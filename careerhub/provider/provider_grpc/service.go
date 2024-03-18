package provider_grpc

import (
	"careerhub-dataprovider/careerhub/provider/domain/company"
	"careerhub-dataprovider/careerhub/provider/domain/jobposting"
	context "context"
	"time"

	"github.com/jae2274/goutils/terr"
)

type ProviderGrpcService struct {
	grpcClient ProviderGrpcClient
}

func NewProviderGrpcService(grpcClient ProviderGrpcClient) *ProviderGrpcService {
	return &ProviderGrpcService{grpcClient: grpcClient}
}

func (pgs *ProviderGrpcService) IsCompanyRegistered(ctx context.Context, in *company.CompanyId) (bool, error) {
	success, err := pgs.grpcClient.IsCompanyRegistered(ctx, &CompanyId{
		Site:      in.Site,
		CompanyId: in.CompanyId,
	})
	if err != nil {
		return false, terr.Wrap(err)
	}

	return success.Success, nil
}

func (pgs *ProviderGrpcService) CloseJobPostings(ctx context.Context, jobpostingIds []*jobposting.JobPostingId) error {
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

func (pgs *ProviderGrpcService) RegisterJobPostingInfo(ctx context.Context, in *jobposting.JobPostingDetail) error {
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
		return terr.New("failed to register job posting info")
	}

	return nil
}
func (pgs *ProviderGrpcService) RegisterCompany(ctx context.Context, in *company.CompanyDetail) error {
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
