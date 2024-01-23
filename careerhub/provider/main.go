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
	"log"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	quit := make(chan app.QuitSignal)
	findNewApp, sendInfoApp := initApp(quit)

	newJobPostingIds, err := findNewApp.Run()
	checkErr(err)

	processedChan, errChan := sendInfoApp.Run(newJobPostingIds, quit)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for err := range errChan {
			log.Println(err)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for processed := range processedChan {
			log.Println(processed)
		}
	}()

	wg.Wait()
	// timeoutQuit := make(chan app.QuitSignal)
	// errorQuit := make(chan app.QuitSignal)
	// go cchan.Timeout(10*time.Minute, 10*time.Minute, processedChan, timeoutQuit)
	// go cchan.TooMuchError(10, 10*time.Minute, errChan, errorQuit)

	// select {
	// case <-errorQuit:
	// 	close(quit)
	// 	log.Fatal("Too much error")
	// case <-timeoutQuit:
	// 	close(quit)
	// 	log.Fatal("Timeout")
	// case <-quit:
	// 	close(errorQuit)
	// 	close(timeoutQuit)
	// 	return
	// }
}

func initApp[QUIT any](quitChan <-chan QUIT) (*app.FindNewJobPostingApp, *app.SendJobPostingApp) {
	jobPostingRepo, companyRepo, grpcClient := initComponents()
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

func initComponents() (*jobposting.JobPostingRepo, *company.CompanyRepo, processor_grpc.DataProcessorClient) {
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

	conn, err := grpc.Dial(envVars.GrpcEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	checkErr(err)

	grpcClient := processor_grpc.NewDataProcessorClient(conn)

	return jobPostingRepo, companyRepo, grpcClient
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
