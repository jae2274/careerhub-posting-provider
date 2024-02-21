package wanted

import (
	"context"
	"fmt"

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
		fmt.Sprintf("https://www.wanted.co.kr/api/v4/jobs?tag_type_ids=%d&limit=%d&offset=%d&country=all&job_sort=job.latest_order&locations=all&years=-1", developmentTagId, offset, limit),
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
