package jobposting

import (
	"careerhub-dataprovider/careerhub/provider/dynamo"
	"careerhub-dataprovider/careerhub/provider/utils/ptr"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type JobPostingId struct {
	Site      string
	PostingId string
}
type JobPosting struct {
	Site      string             `dynamodbav:"site"`
	PostingId string             `dynamodbav:"postingId"`
	State     State              `dynamodbav:"state"`
	CreatedAt dynamo.DynamoTime  `dynamodbav:"createdAt"`
	UpdatedAt *dynamo.DynamoTime `dynamodbav:"updatedAt"`
}

func NewJobPosting(site string, postingId string) *JobPosting {
	return &JobPosting{
		Site:      site,
		PostingId: postingId,
		State:     Hiring,
	}
}

const (
	TableName      = "JobPosting"
	SiteField      = "site"
	PostingIdField = "postingId"
)

func (jp JobPosting) TableDef() dynamo.TableDefinition {
	siteFieldPtr := ptr.P(SiteField)
	postingIdFieldPtr := ptr.P(PostingIdField)
	tableNamePtr := ptr.P(TableName)

	return dynamo.TableDefinition{
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: siteFieldPtr,
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: postingIdFieldPtr,
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: siteFieldPtr,
				KeyType:       types.KeyTypeHash,
			},
			{
				AttributeName: postingIdFieldPtr,
				KeyType:       types.KeyTypeRange,
			},
		},
		TableName: tableNamePtr,
	}
}
