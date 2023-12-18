package company

import (
	"careerhub-dataprovider/careerhub/provider/dynamo"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type CompanyRepo struct {
	dbClient  *dynamodb.Client
	tableName *string
}

func NewCompanyRepo(dbClient *dynamodb.Client) (*CompanyRepo, error) {
	jobPostingModel := Company{}
	err := dynamo.CheckValidTable(dbClient, &jobPostingModel)

	if err != nil {
		fmt.Printf("Error checking match table: %v", err)
		return nil, err
	}

	return &CompanyRepo{
		dbClient:  dbClient,
		tableName: jobPostingModel.TableDef().TableName,
	}, nil
}

func (cr *CompanyRepo) Get(companyId *CompanyId) (*Company, error) {
	return nil, nil
}

func (cr *CompanyRepo) Save(company *Company) (*Company, error) {
	return nil, nil
}
