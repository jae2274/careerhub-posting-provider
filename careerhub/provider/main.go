package main

//aws pipeline trigger
import (
	"careerhub-dataprovider/careerhub/provider/app"
	"careerhub-dataprovider/careerhub/provider/awscfg"
	"careerhub-dataprovider/careerhub/provider/domain/company"
	"careerhub-dataprovider/careerhub/provider/domain/jobposting"
	"careerhub-dataprovider/careerhub/provider/dynamo"
	"careerhub-dataprovider/careerhub/provider/processor_grpc"
	"careerhub-dataprovider/careerhub/provider/source/jumpit"
	"careerhub-dataprovider/careerhub/provider/vars"
	"context"
	"log"
	"os"
	"time"

	"github.com/jae2274/goutils/cchan"
	"github.com/jae2274/goutils/cchan/pipe"
	"github.com/jae2274/goutils/llog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()
	quit := make(chan app.QuitSignal)
	findNewApp, sendInfoApp := initApp(ctx, quit)

	llog.Msg("Start finding new job postings").Log(ctx)
	newJobPostingIds, err := findNewApp.Run()
	checkErr(ctx, err)
	llog.Msg("End finding new job postings").Data("jobPostingCount", len(newJobPostingIds)).Log(ctx)

	llog.Msg("Start getting and sending job posting infos").Log(ctx)
	processedChan, errChan := sendInfoApp.Run(newJobPostingIds, quit)
	loggedProcessedChan, loggedErrChan := justLog(ctx, processedChan, errChan, quit)

	timeoutQuit := make(chan app.QuitSignal)
	errorQuit := make(chan app.QuitSignal)
	go cchan.Timeout(10*time.Minute, 10*time.Minute, loggedProcessedChan, timeoutQuit)
	go cchan.TooMuchError(10, 10*time.Minute, loggedErrChan, errorQuit)

	select {
	case <-errorQuit:
		close(quit)
		log.Fatal("Too much error")
	case <-timeoutQuit:
		close(quit)
		log.Fatal("Timeout")
	case <-quit:
		close(errorQuit)
		close(timeoutQuit)
		return
	}
}

func initApp[QUIT any](ctx context.Context, quitChan <-chan QUIT) (*app.FindNewJobPostingApp, *app.SendJobPostingApp) {
	jobPostingRepo, companyRepo, grpcClient := initComponents(ctx)
	src := jumpit.NewJumpitSource(3000, quitChan)
	return app.NewFindNewJobPostingApp(
			src,
			jobPostingRepo,
			grpcClient,
		), app.NewSendJobPostingApp(
			src,
			jobPostingRepo,
			companyRepo,
			grpcClient,
		)
}

func initComponents(ctx context.Context) (*jobposting.JobPostingRepo, *company.CompanyRepo, processor_grpc.DataProcessorClient) {
	envVars, err := vars.Variables()
	checkErr(ctx, err)

	awsConfig, err := awscfg.Config()
	checkErr(ctx, err)

	dbClient, err := dynamo.NewDbClient(awsConfig, envVars.DbEndpoint)
	checkErr(ctx, err)

	jobPostingRepo, err := jobposting.NewJobPostingRepo(dbClient)
	checkErr(ctx, err)

	companyRepo, err := company.NewCompanyRepo(dbClient)
	checkErr(ctx, err)

	conn, err := grpc.Dial(envVars.GrpcEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	checkErr(ctx, err)

	grpcClient := processor_grpc.NewDataProcessorClient(conn)

	return jobPostingRepo, companyRepo, grpcClient
}

func checkErr(ctx context.Context, err error) {
	if err != nil {
		llog.LogErr(ctx, err)
		os.Exit(1)
	}
}

func justLog(ctx context.Context, processedChan <-chan app.ProcessedSignal, errChan <-chan error, quitChan <-chan app.QuitSignal) (<-chan app.ProcessedSignal, <-chan error) {
	loggedProcessedChan := pipe.PassThrough(processedChan, quitChan, func(signal app.ProcessedSignal) {
		llog.Msg("Processed").Data("site", signal.Site).Data("postingId", signal.PostingId).Log(ctx)
	})

	loggedErrChan := pipe.PassThrough(errChan, quitChan, func(err error) {
		llog.LogErr(ctx, err)
	})

	return loggedProcessedChan, loggedErrChan
}
