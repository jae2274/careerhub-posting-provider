package main

//aws pipeline trigger
import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/jae2274/careerhub-posting-provider/careerhub/posting_service/provider/app"
	"github.com/jae2274/careerhub-posting-provider/careerhub/posting_service/provider/provider_grpc"
	"github.com/jae2274/careerhub-posting-provider/careerhub/posting_service/provider/source"
	"github.com/jae2274/careerhub-posting-provider/careerhub/posting_service/provider/source/jumpit"
	"github.com/jae2274/careerhub-posting-provider/careerhub/posting_service/provider/source/wanted"
	"github.com/jae2274/careerhub-posting-provider/careerhub/posting_service/provider/vars"

	"github.com/jae2274/goutils/cchan"
	"github.com/jae2274/goutils/cchan/pipe"
	"github.com/jae2274/goutils/llog"
	"github.com/jae2274/goutils/mw/grpcmw"
	"github.com/jae2274/goutils/terr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	appName = "posting-provider"
	service = "careerhub"

	ctxKeyTraceID = "trace_id"
)

func initLogger(ctx context.Context) error {
	llog.SetMetadata("service", service)
	llog.SetMetadata("app", appName)
	llog.SetDefaultContextData(ctxKeyTraceID)

	hostname, err := os.Hostname()
	if err != nil {
		return err
	}

	llog.SetMetadata("hostname", hostname)

	return nil
}

func siteFlag() string {
	siteFlag := flag.String("site", "", "site to crawl")
	flag.Parse()

	return *siteFlag
}

func main() {
	mainCtx, quitFunc := context.WithCancel(context.Background())

	err := initLogger(mainCtx)
	checkErr(mainCtx, err)
	llog.Info(mainCtx, "Start Application")

	envVars, err := vars.Variables()
	checkErr(mainCtx, err)
	site := siteFlag()

	llog.SetMetadata("site", site)
	llog.Info(mainCtx, fmt.Sprintf("Site flag is set. site: %s", site))

	findNewApp, sendInfoApp := initApp(mainCtx, site, envVars)

	llog.Msg("Start finding new job postings").Log(mainCtx)
	separateId, err := findNewApp.Run(mainCtx)
	checkErr(mainCtx, err)
	llog.Msg("End finding new job postings").Datas(
		map[string]any{
			"totalCount":         separateId.TotalCount,
			"newPostingCount":    len(separateId.NewPostingIds),
			"closedPostingCount": len(separateId.ClosePostingIds),
		},
	).Log(mainCtx)

	llog.Msg("Start sending job postings").Log(mainCtx)
	processedChan, errChan := sendInfoApp.Run(mainCtx, separateId.NewPostingIds)

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
	jobPostingClient, reviewClient := initComponents(ctx, envVars)
	src, err := jobPostingSource(ctx, site)
	checkErr(ctx, err)

	srv := provider_grpc.NewProviderGrpcService(jobPostingClient, reviewClient)

	return app.NewFindNewJobPostingApp(src, srv), app.NewSendJobPostingApp(src, srv)
}

func initComponents(ctx context.Context, envVars *vars.Vars) (provider_grpc.ProviderGrpcClient, provider_grpc.CrawlingTaskGrpcClient) {
	conn, err := grpc.Dial(envVars.JobPostingGrpcEndpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainStreamInterceptor(grpcmw.SetTraceIdStreamMW()),
		grpc.WithChainUnaryInterceptor(grpcmw.SetTraceIdUnaryMW()),
	)
	checkErr(ctx, err)

	jobPostingClient := provider_grpc.NewProviderGrpcClient(conn)

	conn, err = grpc.Dial(envVars.ReviewGrpcEndpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainStreamInterceptor(grpcmw.SetTraceIdStreamMW()),
		grpc.WithChainUnaryInterceptor(grpcmw.SetTraceIdUnaryMW()),
	)
	checkErr(ctx, err)

	reviewClient := provider_grpc.NewCrawlingTaskGrpcClient(conn)

	return jobPostingClient, reviewClient
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
