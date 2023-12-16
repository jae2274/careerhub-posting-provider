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

// func NewJobPostingId(postingId string, etcInfo map[string]string) *JobPostingId {
// 	return &JobPostingId{
// 		PostingId: postingId,
// 		EtcInfo:   etcInfo,
// 	}
// }

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

// func NewJobPostingDetail(
// 	site, postingId, companyId string,
// 	jobCategory []string,
// 	mainContent MainContent,
// 	requiredSkill []string,
// 	tags []string,
// 	requiredCareer Career,
// 	publishedAt, closedAt *int64,
// 	address []string,
// ) *JobPostingDetail {
// 	return &JobPostingDetail{
// 		Site:           site,
// 		PostingId:      postingId,
// 		CompanyId:      companyId,
// 		JobCategory:    jobCategory,
// 		MainContent:    mainContent,
// 		RequiredSkill:  requiredSkill,
// 		Tags:           tags,
// 		RequiredCareer: requiredCareer,
// 		PublishedAt:    publishedAt,
// 		ClosedAt:       closedAt,
// 		Address:        address,
// 	}
// }

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

// func NewMainContent(
// 	postUrl, title, intro, mainTask, qualifications, preferred, benefits string,
// 	recruitProcess *string,
// ) *MainContent {
// 	return &MainContent{
// 		PostUrl:        postUrl,
// 		Title:          title,
// 		Intro:          intro,
// 		MainTask:       mainTask,
// 		Qualifications: qualifications,
// 		Preferred:      preferred,
// 		Benefits:       benefits,
// 		RecruitProcess: recruitProcess,
// 	}
// }

type Career struct {
	Min *int
	Max *int
}

// func NewCareer(min, max *int) *Career {
// 	return &Career{
// 		Min: min,
// 		Max: max,
// 	}
// }

type Company struct {
	Site          string
	CompanyId     string
	Name          string
	CompanyUrl    *string
	CompanyImages []string
	Description   string
	CompanyLogo   string
}

// func NewCompany(
// 	name string,
// 	companyUrl *string,
// 	companyImages []string,
// 	description string,
// 	companyLogo string,
// ) *Company {
// 	return &Company{
// 		Name:          name,
// 		CompanyUrl:    companyUrl,
// 		CompanyImages: companyImages,
// 		Description:   description,
// 		CompanyLogo:   companyLogo,
// 	}
// }
