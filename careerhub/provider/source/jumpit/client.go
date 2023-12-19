package jumpit

import (
	"careerhub-dataprovider/careerhub/provider/app"
	"fmt"

	"github.com/jae2274/goutils/apiactor"
	"github.com/jae2274/goutils/jjson"
	"github.com/jae2274/goutils/terr"
)

type jumpitApiClient struct {
	aActor *apiactor.ApiActor
}

func newJumpitApiClient(callDelay int64) *jumpitApiClient {
	return &jumpitApiClient{
		aActor: apiactor.NewApiActor(callDelay),
	}
}

func (ja *jumpitApiClient) run(quitChan <-chan app.QuitSignal) {
	apiactor.Run(ja.aActor, quitChan)
}

func (ja *jumpitApiClient) List(page, size int) (*postingList, error) {
	request := apiactor.NewRequest(
		"GET",
		fmt.Sprintf("https://api.jumpit.co.kr/api/positions?page=%d&size=%d&sort=reg_dt&highlight=false", page, size),
	)

	return callApi[postingList](ja.aActor, request)
}

func (ja *jumpitApiClient) Detail(postingId string) (string, *postingDetail, error) {
	postingUrl := fmt.Sprintf("https://api.jumpit.co.kr/api/position/%s", postingId)
	request := apiactor.NewRequest(
		"GET",
		postingUrl,
	)

	result, err := callApi[postingDetail](ja.aActor, request)
	if err != nil {
		return "", nil, terr.Wrap(err)
	}

	return postingUrl, result, nil
}

func (ja *jumpitApiClient) Company(companyId string) (*company, error) {
	companyUrl := fmt.Sprintf("https://api.jumpit.co.kr/api/company/%s", companyId)
	request := apiactor.NewRequest(
		"GET",
		companyUrl,
	)
	response, err := callApi[company](ja.aActor, request)

	if err != nil {
		return nil, terr.Wrap(err)
	}

	return response, nil
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
