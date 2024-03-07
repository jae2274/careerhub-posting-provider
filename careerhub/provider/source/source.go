package source

import "slices"

type Page struct {
	Size int
	Page int
}
type JobPostingSource interface {
	Site() string
	MaxPageSize() int
	List(page, size int) ([]*JobPostingId, error) //가장 최신의 채용공고부터 정렬되도록 반환
	Detail(*JobPostingId) (*JobPostingDetail, error)
	Company(companyId string) (*Company, error)
}

type JobPostingId struct {
	Site      string
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
	RequiredSkill  []string
	Tags           []string
	RequiredCareer Career `validate:"nonzero"`
	PublishedAt    *int64
	ClosedAt       *int64
	Address        []string `validate:"nonzero"`
	ImageUrl       *string
	CompanyImages  []string
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
	Min *int32
	Max *int32
}

type Company struct {
	Site          string `validate:"nonzero"`
	CompanyId     string `validate:"nonzero"`
	Name          string `validate:"nonzero"`
	CompanyUrl    *string
	CompanyImages []string
	Description   string `validate:"nonzero"`
	CompanyLogo   string `validate:"nonzero"`
}

// src에서 모든 채용공고id를 가져온다.
// 가장 오래된 채용공고부터 정렬되도록 반환되어야 한다.
func AllJobPostingIds(src JobPostingSource) ([]*JobPostingId, error) {
	maxPageSize := src.MaxPageSize()

	jobPostingIds := make([]*JobPostingId, 0)

	page := 0
	alreadyListed := make(map[string]bool)
	for {
		page++
		ids, err := src.List(page, maxPageSize)
		if err != nil {
			return nil, err
		}

		if len(ids) == 0 {
			break
		}

		for _, id := range ids {
			if _, ok := alreadyListed[id.PostingId]; !ok {
				alreadyListed[id.PostingId] = true
				jobPostingIds = append(jobPostingIds, id)
			}
		}
	}
	slices.Reverse(jobPostingIds)
	return jobPostingIds, nil
}
