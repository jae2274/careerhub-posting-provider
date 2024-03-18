package company

const (
	IdField          = "_id"
	DefaultNameField = "defaultName"
	CompanyIdField   = "companyId"
	SiteField        = "site"
)

type CompanyId struct {
	Site      string `bson:"site"`
	CompanyId string `bson:"companyId"`
}

type CompanyDetail struct {
	Site          string `validate:"nonzero"`
	CompanyId     string `validate:"nonzero"`
	Name          string `validate:"nonzero"`
	CompanyUrl    *string
	CompanyImages []string
	Description   string `validate:"nonzero"`
	CompanyLogo   string `validate:"nonzero"`
}
