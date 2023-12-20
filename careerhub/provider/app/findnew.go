package app

import (
	"careerhub-dataprovider/careerhub/provider/app/appfunc"
	"careerhub-dataprovider/careerhub/provider/domain/jobposting"
	"careerhub-dataprovider/careerhub/provider/queue"
	"careerhub-dataprovider/careerhub/provider/source"
)

type FindNewJobPostingApp struct {
	src            source.JobPostingSource
	jobpostingRepo *jobposting.JobPostingRepo
	closedJpQueue  *queue.ClosedJobPostingQueue
}

func NewFindNewJobPostingApp(src source.JobPostingSource, jobpostingRepo *jobposting.JobPostingRepo, closedJpQueue *queue.ClosedJobPostingQueue) *FindNewJobPostingApp {
	return &FindNewJobPostingApp{
		src:            src,
		jobpostingRepo: jobpostingRepo,
		closedJpQueue:  closedJpQueue,
	}
}

func (f *FindNewJobPostingApp) Run() ([]*source.JobPostingId, error) {
	separatedIds, err := f.separateIds()
	if err != nil {
		return nil, err
	}

	err = appfunc.SendClosedJobPostings(f.jobpostingRepo, f.closedJpQueue, separatedIds.ClosePostingIds)
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
