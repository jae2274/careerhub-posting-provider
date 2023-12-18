package source

import "slices"

type Page struct {
	Size int
	Page int
}
type JobPostingSource interface {
	Site() string
	MaxPageSize() int
	LastPage() (*Page, error)
	List(page, size int) ([]*JobPostingId, error) //가장 최신의 채용공고부터 정렬되도록 반환
	Detail(*JobPostingId) (*JobPostingDetail, error)
	Company(companyId string) (*Company, error)
}

type JobPostingId struct {
	PostingId string
	EtcInfo   map[string]string
}

type JobPostingDetail struct {
	Site           string      `validate:"nonzero"`
	PostingId      string      `validate:"nonzero"`
	CompanyId      string      `validate:"nonzero"`
	CompanyName    string      `validate:"nonzero"`
	JobCategory    []string    `validate:"nonzero"`
	MainContent    MainContent `validate:"nonzero"`
	RequiredSkill  []string    `validate:"nonzero"`
	Tags           []string    `validate:"nonzero"`
	RequiredCareer Career      `validate:"nonzero"`
	PublishedAt    *int64
	ClosedAt       *int64
	Address        []string `validate:"nonzero"`
}

type MainContent struct {
	PostUrl        string `validate:"nonzero"`
	Title          string `validate:"nonzero"`
	Intro          string `validate:"nonzero"`
	MainTask       string `validate:"nonzero"`
	Qualifications string `validate:"nonzero"`
	Preferred      string `validate:"nonzero"`
	Benefits       string `validate:"nonzero"`
	RecruitProcess *string
}

type Career struct {
	Min *int
	Max *int
}

type Company struct {
	Site          string `validate:"nonzero"`
	CompanyId     string `validate:"nonzero"`
	Name          string `validate:"nonzero"`
	CompanyUrl    *string
	CompanyImages []string `validate:"nonzero"`
	Description   string   `validate:"nonzero"`
	CompanyLogo   string   `validate:"nonzero"`
}

// src에서 모든 채용공고id를 가져온다.
// 가장 오래된 채용공고부터 정렬되도록 반환되어야 한다.
func AllJobPostingIds(src JobPostingSource) ([]*JobPostingId, error) {
	maxPageSize := src.MaxPageSize()

	jobPostingIds := make([]*JobPostingId, 0)

	page := 0
	for {
		page++
		ids, err := src.List(page, maxPageSize)
		if err != nil {
			return nil, err
		}

		if len(ids) == 0 {
			break
		}

		jobPostingIds = append(jobPostingIds, ids...)
	}
	slices.Reverse(jobPostingIds)
	return jobPostingIds, nil
}
