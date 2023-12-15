package jumpit

import (
	"careerhub-dataprovider/careerhub/provider/source"
	"fmt"
	"log"
)

type JumpitSource struct {
	client jumpitApiClient
}

func NewJumpitSource(callDelay int64) *JumpitSource {
	return &JumpitSource{
		client: *newJumpitApiClient(callDelay),
	}
}

func (s *JumpitSource) Site() string {
	return "jumpit"
}

func (s *JumpitSource) MaxPageSize() int {
	return 100
}

func (js *JumpitSource) List(page, size int) ([]source.JobPostingId, error) {
	result, err := js.client.List(page, size)
	if err != nil {
		return nil, err
	}

	postingIds := make([]source.JobPostingId, len(result.Result.Positions))

	for i, position := range result.Result.Positions {
		postingIds[i] = *source.NewJobPostingId(
			fmt.Sprintf("%d", position.Id),
			map[string]string{"jobCategory": position.JobCategory},
		)
	}

	return postingIds, nil
}

func (js *JumpitSource) Detail(jpId source.JobPostingId) (*source.JobPostingDetail, error) {
	postingUrl, response, err := js.client.Detail(jpId.PostingId)
	if err != nil {
		return nil, err
	}

	jobCategory, ok := jpId.EtcInfo["jobCategory"]
	if !ok {
		log.Fatalf("jobCategory is not exist. site:%s, id: %v, etcInfo: %v", js.Site(), jpId.PostingId, jpId.EtcInfo)
	}
	return convertSourceDetail(response, js.Site(), postingUrl, jobCategory)
}

func (js *JumpitSource) Company(companyId string) (*source.Company, error) {
	response, err := js.client.Company(companyId)
	if err != nil {
		return nil, err
	}

	companyResult := response.Result
	return &source.Company{
		Name:        companyResult.CompanyName,
		CompanyUrl:  companyResult.CompanySite,
		CompanyLogo: companyResult.CompanyLogo,
		Description: companyResult.CompanyService.Description,
		// CompanyImages: TODO,
	}, nil
}
