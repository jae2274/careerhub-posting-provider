package source

type JobPostingSource interface {
	Site() string
	MaxPageSize() int
	List(page, size int) ([]*JobPostingId, error)
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

func AllJobPostingIds(src JobPostingSource) ([]*JobPostingId, error) {
	maxPageSize := src.MaxPageSize()

	jobPostingIds := make([]*JobPostingId, maxPageSize*3)

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

	return jobPostingIds, nil
}
