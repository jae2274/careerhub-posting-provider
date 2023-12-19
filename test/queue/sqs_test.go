package queue

import (
	"careerhub-dataprovider/careerhub/provider/vars"
	"careerhub-dataprovider/test/tinit"
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/jae2274/goutils/ptr"
	"github.com/stretchr/testify/require"
)

func TestSQS(t *testing.T) {
	envVars, err := vars.Variables()
	require.NoError(t, err)

	queueNames := []string{
		envVars.JobPostingQueue,
		envVars.ClosedQueue,
		envVars.CompanyQueue,
	}

	t.Run("Send", func(t *testing.T) {
		for _, queueName := range queueNames {

			queue, sqsClient, queueUrl := tinit.InitSQS(t, queueName)

			queue.Send(ptr.P("test"))
			queue.Send(ptr.P("Hello, World!"))

			result, err := sqsClient.ReceiveMessage(context.Background(),
				&sqs.ReceiveMessageInput{
					QueueUrl:            queueUrl,
					MaxNumberOfMessages: 2,
				},
			)
			require.NoError(t, err)
			require.Equal(t, 2, len(result.Messages))
			require.Equal(t, "test", *result.Messages[0].Body)
			require.Equal(t, "Hello, World!", *result.Messages[1].Body)
		}
	})
}
