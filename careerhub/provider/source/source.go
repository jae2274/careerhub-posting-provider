package source

type JobPostingSource interface {
	Site() string
	MaxPageSize() int
	List(page, size int) ([]JobPostingId, error)
	Detail(JobPostingId) (*JobPostingDetail, error)
	Company(companyId string) (*Company, error)
}

type JobPostingId struct {
	PostingId string
	EtcInfo   map[string]string
}

type JobPostingDetail struct {
	Site           string
	PostingId      string
	CompanyId      string
	JobCategory    []string
	MainContent    MainContent
	RequiredSkill  []string
	Tags           []string
	RequiredCareer Career
	PublishedAt    *int64
	ClosedAt       *int64
	Address        []string
}

type MainContent struct {
	PostUrl        string
	Title          string
	Intro          string
	MainTask       string
	Qualifications string
	Preferred      string
	Benefits       string
	RecruitProcess *string
}

type Career struct {
	Min *uint
	Max *uint
}

type Company struct {
	Name          string
	CompanyUrl    *string
	CompanyImages []string
	Description   string
	CompanyLogo   string
}
