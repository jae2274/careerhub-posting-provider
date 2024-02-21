package wanted

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/gocolly/colly/v2"
	"github.com/jae2274/goutils/apiactor"
	"github.com/jae2274/goutils/jjson"
	"github.com/jae2274/goutils/terr"
)

type wantedApiClient struct {
	aActor *apiactor.ApiActor
}

func newWantedApiClient(ctx context.Context, callDelay int64) *wantedApiClient {
	return &wantedApiClient{
		aActor: apiactor.NewApiActor(ctx, callDelay),
	}
}

const developmentTagId = 518

func (ja *wantedApiClient) List(offset, limit int) (*wantedPostingList, error) {
	request := apiactor.NewRequest(
		"GET",
		fmt.Sprintf("https://www.wanted.co.kr/api/v4/jobs?tag_type_ids=%d&limit=%d&offset=%d&country=all&job_sort=job.latest_order&locations=all&years=-1", developmentTagId, limit, offset),
	)

	return callApi[wantedPostingList](ja.aActor, request)
}

func (ja *wantedApiClient) Detail(postingId string) (string, *wantedPostingDetail, error) {
	postingUrl := fmt.Sprintf("https://www.wanted.co.kr/api/v4/jobs/%s", postingId)
	request := apiactor.NewRequest(
		"GET",
		postingUrl,
	)

	result, err := callApi[wantedPostingDetail](ja.aActor, request)

	if err != nil {
		return "", nil, terr.Wrap(err)
	}

	return postingUrl, result, nil
}

func (ja *wantedApiClient) Company(companyId string) (*wantedCompany, error) {
	request := apiactor.NewRequest(
		"GET",
		fmt.Sprintf("https://www.wanted.co.kr/api/v4/companies/%s", companyId),
	)

	return callApi[wantedCompany](ja.aActor, request)
}

func callApi[RESULT any](aActor *apiactor.ApiActor, request *apiactor.Request) (*RESULT, error) {
	rc, err := aActor.Call(request)
	if err != nil {
		return nil, terr.Wrap(err)
	}

	result, err := jjson.UnmarshalReader[RESULT](rc)
	if err != nil {
		return nil, terr.Wrap(err)
	}

	return result, nil
}

func getJobCategoryJson() (*CategoryJson, error) {
	c := colly.NewCollector()

	wg := sync.WaitGroup{}
	wg.Add(1)

	var categoryJson CategoryJson
	var err error
	c.OnHTML("#__NEXT_DATA__", func(e *colly.HTMLElement) {
		defer wg.Done()

		err = json.Unmarshal([]byte(e.Text), &categoryJson)
	})

	c.Visit("https://www.wanted.co.kr/wdlist")
	wg.Wait()
	return &categoryJson, err
}

func GetJobCategoryMap() (map[int]string, error) {
	categoryJson, err := getJobCategoryJson()
	if err != nil {
		return nil, terr.Wrap(err)
	}

	var developmentCategory *Category
	for _, category := range categoryJson.Props.PageProps.Tags.Categories {
		if category.Id == developmentTagId {
			developmentCategory = &category
			break
		}
	}
	if developmentCategory == nil {
		return nil, terr.New("development category not found")
	}

	categoryMap := make(map[int]string)

	for _, subCategory := range developmentCategory.Tags {
		categoryMap[subCategory.Id] = subCategory.Title
	}

	return categoryMap, nil
}
