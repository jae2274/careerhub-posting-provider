package main

//aws pipeline trigger
import (
	"careerhub-dataprovider/careerhub/provider/app"
	"careerhub-dataprovider/careerhub/provider/domain/company"
	"careerhub-dataprovider/careerhub/provider/domain/jobposting"
	"careerhub-dataprovider/careerhub/provider/logger"
	"careerhub-dataprovider/careerhub/provider/mongocfg"
	"careerhub-dataprovider/careerhub/provider/provider_grpc"
	"careerhub-dataprovider/careerhub/provider/source"
	"careerhub-dataprovider/careerhub/provider/source/jumpit"
	"careerhub-dataprovider/careerhub/provider/source/wanted"
	"careerhub-dataprovider/careerhub/provider/vars"
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/jae2274/goutils/cchan"
	"github.com/jae2274/goutils/cchan/pipe"
	"github.com/jae2274/goutils/llog"
	"github.com/jae2274/goutils/terr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	appName = "data-provider"
	service = "careerhub"

	ctxKeyTraceID = "trace_id"
)

func siteFlag() string {
	siteFlag := flag.String("site", "", "site to crawl")
	flag.Parse()

	return *siteFlag
}

func main() {
	mainCtx, quitFunc := context.WithCancel(context.Background())
	envVars := initEnvVars()

	site := siteFlag()
	appLogger := initLogger(mainCtx, site, envVars.PostLogUrl)

	findNewApp, sendInfoApp := initApp(mainCtx, site, envVars)

	llog.Msg("Start finding new job postings").Log(mainCtx)
	newJobPostingIds, err := findNewApp.Run(mainCtx)
	checkErr(mainCtx, err)
	llog.Msg("End finding new job postings").Log(mainCtx)

	llog.Msg("Start sending job postings").Log(mainCtx)
	processedChan, errChan := sendInfoApp.Run(mainCtx, newJobPostingIds)

	loggedErrChan := justLog(mainCtx, errChan)

	cchan.Timeout(10*time.Minute, 10*time.Minute, processedChan, func() { //차이점은 로그 메시지 뿐
		llog.Msg("Timeout caused").Log(mainCtx)
		quitFunc()
	}, quitFunc)

	cchan.TooMuchError(10, 10*time.Minute, loggedErrChan, func() { //차이점은 로그 메시지 뿐
		llog.Msg("Too much errors caused").Log(mainCtx)
		quitFunc()
	}, quitFunc)

	<-mainCtx.Done()
	llog.Msg("Finish This application").Log(mainCtx)
	appLogger.Wg.Wait()
}

func initEnvVars() *vars.Vars {
	envVars, err := vars.Variables()
	checkErr(context.Background(), err)
	return envVars
}

func initLogger(ctx context.Context, site, postLogUrl string) *logger.AppLogger {
	llog.SetMetadata("service", service)
	llog.SetMetadata("app", appName)
	llog.SetMetadata("site", site)
	llog.SetDefaultContextData(ctxKeyTraceID)

	hostname, err := os.Hostname()
	checkErr(ctx, err)

	llog.SetMetadata("hostname", hostname)

	appLogger, err := logger.NewAppLogger(ctx, postLogUrl)
	checkErr(ctx, err)

	llog.SetDefaultLLoger(appLogger)
	return appLogger
}

func jobPostingSource(ctx context.Context, site string) (source.JobPostingSource, error) {
	delayMilis := int64(6000)

	switch site {
	case jumpit.Site:
		return jumpit.NewJumpitSource(ctx, delayMilis), nil
	case wanted.Site:
		return wanted.NewWantedSource(ctx, delayMilis), nil

	default:
		return nil, terr.New(fmt.Sprintf("site flag is not valid. site: %s", site))
	}
}

func initApp(ctx context.Context, site string, envVars *vars.Vars) (*app.FindNewJobPostingApp, *app.SendJobPostingApp) {
	jobPostingRepo, companyRepo, grpcClient := initComponents(ctx, envVars)
	src, err := jobPostingSource(ctx, site)
	checkErr(ctx, err)

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

func initComponents(ctx context.Context, envVars *vars.Vars) (*jobposting.JobPostingRepo, *company.CompanyRepo, provider_grpc.ProviderGrpcClient) {
	db, err := mongocfg.NewDatabase(envVars.MongoUri, envVars.DbName, envVars.DBUser)
	checkErr(ctx, err)

	jobPostingModel := &jobposting.JobPosting{}
	jobPostingCollection := db.Collection(jobPostingModel.Collection())
	err = mongocfg.CheckIndexViaCollection(jobPostingCollection, jobPostingModel.IndexModels())
	checkErr(ctx, err)
	jobPostingRepo := jobposting.NewJobPostingRepo(jobPostingCollection)

	companyModel := &company.Company{}
	companyCollection := db.Collection(companyModel.Collection())
	err = mongocfg.CheckIndexViaCollection(companyCollection, companyModel.IndexModels())
	checkErr(ctx, err)
	companyRepo := company.NewCompanyRepo(companyCollection)

	conn, err := grpc.Dial(envVars.GrpcEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	checkErr(ctx, err)

	grpcClient := provider_grpc.NewProviderGrpcClient(conn)

	return jobPostingRepo, companyRepo, grpcClient
}

func checkErr(ctx context.Context, err error) {
	if err != nil {
		llog.LogErr(ctx, err)
		os.Exit(1)
	}
}

func justLog(ctx context.Context, errChan <-chan error) <-chan error {
	loggedErrChan := pipe.PassThrough(ctx, errChan, func(err error) {
		llog.LogErr(ctx, err)
	})

	return loggedErrChan
}
