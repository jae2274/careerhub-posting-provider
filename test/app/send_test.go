package app

import (
	"careerhub-dataprovider/careerhub/provider/app"
	"careerhub-dataprovider/careerhub/provider/domain/company"
	"careerhub-dataprovider/careerhub/provider/domain/jobposting"
	"careerhub-dataprovider/careerhub/provider/dynamo"
	"careerhub-dataprovider/careerhub/provider/processor_grpc"
	"careerhub-dataprovider/careerhub/provider/source"
	"careerhub-dataprovider/careerhub/provider/source/jumpit"
	"careerhub-dataprovider/test/tinit"
	"context"
	"testing"

	"github.com/jae2274/goutils/cchan"
	"github.com/stretchr/testify/require"
)

func TestSendJobPostingApp(t *testing.T) {
	t.Run("Run", func(t *testing.T) {
		ctx := context.Background()
		src := jumpit.NewJumpitSource(ctx, 3000)

		jobRepo, companyRepo, grpcClient, sendJobApp := initComponents(t, src)

		jpIds, err := src.List(1, 3)
		require.NoError(t, err)

		processedChan, errChan := sendJobApp.Run(ctx, jpIds)

		require.NoError(t, err)

		results := cchan.WaitClosed(processedChan)
		require.Len(t, results, len(jpIds))
		if len(errChan) > 0 {
			for {
				select {
				case err := <-errChan:
					t.Log(err)
				default:
					t.FailNow()
				}
			}
		}
		savedIds, err := jobRepo.GetAllHiring(src.Site())
		require.NoError(t, err)
		require.Len(t, savedIds, len(jpIds))

		jobPostingMessages := grpcClient.GetJobPostingInfo()

		IsEqualSrcJobPostingIds(t, jpIds, jobPostingMessages)
		IsEqualSavedJobPostingIds(t, jpIds, savedIds)

		savedCompanies, err := dynamo.GetAll(companyRepo, context.TODO())
		require.NoError(t, err)

		companyMessages := grpcClient.GetCompany()
		IsEqualSavedCompanies(t, savedCompanies, companyMessages)
		IsEqualJobPostingsAndCompanies(t, jobPostingMessages, companyMessages)
	})
}

func initComponents(t *testing.T, src source.JobPostingSource) (*jobposting.JobPostingRepo, *company.CompanyRepo, tinit.MockGrpcClient, *app.SendJobPostingApp) {
	jobRepo := tinit.InitJobPostingRepo(t)
	companyRepo := tinit.InitCompanyRepo(t)
	grpcClient := tinit.InitGrpcClient(t)

	return jobRepo, companyRepo, grpcClient, app.NewSendJobPostingApp(src, jobRepo, companyRepo, grpcClient)
}

func IsEqualSrcJobPostingIds(t *testing.T, srcJpIds []*source.JobPostingId, jobPostingMessages []*processor_grpc.JobPostingInfo) {
	require.Len(t, jobPostingMessages, len(srcJpIds))
Outer:
	for _, jobPostingMessage := range jobPostingMessages {
		for _, srcJpId := range srcJpIds {
			if jobPostingMessage.JobPostingId.Site == srcJpId.Site && jobPostingMessage.JobPostingId.PostingId == srcJpId.PostingId {
				continue Outer
			}
		}
		t.Errorf("Not found %s %s", jobPostingMessage.JobPostingId.Site, jobPostingMessage.JobPostingId.PostingId)
		t.FailNow()
	}
}

func IsEqualSavedJobPostingIds(t *testing.T, srcJpIds []*source.JobPostingId, savedJpIds []*jobposting.JobPostingId) {
	require.Len(t, savedJpIds, len(srcJpIds))
Outer:
	for _, message := range savedJpIds {
		for _, savedJpId := range srcJpIds {
			if message.Site == savedJpId.Site && message.PostingId == savedJpId.PostingId {
				continue Outer
			}
		}
		t.Errorf("Not found %s %s", message.Site, message.PostingId)
		t.FailNow()
	}
}

func IsEqualSavedCompanies(t *testing.T, savedCompanies []*company.Company, companyMessages []*processor_grpc.Company) {
	require.Len(t, savedCompanies, len(companyMessages))
Outer:
	for _, companyMessage := range companyMessages {
		for _, savedCompany := range savedCompanies {
			if companyMessage.Site == savedCompany.Site && companyMessage.CompanyId == savedCompany.CompanyId {
				continue Outer
			}
		}
		t.Errorf("Not found %s %s", companyMessage.Site, companyMessage.CompanyId)
		t.FailNow()
	}
}

func IsEqualJobPostingsAndCompanies(t *testing.T, jobPostingMessages []*processor_grpc.JobPostingInfo, companyMessages []*processor_grpc.Company) {
	jpCompany := make(map[string]interface{})
	for _, jobPosting := range jobPostingMessages {
		jpCompany[jobPosting.JobPostingId.Site+jobPosting.CompanyId] = false
	}

	for _, company := range companyMessages {
		if _, ok := jpCompany[company.Site+company.CompanyId]; !ok {
			require.Fail(t, "Not found %s %s", company.Site, company.CompanyId)
		}
	}
}
