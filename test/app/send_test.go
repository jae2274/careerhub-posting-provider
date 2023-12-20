package app

import (
	"careerhub-dataprovider/careerhub/provider/app"
	"careerhub-dataprovider/careerhub/provider/domain/company"
	"careerhub-dataprovider/careerhub/provider/domain/jobposting"
	"careerhub-dataprovider/careerhub/provider/dynamo"
	"careerhub-dataprovider/careerhub/provider/queue"
	"careerhub-dataprovider/careerhub/provider/queue/message_v1"
	"careerhub-dataprovider/careerhub/provider/source"
	"careerhub-dataprovider/careerhub/provider/source/jumpit"
	"careerhub-dataprovider/careerhub/provider/vars"
	"careerhub-dataprovider/test/tinit"
	"context"
	"testing"

	"github.com/jae2274/goutils/cchan"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestSendJobPostingApp(t *testing.T) {
	t.Run("Run", func(t *testing.T) {
		src := jumpit.NewJumpitSource(2000)
		src.Run(make(<-chan app.QuitSignal))
		jobRepo, companyRepo, jpQ, companyQ, sendJobApp := initComponents(t, src)

		jpIds, err := src.List(1, 3)
		require.NoError(t, err)

		processedChan, errChan := sendJobApp.Run(jpIds, make(<-chan app.QuitSignal))

		require.NoError(t, err)

		results := cchan.WaitClosed(processedChan)
		require.Len(t, results, 3)
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
		require.Len(t, savedIds, 3)

		jobPostingMessages := getJobPostingMessages(t, jpQ)

		IsEqualSrcJobPostingIds(t, jpIds, jobPostingMessages)
		IsEqualSavedJobPostingIds(t, jpIds, savedIds)

		savedCompanies, err := dynamo.GetAll(companyRepo, context.TODO())
		require.NoError(t, err)

		companyMessages := getCompanyMessages(t, companyQ)
		IsEqualSavedCompanies(t, savedCompanies, companyMessages)
	})
}

func getJobPostingMessages(t *testing.T, jpQ queue.Queue) []message_v1.JobPostingInfo {
	messages, err := jpQ.Recv()
	require.NoError(t, err)

	jobPostingMessages := make([]message_v1.JobPostingInfo, len(messages))
	for i, message := range messages {
		err := proto.Unmarshal(message, &jobPostingMessages[i])
		require.NoError(t, err)
	}
	return jobPostingMessages
}

func initComponents(t *testing.T, src source.JobPostingSource) (*jobposting.JobPostingRepo, *company.CompanyRepo, queue.Queue, queue.Queue, *app.SendJobPostingApp) {
	envVars, err := vars.Variables()
	require.NoError(t, err)

	jobRepo := tinit.InitJobPostingRepo(t)
	companyRepo := tinit.InitCompanyRepo(t)
	jpQueue := tinit.InitSQS(t, envVars.JobPostingQueue)
	companyQueue := tinit.InitSQS(t, envVars.CompanyQueue)

	return jobRepo, companyRepo, jpQueue, companyQueue, app.NewSendJobPostingApp(src, jobRepo, companyRepo, queue.NewJobPostingQueue(jpQueue), queue.NewCompanyQueue(companyQueue))
}

func IsEqualSrcJobPostingIds(t *testing.T, srcJpIds []*source.JobPostingId, jobPostingMessages []message_v1.JobPostingInfo) {
	require.Len(t, jobPostingMessages, len(srcJpIds))
Outer:
	for _, jobPostingMessage := range jobPostingMessages {
		for _, srcJpId := range srcJpIds {
			if jobPostingMessage.Site == srcJpId.Site && jobPostingMessage.PostingId == srcJpId.PostingId {
				continue Outer
			}
		}
		t.Errorf("Not found %s %s", jobPostingMessage.Site, jobPostingMessage.PostingId)
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

func getCompanyMessages(t *testing.T, companyQ queue.Queue) []message_v1.Company {
	messages, err := companyQ.Recv()
	require.NoError(t, err)

	companyMessages := make([]message_v1.Company, len(messages))
	for i, message := range messages {
		err := proto.Unmarshal(message, &companyMessages[i])
		require.NoError(t, err)
	}
	return companyMessages
}

func IsEqualSavedCompanies(t *testing.T, savedCompanies []*company.Company, companyMessages []message_v1.Company) {
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
