package company

import "careerhub-dataprovider/careerhub/provider/dynamo"

type CompanyId struct {
	Site      string
	CompanyId string
}

type Company struct {
	Site      string            `dynamodbav:"site"`
	CompanyId string            `dynamodbav:"companyId"`
	CreatedAt dynamo.DynamoTime `dynamodbav:"createdAt"`
}

func NewCompany(site, companyId string) *Company {
	return &Company{
		Site:      site,
		CompanyId: companyId,
	}
}
