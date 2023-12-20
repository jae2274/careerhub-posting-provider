package queue

import (
	"careerhub-dataprovider/careerhub/provider/vars"
	"careerhub-dataprovider/test/tinit"
	"testing"

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

			queue := tinit.InitSQS(t, queueName)

			err := queue.Send([]byte("test"))
			require.NoError(t, err)
			err = queue.Send([]byte("Hello, World!"))
			require.NoError(t, err)

			results, err := queue.Recv()
			require.NoError(t, err)

			require.Equal(t, 2, len(results))
			require.Equal(t, "test", string(results[0]))
			require.Equal(t, "Hello, World!", string(results[1]))
		}
	})
}
