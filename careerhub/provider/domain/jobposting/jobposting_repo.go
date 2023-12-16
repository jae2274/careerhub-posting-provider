package jobposting

import (
	"careerhub-dataprovider/careerhub/provider/dynamo"
	"careerhub-dataprovider/careerhub/provider/utils/enum"
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type StateValues struct{}

type State = enum.Enum[StateValues]

const (
	Hiring = State("hiring")
	Closed = State("closed")
)

func (StateValues) Values() []string {
	return []string{
		string(Hiring),
		string(Closed),
	}
}

type JobPostingRepo struct {
	dbClient  *dynamodb.Client
	tableName *string
}

func NewJobPostingRepo(dbClient *dynamodb.Client) (*JobPostingRepo, error) {
	jobPostingModel := JobPosting{}
	err := dynamo.CheckValidTable(dbClient, &jobPostingModel)

	if err != nil {
		fmt.Printf("Error checking match table: %v", err)
		return nil, err
	}

	return &JobPostingRepo{
		dbClient:  dbClient,
		tableName: jobPostingModel.TableDef().TableName,
	}, nil
}

func (jpr *JobPostingRepo) Get(id *JobPostingId) (*JobPosting, error) {
	return dynamo.Get(jpr, context.TODO(), newKey(id.Site, id.PostingId))
}

func (jpr *JobPostingRepo) Gets(ids []*JobPostingId) ([]*JobPosting, error) {
	keys := make([]map[string]types.AttributeValue, len(ids))
	for i, id := range ids {
		keys[i] = newKey(id.Site, id.PostingId)
	}

	return dynamo.Gets(jpr, context.TODO(), keys)
}

func (jpr *JobPostingRepo) Save(value *JobPosting) (*JobPosting, error) {
	value.CreatedAt = dynamo.DynamoTime(time.Now())
	return dynamo.Save(jpr, context.TODO(), value)
}

func (jpr *JobPostingRepo) DbClient() *dynamodb.Client {
	return jpr.dbClient
}

func newKey(site string, postingId string) map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"site":      &types.AttributeValueMemberS{Value: site},
		"postingId": &types.AttributeValueMemberS{Value: postingId},
	}
}
