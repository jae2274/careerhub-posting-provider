package jumpit

import (
	"careerhub-dataprovider/careerhub/provider/source"
	"careerhub-dataprovider/careerhub/provider/utils/apiactor"
)

type JumpitSource struct {
	ApiActor *apiactor.ApiActor
}

func NewJumpitSource(callDelay int64) *JumpitSource {
	return &JumpitSource{
		ApiActor: apiactor.NewApiActor(callDelay),
	}
}

func (s *JumpitSource) Site() string {
	return "jumpit"
}

func (s *JumpitSource) MaxPageSize() int {
	return 100
}

func (s *JumpitSource) List(page, size int) ([]source.JobPostingId, error) {
	return nil, nil //TODO
}

func (s *JumpitSource) Detail(jpId source.JobPostingId) (*source.JobPostingDetail, error) {
	return nil, nil //TODO
}

func (s *JumpitSource) Company(companyId string) (*source.Company, error) {
	return nil, nil //TODO
}
