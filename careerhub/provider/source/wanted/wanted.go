package wanted

import (
	"careerhub-dataprovider/careerhub/provider/source"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/jae2274/goutils/llog"
	"github.com/jae2274/goutils/terr"
)

type WantedSource struct {
	client           wantedApiClient
	categoryIdToName map[int]string //TODO: implement
}

func NewWantedSource(ctx context.Context, callDelayMilis int64) (*WantedSource, error) {
	client := *newWantedApiClient(ctx, callDelayMilis)
	jobCategory, err := GetJobCategoryMap()

	if err != nil {
		return nil, err
	}

	return &WantedSource{
		client:           client,
		categoryIdToName: jobCategory,
	}, nil
}

func (js *WantedSource) Site() string {
	return "wanted"
}
func (js *WantedSource) MaxPageSize() int {
	return 100 //TODO: implement
}

func (js *WantedSource) List(page, size int) ([]*source.JobPostingId, error) { //가장 최신의 채용공고부터 정렬되도록 반환
	result, err := js.client.List((page-1)*size, size)
	if err != nil {
		return nil, err
	}

	postingIds := make([]*source.JobPostingId, len(result.Data))

	for i, position := range result.Data {

		var jobCategories []string
		for _, categoryTag := range position.CategoryTags {
			jobCategory, ok := js.categoryIdToName[categoryTag.ID]
			if !ok {
				return nil, terr.New(fmt.Sprintf("categoryIdToName is not exist. categoryId:%d", categoryTag.ID))
			}
			jobCategories = append(jobCategories, jobCategory)
		}

		postingIds[i] = &source.JobPostingId{
			Site:      js.Site(),
			PostingId: fmt.Sprintf("%d", position.Id),
			EtcInfo:   map[string]string{"jobCategory": strings.Join(jobCategories, ",")},
		}
	}

	return postingIds, nil
}

func (js *WantedSource) Detail(jpId *source.JobPostingId) (*source.JobPostingDetail, error) {
	jobCategory, ok := jpId.EtcInfo["jobCategory"]
	if !ok {
		llog.Error(context.Background(), fmt.Sprintf("jobCategory is not exist. site:%s, id: %v, etcInfo: %v", js.Site(), jpId.PostingId, jpId.EtcInfo))
		os.Exit(1)
	}

	postingUrl, response, err := js.client.Detail(jpId.PostingId)
	if err != nil {
		return nil, err
	}

	return convertSourceDetail(response, js.Site(), postingUrl, jobCategory)
}
func (js *WantedSource) Company(companyId string) (*source.Company, error) {

	response, err := js.client.Company(companyId)
	if err != nil {
		return nil, err
	}

	return convertSourceCompany(response, js.Site()), nil
}
