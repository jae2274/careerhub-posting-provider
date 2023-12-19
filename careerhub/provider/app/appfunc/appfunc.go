package appfunc

import (
	"careerhub-dataprovider/careerhub/provider/domain/company"
	"careerhub-dataprovider/careerhub/provider/domain/jobposting"
	"careerhub-dataprovider/careerhub/provider/queue"
	"careerhub-dataprovider/careerhub/provider/queue/message_v1"
	"careerhub-dataprovider/careerhub/provider/source"
)

type SeparatedIds struct {
	NewPostingIds   []*source.JobPostingId
	ClosePostingIds []*jobposting.JobPostingId
}

func SeparateIds(savedJpIds []*jobposting.JobPostingId, hiringJpIds []*source.JobPostingId) *SeparatedIds {

	savedJpMap := make(map[jobposting.JobPostingId]interface{})
	for _, jp := range savedJpIds {
		savedJpMap[*jp] = false
	}

	newJpIds := make([]*source.JobPostingId, 0)
	for _, srcJpId := range hiringJpIds {
		jpId := jobposting.JobPostingId{Site: srcJpId.Site, PostingId: srcJpId.PostingId}
		if _, ok := savedJpMap[jpId]; ok {
			delete(savedJpMap, jpId)
		} else {
			newJpIds = append(newJpIds, srcJpId)
		}
	}

	closeJpIds := make([]*jobposting.JobPostingId, 0)
	for savedJpId := range savedJpMap {
		func(jpId jobposting.JobPostingId) {
			closeJpIds = append(closeJpIds, &jpId)
		}(savedJpId)
	}

	return &SeparatedIds{
		NewPostingIds:   newJpIds,
		ClosePostingIds: closeJpIds,
	}
}

func CallDetail(src source.JobPostingSource, jpId *source.JobPostingId) (*source.JobPostingDetail, error) {
	return src.Detail(jpId)
}

func SendClosedJobPostings(jpRepo *jobposting.JobPostingRepo, queue *queue.ClosedJobPostingQueue, closedJpIds []*jobposting.JobPostingId) error {
	closedJpIdMessages := make([]*message_v1.JobPostingId, len(closedJpIds))

	for i, closedJpId := range closedJpIds {
		closedJpIdMessages[i] = &message_v1.JobPostingId{
			Site:      closedJpId.Site,
			PostingId: closedJpId.PostingId,
		}
	}

	message := &message_v1.ClosedJobPostings{
		JobPostingIds: closedJpIdMessages,
	}

	queue.Send(message)

	//TODO: delete closed job postings
	// jpRepo.DeleteAll(closedJpIds)

	return nil
}

func SendJobPostingInfo(jpRepo *jobposting.JobPostingRepo, queue *queue.JobPostingQueue, detail *source.JobPostingDetail) error {
	message := &message_v1.JobPostingInfo{
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

	queue.Send(message)
	_, err := jpRepo.Save(jobposting.NewJobPosting(message.Site, message.PostingId))
	return err
}

func ProcessCompany(
	src source.JobPostingSource,
	companyRepo *company.CompanyRepo, //TODO: need to implement
	queue *queue.CompanyQueue,
	companyId *company.CompanyId,
) error {

	companyInfo, err := companyRepo.Get(companyId)

	if err != nil {
		//TODO: error handling
		return err
	} else if companyInfo != nil {
		return nil // already processed
	}

	srcCompany, err := src.Company(companyId.CompanyId)

	if err != nil {
		//TODO: error handling
		return err
	}

	message := &message_v1.Company{
		Site:          srcCompany.Site,
		CompanyId:     srcCompany.CompanyId,
		Name:          srcCompany.Name,
		CompanyUrl:    srcCompany.CompanyUrl,
		CompanyImages: srcCompany.CompanyImages,
		Description:   srcCompany.Description,
		CompanyLogo:   srcCompany.CompanyLogo,
	}

	queue.Send(message)

	_, err = companyRepo.Save(company.NewCompany(src.Site(), companyId.CompanyId))

	return err
}
