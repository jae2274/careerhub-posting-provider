package app

import (
	"careerhub-dataprovider/careerhub/provider/domain/jobposting"
	"careerhub-dataprovider/careerhub/provider/queue"
	"careerhub-dataprovider/careerhub/provider/queue/message_v1"
	"careerhub-dataprovider/careerhub/provider/source"
	"fmt"
)

type AppService struct {
	src  source.JobPostingSource
	repo *jobposting.JobPostingRepo
}

func NewAppService(src source.JobPostingSource, repo *jobposting.JobPostingRepo) *AppService {
	return &AppService{
		src:  src,
		repo: repo,
	}
}

func (as *AppService) Run(quitChan <-chan QuitSignal) error {
	savedJpIds, err := as.repo.GetAllHiring(as.src.Site())
	if err != nil {
		return err
	}

	hiringJpIds, err := source.AllJobPostingIds(as.src)
	if err != nil {
		return err
	}

	separatedIds := SeparateIds(savedJpIds, hiringJpIds)

	ProcessNewJobPostings(as.src, separatedIds.NewPostingIds, quitChan)
	ProcessClosedJobPostings(separatedIds.ClosePostingIds)

	return nil
}

func ProcessNewJobPostings(src source.JobPostingSource, newJpIds []*source.JobPostingId, quitChan <-chan QuitSignal) {

}

func processNewJobPostings(src source.JobPostingSource, jpRepo *jobposting.JobPostingRepo, queue queue.Queue, newJpIds []*source.JobPostingId, quitChan <-chan QuitSignal) {
	for _, newJpId := range newJpIds {
		detail, err := src.Detail(newJpId)

		if err != nil {
			//TODO: error handling
			continue
		}

		message := message_v1.JobPostingInfo{
			Site:        detail.Site,
			PostingId:   detail.PostingId,
			CompanyId:   detail.CompanyId,
			CompanyName: detail.CompanyName,
			JobCategory: detail.JobCategory,
			MainContent: &message_v1.MainContent{
				PostUrl:        detail.MainContent.PostUrl,
				Title:          detail.MainContent.Title,
				Intro:          detail.MainContent.Intro,
				MainTask:       detail.MainContent.MainTask,
				Qualifications: detail.MainContent.Qualifications,
				Preferred:      detail.MainContent.Preferred,
				Benefits:       detail.MainContent.Benefits,
				RecruitProcess: detail.MainContent.RecruitProcess,
			},
			RequiredSkill: detail.RequiredSkill,
			Tags:          detail.Tags,
			RequiredCareer: &message_v1.Career{
				Min: detail.RequiredCareer.Min,
				Max: detail.RequiredCareer.Max,
			},
			PublishedAt: detail.PublishedAt,
			ClosedAt:    detail.ClosedAt,
			Address:     detail.Address,
		}

		//TODO: modify queue interface
		// queue.Send(message)
		fmt.Println(message)
		jpRepo.Save(jobposting.NewJobPosting(message.Site, message.CompanyId))
	}
}

func ProcessClosedJobPostings(closedJpIds []*jobposting.JobPostingId) {

}

func processClosedJobPostings(jpRepo *jobposting.JobPostingRepo, queue queue.Queue, closedJpIds []*jobposting.JobPostingId, quitChan <-chan QuitSignal) {
	closedJpIdMessages := make([]*message_v1.JobPostingId, len(closedJpIds))

	for i, closedJpId := range closedJpIds {
		closedJpIdMessages[i] = &message_v1.JobPostingId{
			Site:      closedJpId.Site,
			PostingId: closedJpId.PostingId,
		}
	}

	message := message_v1.ClosedJobPostings{
		JobPostingIds: closedJpIdMessages,
	}

	//TODO: modify queue interface
	// queue.Send(message)
	fmt.Println(message)

	//TODO: delete closed job postings
	// jpRepo.DeleteAll(closedJpIds)
}
