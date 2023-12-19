package app

import (
	"careerhub-dataprovider/careerhub/provider/app/appfunc"
	"careerhub-dataprovider/careerhub/provider/domain/company"
	"careerhub-dataprovider/careerhub/provider/domain/jobposting"
	"careerhub-dataprovider/careerhub/provider/queue"
	"careerhub-dataprovider/careerhub/provider/source"

	"github.com/jae2274/goutils/cchan"
)

type App struct {
	src             source.JobPostingSource
	jobpostingRepo  *jobposting.JobPostingRepo
	companyRepo     *company.CompanyRepo
	jobPostingQueue queue.Queue
	closedJpQueue   queue.Queue
	companyQueue    queue.Queue
}

func NewApp(src source.JobPostingSource, jobpostingRepo *jobposting.JobPostingRepo, companyRepo *company.CompanyRepo, jobPostingQueue queue.Queue, closedJpQueue queue.Queue, companyQueue queue.Queue) *App {
	return &App{
		src:             src,
		jobpostingRepo:  jobpostingRepo,
		companyRepo:     companyRepo,
		jobPostingQueue: jobPostingQueue,
		closedJpQueue:   closedJpQueue,
		companyQueue:    companyQueue,
	}
}

func (as *App) Run(quitChan <-chan QuitSignal) (<-chan ProcessedSignal, <-chan error, error) {
	hiringJpIds, err := source.AllJobPostingIds(as.src)
	if err != nil {
		return nil, nil, err
	}

	savedJpIds, err := as.jobpostingRepo.GetAllHiring(as.src.Site())
	if err != nil {
		return nil, nil, err
	}

	separatedIds := appfunc.SeparateIds(savedJpIds, hiringJpIds)

	err = appfunc.SendClosedJobPostings(as.jobpostingRepo, as.closedJpQueue, separatedIds.ClosePostingIds)
	if err != nil {
		return nil, nil, err
	}

	errChan := make(chan error, 100)
	detailChan := callDetailApi(as.src, separatedIds.NewPostingIds, errChan, quitChan)
	processedDetailChan := processCompany(as.src, as.companyRepo, as.companyQueue, detailChan, errChan, quitChan)

	processedChan := sendJobPostingInfo(as.jobpostingRepo, as.closedJpQueue, processedDetailChan, errChan, quitChan)

	return processedChan, errChan, nil
}

func callDetailApi(src source.JobPostingSource, newJpIds []*source.JobPostingId, errChan chan<- error, quitChan <-chan QuitSignal) <-chan *source.JobPostingDetail {
	resultChan := make(chan *source.JobPostingDetail)

	go func() {
		defer close(resultChan)

		for _, newJpId := range newJpIds {
			detail, err := appfunc.CallDetail(src, newJpId)

			ok := cchan.SendResult(detail, err, resultChan, errChan, quitChan)

			if !ok {
				return
			}
		}
	}()

	return resultChan
}

func sendJobPostingInfo(jpRepo *jobposting.JobPostingRepo, queue queue.Queue, messageChan <-chan *source.JobPostingDetail, errChan chan<- error, quitChan <-chan QuitSignal) <-chan ProcessedSignal {
	processedChan := make(chan ProcessedSignal, 100)

	go func() {
		defer close(processedChan)

		for {
			received, ok := cchan.ReceiveOrQuit(messageChan, quitChan)
			if !ok {
				return
			}

			err := appfunc.SendJobPostingInfo(jpRepo, queue, *received)

			ok = cchan.SendResult(ProcessedSignal{}, err, processedChan, errChan, quitChan)
			if !ok {
				return
			}
		}
	}()

	return processedChan
}

func processCompany(src source.JobPostingSource,
	companyRepo *company.CompanyRepo, //TODO: need to implement
	queue queue.Queue, detailChan <-chan *source.JobPostingDetail, errChan chan<- error, quitChan <-chan QuitSignal) <-chan *source.JobPostingDetail {

	prosessedChan := make(chan *source.JobPostingDetail)

	go func() {
		for {
			received, ok := cchan.ReceiveOrQuit(detailChan, quitChan)
			if !ok {
				return
			}

			detail := *received
			err := appfunc.ProcessCompany(src, companyRepo, queue, &company.CompanyId{
				Site:      detail.Site,
				CompanyId: detail.CompanyId,
			})

			ok = cchan.SendResult(detail, err, prosessedChan, errChan, quitChan)
			if !ok {
				return
			}
		}
	}()

	return prosessedChan
}
