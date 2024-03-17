package jumpit

import (
	"careerhub-dataprovider/careerhub/provider/domain/company"
	"careerhub-dataprovider/careerhub/provider/domain/jobposting"
	"context"
	"fmt"
)

const (
	Site = "jumpit"
)

type JumpitSource struct {
	client jumpitApiClient
}

func NewJumpitSource(ctx context.Context, callDelayMilis int64) *JumpitSource {
	return &JumpitSource{
		client: *newJumpitApiClient(ctx, callDelayMilis),
	}
}

func (s *JumpitSource) Site() string {
	return Site
}

func (s *JumpitSource) MaxPageSize() int {
	return 300
}

func (js *JumpitSource) List(page, size int) ([]*jobposting.JobPostingId, error) {
	result, err := js.client.List(page, size)
	if err != nil {
		return nil, err
	}

	postingIds := make([]*jobposting.JobPostingId, len(result.Result.Positions))

	for i, position := range result.Result.Positions {
		postingIds[i] = &jobposting.JobPostingId{
			Site:      js.Site(),
			PostingId: fmt.Sprintf("%d", position.Id),
		}
	}

	return postingIds, nil
}

func (js *JumpitSource) Detail(jpId *jobposting.JobPostingId) (*jobposting.JobPostingDetail, error) {
	postingUrl, response, err := js.client.Detail(jpId.PostingId)
	if err != nil {
		return nil, err
	}

	return convertSourceDetail(response, js.Site(), postingUrl)
}

func (js *JumpitSource) Company(companyId string) (*company.CompanyDetail, error) {
	response, err := js.client.Company(companyId)
	if err != nil {
		return nil, err
	}

	return convertSourceCompany(response, js.Site()), nil
}
