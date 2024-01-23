package company

import (
	"careerhub-dataprovider/careerhub/provider/dynamo"
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type CompanyRepo struct {
	dbClient  *dynamodb.Client
	tableName *string
}

func NewCompanyRepo(dbClient *dynamodb.Client) (*CompanyRepo, error) {
	model := Company{}
	err := dynamo.CheckValidTable(dbClient, &model)

	if err != nil {
		return nil, err
	}

	return &CompanyRepo{
		dbClient:  dbClient,
		tableName: model.TableDef().TableName,
	}, nil
}

func (cr *CompanyRepo) Get(companyId *CompanyId) (*Company, error) {

	return dynamo.Get(cr, context.TODO(), newKey(companyId.Site, companyId.CompanyId))
}

func (cr *CompanyRepo) Gets(companyIds []*CompanyId) ([]*Company, error) {
	keys := make([]map[string]types.AttributeValue, len(companyIds))
	for i, id := range companyIds {
		keys[i] = newKey(id.Site, id.CompanyId)
	}

	return dynamo.Gets(cr, context.TODO(), keys)
}

func (cr *CompanyRepo) Save(company *Company) (*Company, error) {
	company.CreatedAt = dynamo.DynamoTime(time.Now())
	return dynamo.Save(cr, context.TODO(), company)
}

func (cr *CompanyRepo) DbClient() *dynamodb.Client {
	return cr.dbClient
}

func newKey(site string, companyId string) map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		SiteField:      &types.AttributeValueMemberS{Value: site},
		CompanyIdField: &types.AttributeValueMemberS{Value: companyId},
	}
}
