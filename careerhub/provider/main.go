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
	go Timeout(10*time.Minute, 10*time.Minute, processedChan, quit)
	go TooMuchError(10, 10*time.Minute, errChan, errorQuit)

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

	return app.NewApp(
		src,
		jobPostingRepo,
		companyRepo,
		&queue.FakeQueue{},
		&queue.FakeQueue{},
		&queue.FakeQueue{},
	)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func TooMuchError(periodErrCount uint, limitErrPeriod time.Duration, errChan <-chan error, quitChan chan app.QuitSignal) {
	defer func() {
		log.Default().Println("TooMuchError closed")
	}()

	var errCount uint = 0
	errCaughtTimes := make([]time.Time, periodErrCount)

	for {
		err, ok := cchan.ReceiveOrQuit(errChan, quitChan)
		if !ok {
			return
		}

		log.Default().Println((*err).Error())
		errCount++
		errCaughtTimes = append(errCaughtTimes, time.Now())
		if len(errCaughtTimes) >= int(periodErrCount) {
			lastErrCaughtTime := errCaughtTimes[len(errCaughtTimes)-1]
			recentErrCaughtTime := errCaughtTimes[len(errCaughtTimes)-10]
			errCaughtPeriod := lastErrCaughtTime.Sub(recentErrCaughtTime)

			if errCaughtPeriod.Abs() < limitErrPeriod.Abs() {
				close(quitChan)
				return
			}
			errCaughtTimes = errCaughtTimes[1:]
		}

	}
}

func Timeout(initDuration, duration time.Duration, processedChan <-chan app.ProcessedSignal, quitChan chan app.QuitSignal) {
	defer func() {
		log.Default().Println("Timeout closed")
	}()
	waitDuration := initDuration

	for {

		select {
		case <-quitChan:
			return
		case <-time.After(waitDuration):
			close(quitChan)
			return
		case _, ok := <-processedChan:
			if !ok {
				return
			}
			waitDuration = duration
		}
	}
}
