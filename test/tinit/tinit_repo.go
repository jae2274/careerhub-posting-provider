package tinit

import (
	"careerhub-dataprovider/careerhub/provider/domain/company"
	"careerhub-dataprovider/careerhub/provider/domain/jobposting"
	"testing"
)

func InitJobPostingRepo(t *testing.T) *jobposting.JobPostingRepo {
	dbClient := DB(t)

	jobpostingCollection := dbClient.Collection((&jobposting.JobPosting{}).Collection())
	repo := jobposting.NewJobPostingRepo(jobpostingCollection)

	return repo
}

func InitCompanyRepo(t *testing.T) *company.CompanyRepo {
	dbClient := DB(t)

	companyCollection := dbClient.Collection((&company.Company{}).Collection())
	repo := company.NewCompanyRepo(companyCollection)

	return repo
}
