package company

import (
	"careerhub-dataprovider/careerhub/provider/dynamo"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/jae2274/goutils/ptr"
)

type CompanyId struct {
	Site      string
	CompanyId string
}

type Company struct {
	Site      string            `dynamodbav:"site"`
	CompanyId string            `dynamodbav:"companyId"`
	CreatedAt dynamo.DynamoTime `dynamodbav:"createdAt"`
}

const (
	TableName      = "company"
	SiteField      = "site"
	CompanyIdField = "companyId"
)

func NewCompany(site, companyId string) *Company {
	return &Company{
		Site:      site,
		CompanyId: companyId,
	}
}

func (c Company) TableDef() dynamo.TableDefinition {
	siteFieldPtr := ptr.P(SiteField)
	companyIdFieldPtr := ptr.P(CompanyIdField)
	tableNamePtr := ptr.P(TableName)

	return dynamo.TableDefinition{
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: siteFieldPtr,
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: companyIdFieldPtr,
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: siteFieldPtr,
				KeyType:       types.KeyTypeRange,
			},
			{
				AttributeName: companyIdFieldPtr,
				KeyType:       types.KeyTypeHash,
			},
		},
		TableName: tableNamePtr,
	}
}
