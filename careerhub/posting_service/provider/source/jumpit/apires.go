package jumpit

import (
	"fmt"
	"time"

	"github.com/jae2274/careerhub-posting-provider/careerhub/posting_service/provider/domain/company"
	"github.com/jae2274/careerhub-posting-provider/careerhub/posting_service/provider/domain/jobposting"

	"github.com/jae2274/goutils/ptr"
	"github.com/jae2274/goutils/terr"
)

// list api result
type postingList struct {
	Result listResult `json:"result"`
	Status int        `json:"status"`
}

type listResult struct {
	TotalCount int `json:"totalCount"`
	Positions  []listPosition
}

type listPosition struct {
	Id               int    `json:"id"`
	JobCategory      string `json:"jobCategory"`
	CompanyProfileId int    `json:"companyProfileId"`
}

// detail api result
type postingDetail struct {
	Message string       `json:"message"`
	Status  int          `json:"status"`
	Code    string       `json:"code"`
	Result  detailResult `json:"result"`
}

type detailResult struct {
	ID                    int                 `json:"id"`
	Title                 string              `json:"title"`
	CompanyName           string              `json:"companyName"`
	Logo                  string              `json:"logo"`
	TechStacks            []techStack         `json:"techStacks"`
	ServiceInfo           string              `json:"serviceInfo"`
	Responsibility        string              `json:"responsibility"`
	Qualifications        string              `json:"qualifications"`
	PreferredRequirements string              `json:"preferredRequirements"`
	Welfares              string              `json:"welfares"`
	RecruitProcess        string              `json:"recruitProcess"`
	Newcomer              bool                `json:"newcomer"`
	MinCareer             *int32              `json:"minCareer,omitempty"`
	MaxCareer             *int32              `json:"maxCareer,omitempty"`
	PositionStatus        string              `json:"positionStatus"`
	DeveloperInterviews   []interface{}       `json:"developerInterviews"`
	ItCompanyStory        []interface{}       `json:"itCompanyStory"`
	PublishedAt           string              `json:"publishedAt"`
	ClosedAt              string              `json:"closedAt"`
	Location              *interface{}        `json:"location,omitempty"`
	Tags                  []tag               `json:"tags"`
	CompanyProfileID      int                 `json:"companyProfileId"` //will be deprecated
	EncodedSerialNumber   string              `json:"encodedSerialNumber"`
	CompanyURL            *string             `json:"companyUrl,omitempty"`
	AlwaysOpen            bool                `json:"alwaysOpen"`
	Celebration           int                 `json:"celebration"`
	Graduate              bool                `json:"graduate"`
	WorkingPlaces         []WorkingPlace      `json:"workingPlaces"`
	Follow                bool                `json:"follow"`
	Scrap                 bool                `json:"scrap"`
	Applied               bool                `json:"applied"`
	Draft                 bool                `json:"draft"`
	CacheCompanyImages    []cacheCompanyImage `json:"cacheCompanyImages"`
	JobCategories         []jobCategory       `json:"jobCategories"`
}

type techStack struct {
	Stack     string `json:"stack"`
	ImagePath string `json:"imagePath"`
}

type tag struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Emoticon string `json:"emoticon"`
}

type Address struct {
	AddressRegion  string `json:"addressRegion"`
	AddressCountry string `json:"addressCountry"`
	StreetAddress  string `json:"streetAddress"`
	Type           string `json:"@type"`
}

type WorkingPlace struct {
	Address    string `json:"address"`
	IsDomestic bool   `json:"isDomestic"`
}

type cacheCompanyImage struct {
	ImagePath  string `json:"imagePath"`
	SortNumber int    `json:"sortNumber"`
}

type jobCategory struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func convertSourceDetail(postingDetail *postingDetail, site, postUrl string) (*jobposting.JobPostingDetail, error) {
	result := postingDetail.Result

	publishedAt, err := time.Parse(time.DateTime, result.PublishedAt)
	if err != nil {
		return nil, terr.Wrap(err)
	}

	closedAt, err := time.Parse(time.DateTime, result.ClosedAt)
	if err != nil {
		return nil, terr.Wrap(err)
	}

	var imageUrl *string
	if len(result.CacheCompanyImages) > 0 {
		imageUrl = &(result.CacheCompanyImages[0].ImagePath)
	} else {
		imageUrl = nil
	}

	companyImages := make([]string, len(result.CacheCompanyImages))
	for i, image := range result.CacheCompanyImages {
		companyImages[i] = image.ImagePath
	}
	jobCategory := make([]string, len(result.JobCategories))
	for i, category := range result.JobCategories {
		jobCategory[i] = category.Name
	}

	dPosting := &jobposting.JobPostingDetail{
		Site:          site,
		PostingId:     fmt.Sprintf("%d", result.ID),
		CompanyId:     result.EncodedSerialNumber,
		CompanyName:   result.CompanyName,
		JobCategory:   jobCategory,
		ImageUrl:      imageUrl,
		CompanyImages: companyImages,
		MainContent: jobposting.MainContent{
			PostUrl:        postUrl,
			Title:          result.Title,
			Intro:          result.ServiceInfo,
			MainTask:       result.Responsibility,
			Qualifications: result.Qualifications,
			Preferred:      result.PreferredRequirements,
			Benefits:       result.Welfares,
			RecruitProcess: &result.RecruitProcess,
		},
		RequiredSkill: requiredSkill(result.TechStacks),
		Tags:          tags(result.Tags),
		RequiredCareer: jobposting.Career{
			Min: result.MinCareer,
			Max: result.MaxCareer,
		},
		PublishedAt: ptr.P(publishedAt.UnixMilli()),
		ClosedAt:    ptr.P(closedAt.UnixMilli()),
		Address:     address(result.WorkingPlaces),
	}
	return dPosting, checkValidPostingDetail(dPosting)
}

func checkValidPostingDetail(dPosting *jobposting.JobPostingDetail) error {
	invalidFields := make([]string, 0)

	if dPosting.Site == "" {
		invalidFields = append(invalidFields, "site")
	}

	if dPosting.PostingId == "" {
		invalidFields = append(invalidFields, "postingId")
	}

	if dPosting.CompanyId == "" {
		invalidFields = append(invalidFields, "companyId")
	}

	if dPosting.CompanyName == "" {
		invalidFields = append(invalidFields, "companyName")
	}

	if dPosting.MainContent.PostUrl == "" {
		invalidFields = append(invalidFields, "postUrl")
	}

	if dPosting.MainContent.Title == "" {
		invalidFields = append(invalidFields, "title")
	}

	if dPosting.MainContent.MainTask == "" {
		invalidFields = append(invalidFields, "mainTask")
	}

	if dPosting.MainContent.Qualifications == "" {
		invalidFields = append(invalidFields, "qualifications")
	}

	if dPosting.MainContent.Benefits == "" {
		invalidFields = append(invalidFields, "benefits")
	}

	if dPosting.Address == nil {
		invalidFields = append(invalidFields, "address")
	}

	if len(invalidFields) > 0 {
		return terr.New(fmt.Sprintf("site: %s, postingId: %s, empty fields: %v", dPosting.Site, dPosting.PostingId, invalidFields))
	}

	return nil
}

func requiredSkill(teckStack []techStack) []string {
	result := make([]string, len(teckStack))

	for i, stack := range teckStack {
		result[i] = stack.Stack
	}

	return result
}

func tags(tags []tag) []string {
	result := make([]string, len(tags))

	for i, tag := range tags {
		result[i] = tag.Name
	}

	return result
}

func address(workingPlaces []WorkingPlace) []string {
	result := make([]string, len(workingPlaces))

	for i, workingPlace := range workingPlaces {
		result[i] = workingPlace.Address
	}

	return result
}

// companyRes api result
type companyRes struct {
	Result companyResult `json:"result"`
}

type companyResult struct {
	Id                  int            `json:"id"`
	EncodedSerialNumber string         `json:"encodedSerialNumber"`
	CompanyName         string         `json:"companyName"`
	CompanySite         *string        `json:"companySite,omitempty"`
	CompanyService      companyService `json:"companyService"`
	CompanyLogo         string         `json:"companyLogo"`
	ProfileImages       []profileImage `json:"profileImages"`
}

type profileImage struct {
	ImagePath string `json:"imagePath"`
}

type companyService struct {
	Description string `json:"description"`
}

func convertSourceCompany(companyRes *companyRes, site string) (*company.CompanyDetail, error) {
	result := companyRes.Result

	companyImages := make([]string, len(result.ProfileImages))

	for i, image := range result.ProfileImages {
		companyImages[i] = image.ImagePath
	}

	dCompany := &company.CompanyDetail{
		Site:          site,
		CompanyId:     result.EncodedSerialNumber,
		Name:          result.CompanyName,
		CompanyUrl:    result.CompanySite,
		CompanyImages: companyImages,
		Description:   result.CompanyService.Description,
		CompanyLogo:   result.CompanyLogo,
	}
	return dCompany, checkValidCompany(dCompany)
}

func checkValidCompany(dCompany *company.CompanyDetail) error {
	invalidFields := make([]string, 0)

	if dCompany.Site == "" {
		invalidFields = append(invalidFields, "site")
	}
	if dCompany.CompanyId == "" {
		invalidFields = append(invalidFields, "companyId")
	}
	if dCompany.Name == "" {
		invalidFields = append(invalidFields, "name")
	}

	if len(invalidFields) > 0 {
		return terr.New(fmt.Sprintf("site: %s, companyId: %s, empty fields: %v", dCompany.Site, dCompany.CompanyId, invalidFields))
	}

	return nil
}
