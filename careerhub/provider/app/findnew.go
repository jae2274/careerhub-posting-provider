package app

import (
	"careerhub-dataprovider/careerhub/provider/app/appfunc"
	"careerhub-dataprovider/careerhub/provider/domain/jobposting"
	"careerhub-dataprovider/careerhub/provider/processor_grpc"
	"careerhub-dataprovider/careerhub/provider/source"
)

type FindNewJobPostingApp struct {
	src            source.JobPostingSource
	jobpostingRepo *jobposting.JobPostingRepo
	grpcClient     processor_grpc.DataProcessorClient
}

func NewFindNewJobPostingApp(src source.JobPostingSource, jobpostingRepo *jobposting.JobPostingRepo, closedJpQueue processor_grpc.DataProcessorClient) *FindNewJobPostingApp {
	return &FindNewJobPostingApp{
		src:            src,
		jobpostingRepo: jobpostingRepo,
		grpcClient:     closedJpQueue,
	}
}

func (f *FindNewJobPostingApp) Run() ([]*source.JobPostingId, error) {
	separatedIds, err := f.separateIds()
	if err != nil {
		return nil, err
	}

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
