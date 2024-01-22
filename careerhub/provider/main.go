package main

//aws pipeline trigger
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
	quit := make(chan app.QuitSignal)
	findNewApp, sendInfoApp := initApp(quit)

	newJobPostingIds, err := findNewApp.Run()
	checkErr(err)

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

func initApp[QUIT any](quitChan <-chan QUIT) (*app.FindNewJobPostingApp, *app.SendJobPostingApp) {
	jobPostingRepo, companyRepo, jobPostingQueue, closedQueue, companyQueue := initComponents()
	src := jumpit.NewJumpitSource(3000, quitChan)
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

func initComponents() (*jobposting.JobPostingRepo, *company.CompanyRepo, *queue.JobPostingQueue, *queue.ClosedJobPostingQueue, *queue.CompanyQueue) {
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

	jobPostingQueue, err := queue.NewSQS(awsConfig, envVars.SqsEndpoint, envVars.JobPostingQueue)
	checkErr(err)

	closedQueue, err := queue.NewSQS(awsConfig, envVars.SqsEndpoint, envVars.ClosedQueue)
	checkErr(err)

	companyQueue, err := queue.NewSQS(awsConfig, envVars.SqsEndpoint, envVars.CompanyQueue)
	checkErr(err)

	return jobPostingRepo, companyRepo, queue.NewJobPostingQueue(jobPostingQueue), queue.NewClosedJobPostingQueue(closedQueue), queue.NewCompanyQueue(companyQueue)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
