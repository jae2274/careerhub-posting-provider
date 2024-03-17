package wanted

import (
	"careerhub-dataprovider/careerhub/provider/domain/company"
	"careerhub-dataprovider/careerhub/provider/domain/jobposting"
	"context"
	"fmt"
)

const (
	Site = "wanted"
)

type WantedSource struct {
	client wantedApiClient
}

func NewWantedSource(ctx context.Context, callDelayMilis int64) *WantedSource {
	client := *newWantedApiClient(ctx, callDelayMilis)

	return &WantedSource{
		client: client,
	}
}

func (js *WantedSource) Site() string {
	return Site
}
func (js *WantedSource) MaxPageSize() int {
	return 100 //TODO: implement
}

func (js *WantedSource) List(page, size int) ([]*jobposting.JobPostingId, error) { //가장 최신의 채용공고부터 정렬되도록 반환
	result, err := js.client.List((page-1)*size, size)
	if err != nil {
		return nil, err
	}

	postingIds := make([]*jobposting.JobPostingId, len(result.Data))

	for i, position := range result.Data {

		postingIds[i] = &jobposting.JobPostingId{
			Site:      js.Site(),
			PostingId: fmt.Sprintf("%d", position.Id),
		}
	}

	return postingIds, nil
}

func (js *WantedSource) Detail(jpId *jobposting.JobPostingId) (*jobposting.JobPostingDetail, error) {
	postingUrl, response, err := js.client.Detail(jpId.PostingId)
	if err != nil {
		return nil, err
	}

	return convertSourceDetail(response, js.Site(), postingUrl)
}
func (js *WantedSource) Company(companyId string) (*company.CompanyDetail, error) {

	response, err := js.client.Company(companyId)
	if err != nil {
		return nil, err
	}

	return convertSourceCompany(response, js.Site()), nil
}
