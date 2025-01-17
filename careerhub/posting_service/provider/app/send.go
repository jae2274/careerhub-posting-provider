package app

import (
	"context"

	"github.com/jae2274/careerhub-posting-provider/careerhub/posting_service/provider/domain/company"
	"github.com/jae2274/careerhub-posting-provider/careerhub/posting_service/provider/domain/jobposting"
	"github.com/jae2274/careerhub-posting-provider/careerhub/posting_service/provider/provider_grpc"
	"github.com/jae2274/careerhub-posting-provider/careerhub/posting_service/provider/source"

	"github.com/jae2274/goutils/cchan/pipe"
)

type SendJobPostingApp struct {
	src         source.JobPostingSource
	grpcService provider_grpc.ProviderGrpcService
}

func NewSendJobPostingApp(src source.JobPostingSource, grpcService provider_grpc.ProviderGrpcService) *SendJobPostingApp {
	return &SendJobPostingApp{
		src:         src,
		grpcService: grpcService,
	}
}

func (s *SendJobPostingApp) Run(ctx context.Context, newIds []*jobposting.JobPostingId) (<-chan ProcessedSignal, <-chan error) {

	processedChan, errChan := s.createPipeline(ctx, newIds)

	return processedChan, errChan
}

func (s *SendJobPostingApp) createPipeline(ctx context.Context, newJpIds []*jobposting.JobPostingId) (<-chan ProcessedSignal, <-chan error) {
	jobPostingIdChan := newJobPostingChan(newJpIds)

	step1 := pipe.NewStep(nil,
		func(jpId *jobposting.JobPostingId) (*jobposting.JobPostingDetail, error) {
			return s.src.Detail(jpId)
		})
	step2 := pipe.NewStep(nil,
		func(detail *jobposting.JobPostingDetail) (*jobposting.JobPostingDetail, error) {

			isRegistered, err := s.grpcService.IsCompanyRegistered(context.TODO(), &company.CompanyId{
				Site:      detail.Site,
				CompanyId: detail.CompanyId,
			})

			if err != nil {
				return detail, err
			} else if isRegistered {
				return detail, nil // already processed
			}

			srcCompany, err := s.src.Company(detail.CompanyId)

			if err != nil {
				return detail, err
			}

			err = s.grpcService.RegisterCompany(context.TODO(), srcCompany)
			if err != nil {
				return detail, err
			}

			return detail, nil
		})
	step3 := pipe.NewStep(nil,
		func(detail *jobposting.JobPostingDetail) (ProcessedSignal, error) {
			err := s.grpcService.RegisterJobPostingInfo(context.TODO(), detail)
			if err != nil {
				return ProcessedSignal{Site: detail.Site, PostingId: detail.PostingId}, err
			}

			return ProcessedSignal{Site: detail.Site, PostingId: detail.PostingId}, nil
		})

	return pipe.Pipeline3(ctx, jobPostingIdChan, step1, step2, step3)
}

func newJobPostingChan(newJpIds []*jobposting.JobPostingId) <-chan *jobposting.JobPostingId {
	resultChan := make(chan *jobposting.JobPostingId)

	go func() {
		defer close(resultChan)

		for _, newJpId := range newJpIds {
			resultChan <- newJpId
		}
	}()

	return resultChan
}
