package tinit

import (
	"context"
	"fmt"
	"runtime"
	"testing"

	"github.com/jae2274/careerhub-posting-provider/careerhub/posting_service/provider/domain/company"
	"github.com/jae2274/careerhub-posting-provider/careerhub/posting_service/provider/domain/jobposting"
	"github.com/jae2274/careerhub-posting-provider/careerhub/posting_service/provider/provider_grpc"
	"github.com/jae2274/careerhub-posting-provider/careerhub/posting_service/provider/vars"

	"google.golang.org/grpc"
)

type MockJobPostingGrpcClient interface {
	provider_grpc.ProviderGrpcClient
	GetCompanyCount() int
	GetJobPostingCount() int
	GetCompany(*company.CompanyId) *company.CompanyDetail
	GetJobPosting(*jobposting.JobPostingId) *jobposting.JobPostingDetail
}

type MockJobPostingGrpcClientImpl struct {
	JobPostingInfos []*provider_grpc.JobPostingInfo
	Companies       []*provider_grpc.Company
}

func (m *MockJobPostingGrpcClientImpl) IsCompanyRegistered(ctx context.Context, in *provider_grpc.CompanyId, opts ...grpc.CallOption) (*provider_grpc.BoolResponse, error) {
	for _, company := range m.Companies {
		if company.CompanyId == in.CompanyId && company.Site == in.Site {
			return &provider_grpc.BoolResponse{Success: true}, nil
		}
	}

	return &provider_grpc.BoolResponse{Success: false}, nil
}

func (m *MockJobPostingGrpcClientImpl) GetAllHiring(ctx context.Context, in *provider_grpc.Site, opts ...grpc.CallOption) (*provider_grpc.JobPostings, error) {

	jobPostingIds := make([]*provider_grpc.JobPostingId, len(m.JobPostingInfos))
	for i, jp := range m.JobPostingInfos {
		jobPostingIds[i] = &provider_grpc.JobPostingId{Site: jp.JobPostingId.Site, PostingId: jp.JobPostingId.PostingId}
	}

	return &provider_grpc.JobPostings{JobPostingIds: jobPostingIds}, nil
}

func (m *MockJobPostingGrpcClientImpl) CloseJobPostings(ctx context.Context, in *provider_grpc.JobPostings, opts ...grpc.CallOption) (*provider_grpc.BoolResponse, error) {

	newJobPostings := make([]*provider_grpc.JobPostingInfo, 0)
	for _, existedJp := range m.JobPostingInfos {
		for _, closedJp := range in.JobPostingIds {
			if existedJp.JobPostingId.Site == closedJp.Site && existedJp.JobPostingId.PostingId == closedJp.PostingId {
				continue
			}
			newJobPostings = append(newJobPostings, existedJp)
		}
	}

	m.JobPostingInfos = newJobPostings
	return &provider_grpc.BoolResponse{Success: true}, nil
}

func (m *MockJobPostingGrpcClientImpl) RegisterJobPostingInfo(ctx context.Context, in *provider_grpc.JobPostingInfo, opts ...grpc.CallOption) (*provider_grpc.BoolResponse, error) {
	m.JobPostingInfos = append(m.JobPostingInfos, in)
	return &provider_grpc.BoolResponse{Success: true}, nil
}

func (m *MockJobPostingGrpcClientImpl) RegisterCompany(ctx context.Context, in *provider_grpc.Company, opts ...grpc.CallOption) (*provider_grpc.BoolResponse, error) {
	m.Companies = append(m.Companies, in)
	return &provider_grpc.BoolResponse{Success: true}, nil
}

func (m *MockJobPostingGrpcClientImpl) GetCompanyCount() int {
	return len(m.Companies)
}

func (m *MockJobPostingGrpcClientImpl) GetJobPostingCount() int {
	return len(m.JobPostingInfos)
}

func (m *MockJobPostingGrpcClientImpl) GetCompany(in *company.CompanyId) *company.CompanyDetail {
	for _, c := range m.Companies {
		if c.CompanyId == in.CompanyId && c.Site == in.Site {
			return &company.CompanyDetail{
				Site:          c.Site,
				CompanyId:     c.CompanyId,
				Name:          c.Name,
				CompanyUrl:    c.CompanyUrl,
				CompanyImages: c.CompanyImages,
				Description:   c.Description,
				CompanyLogo:   c.CompanyLogo,
			}
		}
	}

	return nil
}

type MockReviewGrpcClient interface {
	provider_grpc.CrawlingTaskGrpcClient
	GetCrawlingTasks() []*provider_grpc.AddCrawlingTaskRequest
	GetCrawlingTask(companyName string) *provider_grpc.AddCrawlingTaskRequest
}

type MockReviewGrpcClientImpl struct {
	CrawlingTasks []*provider_grpc.AddCrawlingTaskRequest
}

func (m *MockReviewGrpcClientImpl) AddCrawlingTask(ctx context.Context, in *provider_grpc.AddCrawlingTaskRequest, opts ...grpc.CallOption) (*provider_grpc.AddCrawlingTaskResponse, error) {
	m.CrawlingTasks = append(m.CrawlingTasks, in)
	return &provider_grpc.AddCrawlingTaskResponse{Status: "success"}, nil
}

func (m *MockReviewGrpcClientImpl) GetCrawlingTasks() []*provider_grpc.AddCrawlingTaskRequest {
	return m.CrawlingTasks
}

func (m *MockReviewGrpcClientImpl) GetCrawlingTask(companyName string) *provider_grpc.AddCrawlingTaskRequest {
	for _, task := range m.CrawlingTasks {
		if task.CompanyName == companyName {
			return task
		}
	}

	return nil
}

func (m *MockJobPostingGrpcClientImpl) GetJobPosting(in *jobposting.JobPostingId) *jobposting.JobPostingDetail {
	for _, jp := range m.JobPostingInfos {
		if jp.JobPostingId.PostingId == in.PostingId && jp.JobPostingId.Site == in.Site {
			return &jobposting.JobPostingDetail{
				Site:        jp.JobPostingId.Site,
				PostingId:   jp.JobPostingId.PostingId,
				CompanyId:   jp.CompanyId,
				CompanyName: jp.CompanyName,
				JobCategory: jp.JobCategory,
				MainContent: jobposting.MainContent{
					PostUrl:        jp.MainContent.PostUrl,
					Title:          jp.MainContent.Title,
					Intro:          jp.MainContent.Intro,
					MainTask:       jp.MainContent.MainTask,
					Qualifications: jp.MainContent.Qualifications,
					Preferred:      jp.MainContent.Preferred,
					Benefits:       jp.MainContent.Benefits,
					RecruitProcess: jp.MainContent.RecruitProcess,
				},
				RequiredSkill: jp.RequiredSkill,
				Tags:          jp.Tags,
				RequiredCareer: jobposting.Career{
					Min: jp.RequiredCareer.Min,
					Max: jp.RequiredCareer.Max,
				},
				PublishedAt:   jp.PublishedAt,
				ClosedAt:      jp.ClosedAt,
				Address:       jp.Address,
				ImageUrl:      jp.ImageUrl,
				CompanyImages: jp.CompanyImages,
			}
		}
	}

	return nil
}

func InitGrpcClient(t *testing.T) (MockJobPostingGrpcClient, MockReviewGrpcClient) {
	variables, err := vars.Variables()
	checkError(t, err)

	if variables.JobPostingGrpcEndpoint == "" {
		t.Fatal("JOB_POSTING_GRPC_ENDPOINT is not set")
		t.FailNow()
	}

	return &MockJobPostingGrpcClientImpl{
			JobPostingInfos: make([]*provider_grpc.JobPostingInfo, 0),
			Companies:       make([]*provider_grpc.Company, 0),
		}, &MockReviewGrpcClientImpl{
			CrawlingTasks: make([]*provider_grpc.AddCrawlingTaskRequest, 0),
		}
}

func checkError(t *testing.T, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d\n", file, line)
		t.Error(err)
		t.FailNow()
	}
}
