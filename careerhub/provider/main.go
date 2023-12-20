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
	findNewApp, sendInfoApp := initApp()

	newJobPostingIds, err := findNewApp.Run()
	checkErr(err)

	quit := make(chan app.QuitSignal)
	processedChan, errChan := sendInfoApp.Run(newJobPostingIds, quit)

	timeoutQuit := make(chan app.QuitSignal)
	errorQuit := make(chan app.QuitSignal)
	go cchan.Timeout(10*time.Minute, 10*time.Minute, processedChan, timeoutQuit)
	go cchan.TooMuchError(10, 10*time.Minute, errChan, errorQuit)

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

func initApp() (*app.FindNewJobPostingApp, *app.SendJobPostingApp) {
	src, jobPostingRepo, companyRepo, jobPostingQueue, closedQueue, companyQueue := initComponents()

	return app.NewFindNewJobPostingApp(
			src,
			jobPostingRepo,
			closedQueue,
		), app.NewSendJobPostingApp(
			src,
			jobPostingRepo,
			companyRepo,
			jobPostingQueue,
			companyQueue,
		)
}

func initComponents() (*jumpit.JumpitSource, *jobposting.JobPostingRepo, *company.CompanyRepo, *queue.JobPostingQueue, *queue.ClosedJobPostingQueue, *queue.CompanyQueue) {
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

	return src, jobPostingRepo, companyRepo, queue.NewJobPostingQueue(jobPostingQueue), queue.NewClosedJobPostingQueue(closedQueue), queue.NewCompanyQueue(companyQueue)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
