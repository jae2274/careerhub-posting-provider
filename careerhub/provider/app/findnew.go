package app

import (
	"careerhub-dataprovider/careerhub/provider/app/appfunc"
	"careerhub-dataprovider/careerhub/provider/domain/jobposting"
	"careerhub-dataprovider/careerhub/provider/provider_grpc"
	"careerhub-dataprovider/careerhub/provider/source"
	"context"

	"github.com/jae2274/goutils/llog"
)

type FindNewJobPostingApp struct {
	src            source.JobPostingSource
	jobpostingRepo *jobposting.JobPostingRepo
	grpcClient     provider_grpc.ProviderGrpcClient
}

func NewFindNewJobPostingApp(src source.JobPostingSource, jobpostingRepo *jobposting.JobPostingRepo, closedJpQueue provider_grpc.ProviderGrpcClient) *FindNewJobPostingApp {
	return &FindNewJobPostingApp{
		src:            src,
		jobpostingRepo: jobpostingRepo,
		grpcClient:     closedJpQueue,
	}
}

func (f *FindNewJobPostingApp) Run(ctx context.Context) ([]*source.JobPostingId, error) {
	separatedIds, err := f.separateIds()
	if err != nil {
		return nil, err
	}
	llog.Msg("End finding new job postings").Datas(
		map[string]any{
			"totalCount":         separatedIds.TotalCount,
			"newPostingCount":    len(separatedIds.NewPostingIds),
			"closedPostingCount": len(separatedIds.ClosePostingIds),
		},
	).Log(ctx)

	err = appfunc.SendClosedJobPostings(f.jobpostingRepo, f.grpcClient, separatedIds.ClosePostingIds)
	if err != nil {
		return nil, err
	}

	return separatedIds.NewPostingIds, nil
}

func (a *FindNewJobPostingApp) separateIds() (*appfunc.SeparatedIds, error) {
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
