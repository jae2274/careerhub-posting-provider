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
	"os"
	"time"

	"github.com/jae2274/goutils/cchan"
	"github.com/jae2274/goutils/cchan/pipe"
	"github.com/jae2274/goutils/llog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	appName = "data-processor"
	service = "careerhub"

	ctxKeyTraceID = "trace_id"
)

func main() {
	mainCtx, quitFunc := context.WithCancel(context.Background())
	initLogger(mainCtx)

	findNewApp, sendInfoApp := initApp(mainCtx)

	llog.Msg("Start finding new job postings").Log(mainCtx)
	newJobPostingIds, err := findNewApp.Run(mainCtx)
	checkErr(mainCtx, err)
	llog.Msg("End finding new job postings").Log(mainCtx)

	llog.Msg("Start sending job postings").Log(mainCtx)
	processedChan, errChan := sendInfoApp.Run(mainCtx, newJobPostingIds)

	loggedProcessedChan, loggedErrChan := justLog(mainCtx, processedChan, errChan)

	cchan.Timeout(10*time.Minute, 10*time.Minute, loggedProcessedChan, func() { //차이점은 로그 메시지 뿐
		llog.Msg("Timeout caused").Log(mainCtx)
		quitFunc()
	}, quitFunc)

	cchan.TooMuchError(10, 10*time.Minute, loggedErrChan, func() { //차이점은 로그 메시지 뿐
		llog.Msg("Too much errors caused").Log(mainCtx)
		quitFunc()
	}, quitFunc)

	<-mainCtx.Done()
	llog.Msg("Finish This application").Log(mainCtx)
}

func initLogger(ctx context.Context) {
	llog.SetMetadata("service", service)
	llog.SetMetadata("app", appName)
	llog.SetDefaultContextData(ctxKeyTraceID)

	hostname, err := os.Hostname()
	checkErr(ctx, err)

	llog.SetMetadata("hostname", hostname)
}

func initApp(ctx context.Context) (*app.FindNewJobPostingApp, *app.SendJobPostingApp) {
	jobPostingRepo, companyRepo, grpcClient := initComponents(ctx)
	src := jumpit.NewJumpitSource(ctx, 6000)
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

func justLog(ctx context.Context, processedChan <-chan app.ProcessedSignal, errChan <-chan error) (<-chan app.ProcessedSignal, <-chan error) {
	loggedProcessedChan := pipe.PassThrough(ctx, processedChan, func(signal app.ProcessedSignal) {
		llog.Msg("Processed").Data("site", signal.Site).Data("postingId", signal.PostingId).Log(ctx)
	})

	loggedErrChan := pipe.PassThrough(ctx, errChan, func(err error) {
		llog.LogErr(ctx, err)
	})

	return loggedProcessedChan, loggedErrChan
}
