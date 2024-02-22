package wanted

import (
	"careerhub-dataprovider/careerhub/provider/source"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type wantedPostingList struct {
	Data []jobItem `json:"data"`
}

type jobItem struct {
	Id           int           `json:"id"`
	CategoryTags []CategoryTag `json:"category_tags"`
}

type CategoryTag struct {
	ParentID int `json:"parent_id"`
	ID       int `json:"id"`
}

type wantedPostingDetail struct {
	Job Job `json:"job"`
}

type Job struct {
	Address         Address        `json:"address"`
	IsCrossboarder  bool           `json:"is_crossboarder"`
	ID              int            `json:"id"`
	Detail          JobDetail      `json:"detail"`
	DueTime         *string        `json:"due_time,omitempty"`
	Score           float64        `json:"score"`
	CompanyImages   []CompanyImage `json:"company_images"`
	Hidden          bool           `json:"hidden"`
	SkillTags       []SkillTag     `json:"skill_tags"`
	Status          string         `json:"status"`
	IsBookmark      bool           `json:"is_bookmark"`
	Company         Company        `json:"company"`
	IsCompanyFollow bool           `json:"is_company_follow"`
	CompareCountry  bool           `json:"compare_country"`
	LogoImg         Image          `json:"logo_img"`
	CompanyTags     []CompanyTag   `json:"company_tags"`
	ShortLink       interface{}    `json:"short_link,omitempty"`
	TitleImg        Image          `json:"title_img"`
	Position        string         `json:"position"`
	CategoryTags    []CategoryTag  `json:"category_tags"`
}

type JobDetail struct {
	Requirements    string `json:"requirements"`
	MainTasks       string `json:"main_tasks"`
	Intro           string `json:"intro"`
	Benefits        string `json:"benefits"`
	PreferredPoints string `json:"preferred_points"`
}

type Address struct {
	Country      string       `json:"country"`
	FullLocation string       `json:"full_location"`
	GeoLocation  *GeoLocation `json:"geo_location,omitempty"`
	Location     string       `json:"location"`
	CountryCode  string       `json:"country_code"`
}

type GeoLocation struct {
	NLocation    NLocation   `json:"n_location"`
	Location     Location    `json:"location"`
	LocationType string      `json:"location_type"`
	Viewport     Viewport    `json:"viewport"`
	Bounds       interface{} `json:"bounds,omitempty"`
}

type NLocation struct {
	Lat     *float64 `json:"lat,omitempty"`
	Lng     *float64 `json:"lng,omitempty"`
	Address *string  `json:"address,omitempty"`
}

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Viewport struct {
	Northeast Coordinate `json:"northeast"`
	Southwest Coordinate `json:"southwest"`
}

type Coordinate struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type CompanyImage struct {
	URL string `json:"url"`
	ID  int    `json:"id"`
}

type SkillTag struct {
	Title     string `json:"title"`
	ID        int    `json:"id"`
	KindTitle string `json:"kind_title"`
}

type Company struct {
	ID           int    `json:"id"`
	IndustryName string `json:"industry_name"`
	Name         string `json:"name"`
}

type CompanyTag struct {
	Title     string `json:"title"`
	ID        int    `json:"id"`
	KindTitle string `json:"kind_title"`
}

type Image struct {
	Origin string `json:"origin"`
	Thumb  string `json:"thumb"`
}

func convertSourceDetail(detail *wantedPostingDetail, site string, postingUrl string, jobCategory string) (*source.JobPostingDetail, error) {
	job := detail.Job

	var closedAt int64
	if job.DueTime != nil {
		closedDate, _ := time.Parse("2006-01-02", *job.DueTime)
		closedAt = closedDate.UnixMilli()
	}

	var companyTags []string
	for _, tag := range job.CompanyTags {
		companyTags = append(companyTags, tag.Title)
	}

	requiredCareer, err := extractCareers(job)
	if err != nil {
		return nil, err
	}

	return &source.JobPostingDetail{
		Site:        site,
		PostingId:   fmt.Sprintf("%d", job.ID),
		CompanyId:   fmt.Sprintf("%d", job.Company.ID),
		CompanyName: job.Company.Name,
		JobCategory: strings.Split(jobCategory, ","),
		MainContent: source.MainContent{
			PostUrl:        postingUrl,
			Title:          job.Position,
			Intro:          job.Detail.Intro,
			MainTask:       job.Detail.MainTasks,
			Qualifications: job.Detail.Requirements,
			Preferred:      job.Detail.PreferredPoints,
			Benefits:       job.Detail.Benefits,
			RecruitProcess: nil,
		},
		RequiredSkill:  []string{}, //wanted의 skill에 대한 정보에 신뢰성이 없어 사용하지 않음
		Tags:           companyTags,
		RequiredCareer: requiredCareer,
		PublishedAt:    nil,
		ClosedAt:       &closedAt,
		Address:        []string{job.Address.FullLocation},
	}, nil
}

func extractCareers(job Job) (source.Career, error) {
	min, err := extractCareer(job, MIN)
	if err != nil {
		return source.Career{}, err
	}

	max, err := extractCareer(job, MAX)
	if err != nil {
		return source.Career{}, err
	}

	return source.Career{
		Min: min,
		Max: max,
	}, nil
}

func extractCareer(job Job, careerType CareerType) (*int32, error) {
	min, err := Career(job.Position, careerType)
	if err != nil {
		return nil, err
	} else if min == nil {
		min, err = Career(job.Detail.Requirements, careerType)
		if err != nil {
			return nil, err
		}
	}

	return min, nil
}

type CategoryJson struct {
	Props `json:"props"`
}

type Props struct {
	PageProps `json:"pageProps"`
}

type PageProps struct {
	Tags `json:"tags"`
}

type Tags struct {
	Categories []Category `json:"category"`
}

type Category struct {
	Id    int           `json:"id"`
	Tags  []SubCategory `json:"tags"`
	Title string        `json:"title"`
}

type SubCategory struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type wantedCompany struct {
	Company company `json:"company"`
}

type company struct {
	Id            int              `json:"id"`
	Name          string           `json:"name"`
	Detail        companyDetail    `json:"detail"`
	CompanyImages []companyImage   `json:"company_images"`
	LogoImg       companyLogoImage `json:"logo_img"`
}

type companyDetail struct {
	Description string  `json:"description"`
	Link        *string `json:"link"`
}

type companyImage struct {
	Url string `json:"url"`
}

type companyLogoImage struct {
	Origin string `json:"origin"`
}

func convertSourceCompany(company *wantedCompany, site string) *source.Company {
	result := company.Company

	companyImages := make([]string, len(result.CompanyImages))

	for i, image := range result.CompanyImages {
		companyImages[i] = image.Url
	}

	return &source.Company{
		Site:          site,
		CompanyId:     strconv.Itoa(result.Id),
		Name:          result.Name,
		CompanyUrl:    result.Detail.Link,
		CompanyImages: companyImages,
		Description:   result.Detail.Description,
		CompanyLogo:   result.LogoImg.Origin,
	}
}
