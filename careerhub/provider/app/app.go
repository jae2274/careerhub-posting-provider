package app

import (
	"careerhub-dataprovider/careerhub/provider/app/appfunc"
	"careerhub-dataprovider/careerhub/provider/domain/company"
	"careerhub-dataprovider/careerhub/provider/domain/jobposting"
	"careerhub-dataprovider/careerhub/provider/queue"
	"careerhub-dataprovider/careerhub/provider/source"

	"github.com/jae2274/goutils/cchan/pipe"
)

type App struct {
	src             source.JobPostingSource
	jobpostingRepo  *jobposting.JobPostingRepo
	companyRepo     *company.CompanyRepo
	jobPostingQueue *queue.JobPostingQueue
	closedJpQueue   *queue.ClosedJobPostingQueue
	companyQueue    *queue.CompanyQueue
}

func NewApp(src source.JobPostingSource, jobpostingRepo *jobposting.JobPostingRepo, companyRepo *company.CompanyRepo, jobPostingQueue *queue.JobPostingQueue, closedJpQueue *queue.ClosedJobPostingQueue, companyQueue *queue.CompanyQueue) *App {
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
	separatedIds, err := as.separateIds()
	if err != nil {
		return nil, nil, err
	}

	err = appfunc.SendClosedJobPostings(as.jobpostingRepo, as.closedJpQueue, separatedIds.ClosePostingIds)
	if err != nil {
		return nil, nil, err
	}

	processedChan, errChan := as.createPipeline(separatedIds.NewPostingIds, quitChan)

	return processedChan, errChan, nil
}

func (a *App) separateIds() (*appfunc.SeparatedIds, error) {
	hiringJpIds, err := source.AllJobPostingIds(a.src)
	if err != nil {
		return nil, err
	}

	savedJpIds, err := a.jobpostingRepo.GetAllHiring(a.src.Site())
	if err != nil {
		return nil, err
	}

	return appfunc.SeparateIds(savedJpIds, hiringJpIds), nil
}

func (a *App) createPipeline(newJpIds []*source.JobPostingId, quitChan <-chan QuitSignal) (<-chan ProcessedSignal, <-chan error) {
	jobPostingIdChan := newJobPostingChan(newJpIds)

	step1 := pipe.NewStep(nil,
		func(jpId *source.JobPostingId) (*source.JobPostingDetail, error) {
			return appfunc.CallDetail(a.src, jpId)
		})
	step2 := pipe.NewStep(nil,
		func(detail *source.JobPostingDetail) (*source.JobPostingDetail, error) {
			return detail, appfunc.ProcessCompany(a.src, a.companyRepo, a.companyQueue, &company.CompanyId{
				Site:      detail.Site,
				CompanyId: detail.CompanyId,
			})
		})
	step3 := pipe.NewStep(nil,
		func(detail *source.JobPostingDetail) (ProcessedSignal, error) {
			return ProcessedSignal{}, appfunc.SendJobPostingInfo(a.jobpostingRepo, a.jobPostingQueue, detail)
		})

	return pipe.Pipeline3(jobPostingIdChan, quitChan, 100, step1, step2, step3)
}

func newJobPostingChan(newJpIds []*source.JobPostingId) <-chan *source.JobPostingId {
	resultChan := make(chan *source.JobPostingId)

	go func() {
		defer close(resultChan)

		for _, newJpId := range newJpIds {
			resultChan <- newJpId
		}
	}()

	return resultChan
}
