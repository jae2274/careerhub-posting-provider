package jobposting

import (
	"careerhub-dataprovider/careerhub/provider/dynamo"
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/jae2274/goutils/enum"
	"github.com/jae2274/goutils/terr"
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

func (jpr *JobPostingRepo) GetAllHiring(site string) ([]*JobPostingId, error) {
	filtEx := expression.Name(SiteField).Equal(expression.Value(site))
	projEx := expression.NamesList(
		expression.Name(SiteField), expression.Name(PostingIdField))

	expr, err := expression.NewBuilder().WithFilter(filtEx).WithProjection(projEx).Build()

	if err != nil {
		return nil, terr.Wrap(err)
	}

	response, err := jpr.dbClient.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName:                 jpr.tableName,
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
	})

	if err != nil {
		return nil, terr.Wrap(err)
	}

	var jobPostingIds []*JobPostingId
	err = attributevalue.UnmarshalListOfMaps(response.Items, &jobPostingIds)
	if err != nil {
		return nil, terr.Wrap(err)
	}

	return jobPostingIds, nil
}

func (jpr *JobPostingRepo) DeleteAll(ids []*JobPostingId) error {
	if len(ids) == 0 {
		return nil
	}

	keys := make([]map[string]types.AttributeValue, len(ids))
	for i, id := range ids {
		keys[i] = newKey(id.Site, id.PostingId)
	}

	return dynamo.Deletes(jpr, context.TODO(), keys)
}

func (jpr *JobPostingRepo) DbClient() *dynamodb.Client {
	return jpr.dbClient
}

func newKey(site string, postingId string) map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		SiteField:      &types.AttributeValueMemberS{Value: site},
		PostingIdField: &types.AttributeValueMemberS{Value: postingId},
	}
}
