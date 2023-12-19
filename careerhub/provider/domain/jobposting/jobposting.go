package jobposting

import (
	"careerhub-dataprovider/careerhub/provider/dynamo"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/jae2274/goutils/ptr"
)

type JobPostingId struct {
	Site      string
	PostingId string
}
type JobPosting struct {
	Site      string            `dynamodbav:"site"`
	PostingId string            `dynamodbav:"postingId"`
	CreatedAt dynamo.DynamoTime `dynamodbav:"createdAt"`
}

func NewJobPosting(site string, postingId string) *JobPosting {
	return &JobPosting{
		Site:      site,
		PostingId: postingId,
	}
}

const (
	TableName      = "JobPosting"
	SiteField      = "site"
	PostingIdField = "postingId"
	StateField     = "state"
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

// func (jp JobPosting) GlobalSecondaryIndexes() []types.GlobalSecondaryIndex {
// 	secIndex := types.GlobalSecondaryIndex{
// 		IndexName: ptr.P("StateIndex"), // Set the name of the index
// 		KeySchema: []types.KeySchemaElement{
// 			{
// 				AttributeName: ptr.P(StateField), // Set the attribute name for the index
// 				KeyType:       types.KeyTypeRange,
// 			},
// 		},
// 		Projection: &types.Projection{
// 			ProjectionType: types.ProjectionTypeAll, // Set the projection type for the index
// 		},
// 	}

// 	return []types.GlobalSecondaryIndex{secIndex}
// }
