package tinit

import (
	"careerhub-dataprovider/careerhub/provider/domain/jobposting"
	"testing"
)

func InitJobPostingRepo(t *testing.T) *jobposting.JobPostingRepo {
	dbClient := DB(t)

	repo, err := jobposting.NewJobPostingRepo(dbClient)

	if err != nil {
		t.Errorf("Error creating summoner repo: %v", err)
		t.FailNow()
	}

	return repo
}
