package jumpit

import (
	"context"
	"fmt"

	"github.com/jae2274/careerhub-posting-provider/careerhub/posting_service/provider/source"
	"github.com/jae2274/goutils/apiactor"
	"github.com/jae2274/goutils/terr"
)

type jumpitApiClient struct {
	aActor *apiactor.ApiActor
}

func newJumpitApiClient(ctx context.Context, callDelay int64) *jumpitApiClient {
	return &jumpitApiClient{
		aActor: apiactor.NewApiActor(ctx, callDelay),
	}
}

func (ja *jumpitApiClient) List(page, size int) (*postingList, error) {
	request := apiactor.NewRequest(
		"GET",
		fmt.Sprintf("https://api.jumpit.co.kr/api/positions?page=%d&size=%d&sort=reg_dt&highlight=false", page, size),
	)

	return source.CallApi[postingList](ja.aActor, request)
}

func (ja *jumpitApiClient) Detail(postingId string) (string, *postingDetail, error) {
	postingUrl := fmt.Sprintf("https://api.jumpit.co.kr/api/position/%s", postingId)
	request := apiactor.NewRequest(
		"GET",
		postingUrl,
	)

	result, err := source.CallApi[postingDetail](ja.aActor, request)
	if err != nil {
		return "", nil, terr.Wrap(err)
	}

	return postingUrl, result, nil
}

func (ja *jumpitApiClient) Company(companyId string) (*companyRes, error) {
	companyUrl := fmt.Sprintf("https://api.jumpit.co.kr/api/company/%s", companyId)
	request := apiactor.NewRequest(
		"GET",
		companyUrl,
	)
	response, err := source.CallApi[companyRes](ja.aActor, request)

	if err != nil {
		return nil, terr.Wrap(err)
	}

	return response, nil
}
