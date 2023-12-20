package app

import (
	"careerhub-dataprovider/careerhub/provider/app/appfunc"
	"careerhub-dataprovider/careerhub/provider/domain/company"
	"careerhub-dataprovider/careerhub/provider/domain/jobposting"
	"careerhub-dataprovider/careerhub/provider/queue"
	"careerhub-dataprovider/careerhub/provider/source"

	"github.com/jae2274/goutils/cchan/pipe"
)

type SendJobPostingApp struct {
	src             source.JobPostingSource
	jobpostingRepo  *jobposting.JobPostingRepo
	companyRepo     *company.CompanyRepo
	jobPostingQueue *queue.JobPostingQueue
	companyQueue    *queue.CompanyQueue
}

func NewSendJobPostingApp(src source.JobPostingSource, jobpostingRepo *jobposting.JobPostingRepo, companyRepo *company.CompanyRepo, jobPostingQueue *queue.JobPostingQueue, companyQueue *queue.CompanyQueue) *SendJobPostingApp {
	return &SendJobPostingApp{
		src:             src,
		jobpostingRepo:  jobpostingRepo,
		companyRepo:     companyRepo,
		jobPostingQueue: jobPostingQueue,
		companyQueue:    companyQueue,
	}
}

func (s *SendJobPostingApp) Run(newIds []*source.JobPostingId, quitChan <-chan QuitSignal) (<-chan ProcessedSignal, <-chan error) {

	processedChan, errChan := s.createPipeline(newIds, quitChan)

	return processedChan, errChan
}

func (s *SendJobPostingApp) createPipeline(newJpIds []*source.JobPostingId, quitChan <-chan QuitSignal) (<-chan ProcessedSignal, <-chan error) {
	jobPostingIdChan := newJobPostingChan(newJpIds)

	step1 := pipe.NewStep(nil,
		func(jpId *source.JobPostingId) (*source.JobPostingDetail, error) {
			return appfunc.CallDetail(s.src, jpId)
		})
	step2 := pipe.NewStep(nil,
		func(detail *source.JobPostingDetail) (*source.JobPostingDetail, error) {
			return detail, appfunc.ProcessCompany(s.src, s.companyRepo, s.companyQueue, &company.CompanyId{
				Site:      detail.Site,
				CompanyId: detail.CompanyId,
			})
		})
	step3 := pipe.NewStep(nil,
		func(detail *source.JobPostingDetail) (ProcessedSignal, error) {
			return ProcessedSignal{}, appfunc.SendJobPostingInfo(s.jobpostingRepo, s.jobPostingQueue, detail)
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
