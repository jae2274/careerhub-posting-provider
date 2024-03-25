package wanted

import (
	"fmt"
	"strconv"
	"time"

	"github.com/jae2274/careerhub-posting-provider/careerhub/posting_service/provider/domain/company"
	"github.com/jae2274/careerhub-posting-provider/careerhub/posting_service/provider/domain/jobposting"
)

type wantedPostingList struct {
	Data []jobItem `json:"data"`
}

type jobItem struct {
	Id int `json:"id"`
}

type wantedPostingDetail struct {
	Job Job `json:"job"`
}

type Job struct {
	ID          int         `json:"id"`
	Address     Address     `json:"address"`
	AnnualFrom  *int32      `json:"annual_from"`
	AnnualTo    *int32      `json:"annual_to"`
	CategoryTag CategoryTag `json:"category_tag"`
	Company     Company     `json:"company"`
	Detail      JobDetail   `json:"detail"`
	TitleImages []string    `json:"title_images"`
	DueTime     *string     `json:"due_time"`
}

type CategoryTag struct {
	ChildTags []ChildTag `json:"child_tags"`
}

type ChildTag struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
}

type JobDetail struct {
	Position        string `json:"position"`
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
	ID            int      `json:"id"`
	IndustryName  string   `json:"industry_name"`
	Name          string   `json:"name"`
	HighlightTags []string `json:"highlight_tags"`
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

const (
	maxTimeMillis = 253402300799000 //9999-12-31 23:59:59.000
)

func convertSourceDetail(detail *wantedPostingDetail, site string, postingUrl string) (*jobposting.JobPostingDetail, error) {
	job := detail.Job

	var closedAt int64
	if job.DueTime != nil {
		closedDate, _ := time.Parse("2006-01-02", *job.DueTime)
		closedAt = closedDate.UnixMilli()
	} else {
		closedAt = maxTimeMillis
	}

	var jobCategories []string
	for _, tag := range job.CategoryTag.ChildTags {
		jobCategories = append(jobCategories, tag.Text)
	}

	var imageUrl *string
	if len(job.TitleImages) > 0 {
		imageUrl = &(job.TitleImages[0])
	} else {
		imageUrl = nil
	}

	return &jobposting.JobPostingDetail{
		Site:          site,
		PostingId:     fmt.Sprintf("%d", job.ID),
		CompanyId:     fmt.Sprintf("%d", job.Company.ID),
		CompanyName:   job.Company.Name,
		JobCategory:   jobCategories,
		ImageUrl:      imageUrl,
		CompanyImages: job.TitleImages,
		MainContent: jobposting.MainContent{
			PostUrl:        postingUrl,
			Title:          job.Detail.Position,
			Intro:          job.Detail.Intro,
			MainTask:       job.Detail.MainTasks,
			Qualifications: job.Detail.Requirements,
			Preferred:      job.Detail.PreferredPoints,
			Benefits:       job.Detail.Benefits,
			RecruitProcess: nil,
		},
		RequiredSkill: []string{}, //wanted의 skill에 대한 정보에 신뢰성이 없어 사용하지 않음
		Tags:          job.Company.HighlightTags,
		RequiredCareer: jobposting.Career{
			Min: job.AnnualFrom,
			Max: job.AnnualTo,
		},
		PublishedAt: nil,
		ClosedAt:    &closedAt,
		Address:     []string{job.Address.FullLocation},
	}, nil
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
	Company companyRes `json:"company"`
}

type companyRes struct {
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

func convertSourceCompany(companyRes *wantedCompany, site string) *company.CompanyDetail {
	result := companyRes.Company

	companyImages := make([]string, len(result.CompanyImages))

	for i, image := range result.CompanyImages {
		companyImages[i] = image.Url
	}

	return &company.CompanyDetail{
		Site:          site,
		CompanyId:     strconv.Itoa(result.Id),
		Name:          result.Name,
		CompanyUrl:    result.Detail.Link,
		CompanyImages: companyImages,
		Description:   result.Detail.Description,
		CompanyLogo:   result.LogoImg.Origin,
	}
}
