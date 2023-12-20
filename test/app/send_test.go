package app

import (
	"careerhub-dataprovider/careerhub/provider/app"
	"careerhub-dataprovider/careerhub/provider/awscfg"
	"careerhub-dataprovider/careerhub/provider/domain/company"
	"careerhub-dataprovider/careerhub/provider/domain/jobposting"
	"careerhub-dataprovider/careerhub/provider/queue"
	"careerhub-dataprovider/careerhub/provider/queue/message_v1"
	"careerhub-dataprovider/careerhub/provider/source"
	"careerhub-dataprovider/careerhub/provider/source/jumpit"
	"careerhub-dataprovider/careerhub/provider/vars"
	"careerhub-dataprovider/test/tinit"
	"context"
	"encoding/base64"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/jae2274/goutils/cchan"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestSendJobPostingApp(t *testing.T) {
	t.Run("Run", func(t *testing.T) {
		src := jumpit.NewJumpitSource(2000)
		src.Run(make(<-chan app.QuitSignal))
		jobRepo, _, sendJobApp := initComponents(t, src)

		jpIds, err := src.List(1, 3)
		require.NoError(t, err)

		processedChan, errChan := sendJobApp.Run(jpIds, make(<-chan app.QuitSignal))

		require.NoError(t, err)

		cchan.WaitClosed(processedChan)
		if len(errChan) > 0 {
			for {
				select {
				case err := <-errChan:
					t.Log(err)
				default:
					t.Fail()
				}
			}
		}
		savedIds, err := jobRepo.GetAllHiring(src.Site())
		require.NoError(t, err)
		require.Len(t, savedIds, 3)

		messages := getFromJobPostingQueue(t)
		require.Len(t, messages, 3)

	Outer:
		for _, message := range messages {
			for _, jpId := range jpIds {
				if message.Site == jpId.Site && message.PostingId == jpId.PostingId {
					continue Outer
				}
			}
			t.Errorf("Not found %s %s", message.Site, message.PostingId)
			t.FailNow()
		}
	})

}

func initComponents(t *testing.T, src source.JobPostingSource) (*jobposting.JobPostingRepo, *company.CompanyRepo, *app.SendJobPostingApp) {
	envVars, err := vars.Variables()
	require.NoError(t, err)

	jobRepo := tinit.InitJobPostingRepo(t)
	companyRepo := tinit.InitCompanyRepo(t)
	jpQueue, _, _ := tinit.InitSQS(t, envVars.JobPostingQueue)
	companyQueue, _, _ := tinit.InitSQS(t, envVars.CompanyQueue)

	return jobRepo, companyRepo, app.NewSendJobPostingApp(src, jobRepo, companyRepo, queue.NewJobPostingQueue(jpQueue), queue.NewCompanyQueue(companyQueue))
}

func getFromJobPostingQueue(t *testing.T) []*message_v1.JobPostingInfo {
	envVars, err := vars.Variables()
	require.NoError(t, err)

	cfg, err := awscfg.Config()
	require.NoError(t, err)

	sqsClient := queue.NewClient(cfg, envVars.SqsEndpoint)
	result, err := sqsClient.GetQueueUrl(context.Background(), &sqs.GetQueueUrlInput{
		QueueName: &envVars.JobPostingQueue,
	})
	require.NoError(t, err)

	messages, err := getAll(sqsClient, result.QueueUrl)
	require.NoError(t, err)

	return messages
}

func getAll(sqsClient *sqs.Client, queueUrl *string) ([]*message_v1.JobPostingInfo, error) {
	messages := make([]*message_v1.JobPostingInfo, 0)

	for {
		result, err := sqsClient.ReceiveMessage(context.Background(),
			&sqs.ReceiveMessageInput{
				QueueUrl:            queueUrl,
				MaxNumberOfMessages: 10,
			},
		)
		if err != nil {
			return nil, err
		}

		if len(result.Messages) == 0 {
			break
		}

		for _, msg := range result.Messages {
			decodedBody, err := base64.StdEncoding.DecodeString(*msg.Body)
			if err != nil {
				return nil, err
			}

			var jobPostingInfo message_v1.JobPostingInfo
			err = proto.Unmarshal(decodedBody, &jobPostingInfo)
			if err != nil {
				return nil, err
			}

			messages = append(messages, &jobPostingInfo)
		}
	}

	return messages, nil
}
