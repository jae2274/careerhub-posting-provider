package appfunc

import (
	"careerhub-dataprovider/careerhub/provider/domain/company"
	"careerhub-dataprovider/careerhub/provider/domain/jobposting"
	"careerhub-dataprovider/careerhub/provider/provider_grpc"
	"careerhub-dataprovider/careerhub/provider/source"
	"context"
	"time"

	"github.com/jae2274/goutils/terr"
)

type SeparatedIds struct {
	TotalCount      int
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
		TotalCount:      len(hiringJpIds),
		NewPostingIds:   newJpIds,
		ClosePostingIds: closeJpIds,
	}
}

func CallDetail(src source.JobPostingSource, jpId *source.JobPostingId) (*source.JobPostingDetail, error) {
	return src.Detail(jpId)
}

func SendClosedJobPostings(jpRepo *jobposting.JobPostingRepo, grpcClient provider_grpc.ProviderGrpcClient, closedJpIds []*jobposting.JobPostingId) error {

	closedJpIdMessages := make([]*provider_grpc.JobPostingId, len(closedJpIds))

	for i, closedJpId := range closedJpIds {
		closedJpIdMessages[i] = &provider_grpc.JobPostingId{
			Site:      closedJpId.Site,
			PostingId: closedJpId.PostingId,
		}
	}

	message := &provider_grpc.JobPostings{
		JobPostingIds: closedJpIdMessages,
	}

	_, err := grpcClient.CloseJobPostings(context.TODO(), message)
	if err != nil {
		return err
	}

	//TODO: delete closed job postings

	return jpRepo.DeleteAll(closedJpIds)
}

func SendJobPostingInfo(jpRepo *jobposting.JobPostingRepo, grpcClient provider_grpc.ProviderGrpcClient, detail *source.JobPostingDetail) error {
	message := &provider_grpc.JobPostingInfo{
		JobPostingId: &provider_grpc.JobPostingId{
			Site:      detail.Site,
			PostingId: detail.PostingId,
		},
		CompanyId:   detail.CompanyId,
		CompanyName: detail.CompanyName,
		JobCategory: detail.JobCategory,
		MainContent: &provider_grpc.MainContent{
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
		RequiredCareer: &provider_grpc.Career{
			Min: detail.RequiredCareer.Min,
			Max: detail.RequiredCareer.Max,
		},
		PublishedAt:   detail.PublishedAt,
		ClosedAt:      detail.ClosedAt,
		Address:       detail.Address,
		CreatedAt:     time.Now().UnixMilli(),
		ImageUrl:      detail.ImageUrl,
		CompanyImages: detail.CompanyImages,
	}

	_, err := grpcClient.RegisterJobPostingInfo(context.TODO(), message)
	if err != nil {
		return terr.Wrap(err)
	}

	_, err = jpRepo.Save(jobposting.NewJobPosting(message.JobPostingId.Site, message.JobPostingId.PostingId))
	return err
}

func ProcessCompany(
	src source.JobPostingSource,
	companyRepo *company.CompanyRepo, //TODO: need to implement
	grpcClient provider_grpc.ProviderGrpcClient,
	companyId *company.CompanyId,
) error {

	companyInfo, err := companyRepo.Get(companyId)

	if err != nil {
		return err
	} else if companyInfo != nil {
		return nil // already processed
	}

	srcCompany, err := src.Company(companyId.CompanyId)

	if err != nil {
		return err
	}

	message := &provider_grpc.Company{
		Site:          srcCompany.Site,
		CompanyId:     srcCompany.CompanyId,
		Name:          srcCompany.Name,
		CompanyUrl:    srcCompany.CompanyUrl,
		CompanyImages: srcCompany.CompanyImages,
		Description:   srcCompany.Description,
		CompanyLogo:   srcCompany.CompanyLogo,
		CreatedAt:     time.Now().UnixMilli(),
	}

	_, err = grpcClient.RegisterCompany(context.TODO(), message)
	if err != nil {
		return err
	}

	_, err = companyRepo.Save(company.NewCompany(src.Site(), companyId.CompanyId))

	return err
}
