package app

import (
	"careerhub-dataprovider/careerhub/provider/app/appfunc"
	"careerhub-dataprovider/careerhub/provider/domain/company"
	"careerhub-dataprovider/careerhub/provider/domain/jobposting"
	"careerhub-dataprovider/careerhub/provider/processor_grpc"
	"careerhub-dataprovider/careerhub/provider/source"
	"context"

	"github.com/jae2274/goutils/cchan/pipe"
)

type SendJobPostingApp struct {
	src            source.JobPostingSource
	jobpostingRepo *jobposting.JobPostingRepo
	companyRepo    *company.CompanyRepo
	grpcClient     processor_grpc.DataProcessorClient
}

func NewSendJobPostingApp(src source.JobPostingSource, jobpostingRepo *jobposting.JobPostingRepo, companyRepo *company.CompanyRepo, grpcClient processor_grpc.DataProcessorClient) *SendJobPostingApp {
	return &SendJobPostingApp{
		src:            src,
		jobpostingRepo: jobpostingRepo,
		companyRepo:    companyRepo,
		grpcClient:     grpcClient,
	}
}

func (s *SendJobPostingApp) Run(ctx context.Context, newIds []*source.JobPostingId) (<-chan ProcessedSignal, <-chan error) {

	processedChan, errChan := s.createPipeline(ctx, newIds)

	return processedChan, errChan
}

func (s *SendJobPostingApp) createPipeline(ctx context.Context, newJpIds []*source.JobPostingId) (<-chan ProcessedSignal, <-chan error) {
	jobPostingIdChan := newJobPostingChan(newJpIds)

	step1 := pipe.NewStep(nil,
		func(jpId *source.JobPostingId) (*source.JobPostingDetail, error) {
			return appfunc.CallDetail(s.src, jpId)
		})
	step2 := pipe.NewStep(nil,
		func(detail *source.JobPostingDetail) (*source.JobPostingDetail, error) {
			return detail, appfunc.ProcessCompany(s.src, s.companyRepo, s.grpcClient, &company.CompanyId{
				Site:      detail.Site,
				CompanyId: detail.CompanyId,
			})
		})
	step3 := pipe.NewStep(nil,
		func(detail *source.JobPostingDetail) (ProcessedSignal, error) {
			return ProcessedSignal{Site: detail.Site, PostingId: detail.PostingId}, appfunc.SendJobPostingInfo(s.jobpostingRepo, s.grpcClient, detail)
		})

	errChan := make(chan error, 100)
	return pipe.Pipeline3(ctx, jobPostingIdChan, errChan, step1, step2, step3), errChan
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
