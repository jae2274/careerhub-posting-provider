package tinit

import (
	"careerhub-dataprovider/careerhub/provider/processor_grpc"
	"careerhub-dataprovider/careerhub/provider/vars"
	"context"
	"fmt"
	"runtime"
	"testing"

	"google.golang.org/grpc"
)

type MockGrpcClient interface {
	processor_grpc.DataProcessorClient
	GetClosedJpIds() []*processor_grpc.JobPostings
	GetJobPostingInfo() []*processor_grpc.JobPostingInfo
	GetCompany() []*processor_grpc.Company
}

type MockGrpcClientImpl struct {
	JobPostingInfos []*processor_grpc.JobPostingInfo
	JobPostings     []*processor_grpc.JobPostings
	Companies       []*processor_grpc.Company
}

func (m *MockGrpcClientImpl) CloseJobPostings(ctx context.Context, in *processor_grpc.JobPostings, opts ...grpc.CallOption) (*processor_grpc.BoolResponse, error) {
	m.JobPostings = append(m.JobPostings, in)
	return &processor_grpc.BoolResponse{Success: true}, nil
}

func (m *MockGrpcClientImpl) RegisterJobPostingInfo(ctx context.Context, in *processor_grpc.JobPostingInfo, opts ...grpc.CallOption) (*processor_grpc.BoolResponse, error) {
	m.JobPostingInfos = append(m.JobPostingInfos, in)
	return &processor_grpc.BoolResponse{Success: true}, nil
}

func (m *MockGrpcClientImpl) RegisterCompany(ctx context.Context, in *processor_grpc.Company, opts ...grpc.CallOption) (*processor_grpc.BoolResponse, error) {
	m.Companies = append(m.Companies, in)
	return &processor_grpc.BoolResponse{Success: true}, nil
}

func (m *MockGrpcClientImpl) GetClosedJpIds() []*processor_grpc.JobPostings {
	return m.JobPostings
}

func (m *MockGrpcClientImpl) GetJobPostingInfo() []*processor_grpc.JobPostingInfo {
	return m.JobPostingInfos
}

func (m *MockGrpcClientImpl) GetCompany() []*processor_grpc.Company {
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
		JobPostingInfos: make([]*processor_grpc.JobPostingInfo, 0),
		JobPostings:     make([]*processor_grpc.JobPostings, 0),
		Companies:       make([]*processor_grpc.Company, 0),
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
