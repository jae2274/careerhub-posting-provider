package tinit

import (
	"careerhub-dataprovider/careerhub/provider/provider_grpc"
	"careerhub-dataprovider/careerhub/provider/vars"
	"context"
	"fmt"
	"runtime"
	"testing"

	"google.golang.org/grpc"
)

type MockGrpcClient interface {
	provider_grpc.ProviderGrpcClient
	GetClosedJpIds() []*provider_grpc.JobPostings
	GetJobPostingInfo() []*provider_grpc.JobPostingInfo
	GetCompany() []*provider_grpc.Company
}

type MockGrpcClientImpl struct {
	JobPostingInfos []*provider_grpc.JobPostingInfo
	JobPostings     []*provider_grpc.JobPostings
	Companies       []*provider_grpc.Company
}

func (m *MockGrpcClientImpl) IsCompanyRegistered(ctx context.Context, in *provider_grpc.CompanyId, opts ...grpc.CallOption) (*provider_grpc.BoolResponse, error) {
	for _, company := range m.Companies {
		if company.CompanyId == in.CompanyId && company.Site == in.Site {
			return &provider_grpc.BoolResponse{Success: true}, nil
		}
	}

	return &provider_grpc.BoolResponse{Success: false}, nil
}

func (m *MockGrpcClientImpl) CloseJobPostings(ctx context.Context, in *provider_grpc.JobPostings, opts ...grpc.CallOption) (*provider_grpc.BoolResponse, error) {
	m.JobPostings = append(m.JobPostings, in)
	return &provider_grpc.BoolResponse{Success: true}, nil
}

func (m *MockGrpcClientImpl) RegisterJobPostingInfo(ctx context.Context, in *provider_grpc.JobPostingInfo, opts ...grpc.CallOption) (*provider_grpc.BoolResponse, error) {
	m.JobPostingInfos = append(m.JobPostingInfos, in)
	return &provider_grpc.BoolResponse{Success: true}, nil
}

func (m *MockGrpcClientImpl) RegisterCompany(ctx context.Context, in *provider_grpc.Company, opts ...grpc.CallOption) (*provider_grpc.BoolResponse, error) {
	m.Companies = append(m.Companies, in)
	return &provider_grpc.BoolResponse{Success: true}, nil
}

func (m *MockGrpcClientImpl) GetClosedJpIds() []*provider_grpc.JobPostings {
	return m.JobPostings
}

func (m *MockGrpcClientImpl) GetJobPostingInfo() []*provider_grpc.JobPostingInfo {
	return m.JobPostingInfos
}

func (m *MockGrpcClientImpl) GetCompany() []*provider_grpc.Company {
	return m.Companies
}

func InitGrpcClient(t *testing.T) MockGrpcClient {
	variables, err := vars.Variables()
	checkError(t, err)

	if variables.GrpcEndpoint == "" {
		t.Fatal("GRPC_ENDPOINT is not set")
		t.FailNow()
	}

	return &MockGrpcClientImpl{
		JobPostingInfos: make([]*provider_grpc.JobPostingInfo, 0),
		JobPostings:     make([]*provider_grpc.JobPostings, 0),
		Companies:       make([]*provider_grpc.Company, 0),
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
