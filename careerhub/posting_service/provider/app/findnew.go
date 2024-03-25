package app

import (
	"context"

	"github.com/jae2274/careerhub-posting-provider/careerhub/posting_service/provider/app/appfunc"
	"github.com/jae2274/careerhub-posting-provider/careerhub/posting_service/provider/provider_grpc"
	"github.com/jae2274/careerhub-posting-provider/careerhub/posting_service/provider/source"
)

type FindNewJobPostingApp struct {
	src         source.JobPostingSource
	grpcService provider_grpc.ProviderGrpcService
}

func NewFindNewJobPostingApp(src source.JobPostingSource, grpcService provider_grpc.ProviderGrpcService) *FindNewJobPostingApp {
	return &FindNewJobPostingApp{
		src:         src,
		grpcService: grpcService,
	}
}

func (f *FindNewJobPostingApp) Run(ctx context.Context) (*appfunc.SeparatedIds, error) {
	separatedIds, err := f.separateIds()
	if err != nil {
		return nil, err
	}

	err = f.grpcService.CloseJobPostings(ctx, separatedIds.ClosePostingIds)
	if err != nil {
		return nil, err
	}

	return separatedIds, nil
}

func (a *FindNewJobPostingApp) separateIds() (*appfunc.SeparatedIds, error) {
	hiringJpIds, err := source.AllJobPostingIds(a.src)
	if err != nil {
		return nil, err
	}

	savedJpIds, err := a.grpcService.GetAllHiring(context.Background(), a.src.Site())
	if err != nil {
		return nil, err
	}

	return appfunc.SeparateIds(savedJpIds, hiringJpIds), nil
}
