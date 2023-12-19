package queue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/jae2274/goutils/terr"
)

type SQS struct {
	client   *sqs.Client
	queueUrl string
}

func NewClient(cfg *aws.Config, endpoint *string) *sqs.Client {
	return sqs.NewFromConfig(*cfg,
		func(options *sqs.Options) {
			if endpoint != nil {
				options.BaseEndpoint = endpoint
			}
		},
	)
}

func NewSQS(cfg *aws.Config, endpoint *string, queueName string) (*SQS, error) {
	client := NewClient(cfg, endpoint)
	result, err := client.GetQueueUrl(context.Background(), &sqs.GetQueueUrlInput{
		QueueName: &queueName,
	})

	if err != nil {
		return nil, terr.Wrap(err)
	}

	return &SQS{
		client:   client,
		queueUrl: *result.QueueUrl,
	}, nil
}

func (q *SQS) Send(message *string) error {

	_, err := q.client.SendMessage(context.Background(), &sqs.SendMessageInput{
		MessageBody: message,
		QueueUrl:    &q.queueUrl,
	})

	if err != nil {
		return terr.Wrap(err)
	}

	return nil
}
