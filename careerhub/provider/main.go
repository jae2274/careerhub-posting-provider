package main

import (
	"careerhub-dataprovider/careerhub/provider/app"
	"careerhub-dataprovider/careerhub/provider/awscfg"
	"careerhub-dataprovider/careerhub/provider/domain/company"
	"careerhub-dataprovider/careerhub/provider/domain/jobposting"
	"careerhub-dataprovider/careerhub/provider/dynamo"
	"careerhub-dataprovider/careerhub/provider/queue"
	"careerhub-dataprovider/careerhub/provider/source/jumpit"
	"careerhub-dataprovider/careerhub/provider/vars"
	"log"
	"time"

	"github.com/jae2274/goutils/cchan"
)

func main() {
	application := initApp()
	quit := make(chan app.QuitSignal)
	processedChan, errChan, err := application.Run(quit)
	checkErr(err)

	errorQuit := make(chan app.QuitSignal)
	go cchan.Timeout(10*time.Minute, 10*time.Minute, processedChan, quit)
	go cchan.TooMuchError(10, 10*time.Minute, errChan, errorQuit)

	select {
	case <-errorQuit:
		close(quit)
		log.Fatal("Too much error")
	case <-quit:
		close(errorQuit)
		return
	}
}

func initApp() *app.App {
	envVars, err := vars.Variables()
	checkErr(err)

	awsConfig, err := awscfg.Config()
	checkErr(err)

	dbClient, err := dynamo.NewDbClient(awsConfig, envVars.DbEndpoint)
	checkErr(err)

	jobPostingRepo, err := jobposting.NewJobPostingRepo(dbClient)
	checkErr(err)

	companyRepo, err := company.NewCompanyRepo(dbClient)
	checkErr(err)

	src := jumpit.NewJumpitSource(3000)

	jobPostingQueue, err := queue.NewSQS(awsConfig, envVars.SqsEndpoint, envVars.JobPostingQueue)
	checkErr(err)

	closedQueue, err := queue.NewSQS(awsConfig, envVars.SqsEndpoint, envVars.ClosedQueue)
	checkErr(err)

	companyQueue, err := queue.NewSQS(awsConfig, envVars.SqsEndpoint, envVars.CompanyQueue)
	checkErr(err)

	return app.NewApp(
		src,
		jobPostingRepo,
		companyRepo,
		queue.NewJobPostingQueue(jobPostingQueue),
		queue.NewClosedJobPostingQueue(closedQueue),
		queue.NewCompanyQueue(companyQueue),
	)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
