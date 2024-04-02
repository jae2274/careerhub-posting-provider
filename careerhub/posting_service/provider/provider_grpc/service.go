package provider_grpc

import (
	context "context"
	"time"

	"github.com/jae2274/careerhub-posting-provider/careerhub/posting_service/provider/domain/company"
	"github.com/jae2274/careerhub-posting-provider/careerhub/posting_service/provider/domain/jobposting"

	"github.com/jae2274/goutils/terr"
)

type ProviderGrpcService interface {
	CloseJobPostings(ctx context.Context, jobpostingIds []*jobposting.JobPostingId) error
	GetAllHiring(ctx context.Context, site string) ([]*jobposting.JobPostingId, error)
	IsCompanyRegistered(ctx context.Context, in *company.CompanyId) (bool, error)
	RegisterCompany(ctx context.Context, in *company.CompanyDetail) error
	RegisterJobPostingInfo(ctx context.Context, in *jobposting.JobPostingDetail) error
}
type ProviderGrpcServiceImpl struct {
	grpcClient ProviderGrpcClient
}

func NewProviderGrpcService(grpcClient ProviderGrpcClient) ProviderGrpcService {
	return &ProviderGrpcServiceImpl{grpcClient: grpcClient}
}

func (pgs *ProviderGrpcServiceImpl) IsCompanyRegistered(ctx context.Context, in *company.CompanyId) (bool, error) {
	success, err := pgs.grpcClient.IsCompanyRegistered(ctx, &CompanyId{
		Site:      in.Site,
		CompanyId: in.CompanyId,
	})
	if err != nil {
		return false, terr.Wrap(err)
	}

	return success.Success, nil
}

func (pgs *ProviderGrpcServiceImpl) GetAllHiring(ctx context.Context, site string) ([]*jobposting.JobPostingId, error) {
	jobpostings, err := pgs.grpcClient.GetAllHiring(ctx, &Site{Site: site})
	if err != nil {
		return nil, terr.Wrap(err)
	}

	jobpostingIds := make([]*jobposting.JobPostingId, len(jobpostings.JobPostingIds))
	for i, jp := range jobpostings.JobPostingIds {
		jobpostingIds[i] = &jobposting.JobPostingId{Site: jp.Site, PostingId: jp.PostingId}
	}

	return jobpostingIds, nil
}

func (pgs *ProviderGrpcServiceImpl) CloseJobPostings(ctx context.Context, jobpostingIds []*jobposting.JobPostingId) error {
	pbJobpostingIds := make([]*JobPostingId, 0, len(jobpostingIds))
	for _, id := range jobpostingIds {
		pbJobpostingIds = append(pbJobpostingIds, &JobPostingId{Site: id.Site, PostingId: id.PostingId})
	}

	successRes, err := pgs.grpcClient.CloseJobPostings(ctx, &JobPostings{JobPostingIds: pbJobpostingIds})
	if err != nil {
		return terr.Wrap(err)
	}

	if !successRes.Success {
		return terr.New("failed to close job postings")
	}

	return nil
}

func (pgs *ProviderGrpcServiceImpl) RegisterJobPostingInfo(ctx context.Context, in *jobposting.JobPostingDetail) error {
	pbJobPosting := &JobPostingInfo{
		JobPostingId: &JobPostingId{Site: in.Site, PostingId: in.PostingId},
		CompanyId:    in.CompanyId,
		CompanyName:  in.CompanyName,
		JobCategory:  in.JobCategory,
		MainContent: &MainContent{
			PostUrl:        in.MainContent.PostUrl,
			Title:          in.MainContent.Title,
			Intro:          in.MainContent.Intro,
			MainTask:       in.MainContent.MainTask,
			Qualifications: in.MainContent.Qualifications,
			Preferred:      in.MainContent.Preferred,
			Benefits:       in.MainContent.Benefits,
			RecruitProcess: in.MainContent.RecruitProcess,
		},
		RequiredSkill: in.RequiredSkill,
		Tags:          in.Tags,
		RequiredCareer: &Career{
			Min: in.RequiredCareer.Min,
			Max: in.RequiredCareer.Max,
		},
		PublishedAt:   in.PublishedAt,
		ClosedAt:      in.ClosedAt,
		Address:       in.Address,
		CreatedAt:     time.Now().UnixMilli(),
		ImageUrl:      in.ImageUrl,
		CompanyImages: in.CompanyImages,
	}

	successRes, err := pgs.grpcClient.RegisterJobPostingInfo(ctx, pbJobPosting)

	if err != nil {
		return terr.Wrap(err)
	}

	if !successRes.Success {
		return terr.New("failed to register job posting info. site: " + in.Site + ", postingId: " + in.PostingId)
	}

	return nil
}
func (pgs *ProviderGrpcServiceImpl) RegisterCompany(ctx context.Context, in *company.CompanyDetail) error {
	successRes, err := pgs.grpcClient.RegisterCompany(ctx, &Company{
		Site:          in.Site,
		CompanyId:     in.CompanyId,
		Name:          in.Name,
		CompanyUrl:    in.CompanyUrl,
		CompanyImages: in.CompanyImages,
		Description:   in.Description,
		CompanyLogo:   in.CompanyLogo,
		CreatedAt:     time.Now().UnixMilli(),
	})

	if err != nil {
		return terr.Wrap(err)
	}

	if !successRes.Success {
		return terr.New("failed to register company")
	}

	return nil
}