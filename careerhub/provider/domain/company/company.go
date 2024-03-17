package company

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

type Company struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Site      string             `bson:"site"`
	CompanyId string             `bson:"companyId"`
	CreatedAt time.Time          `bson:"createdAt"`
}

func NewCompany(site, companyId string) *Company {
	return &Company{
		Site:      site,
		CompanyId: companyId,
	}
}

func (*Company) Collection() string {
	return "company"
}

func (*Company) IndexModels() map[string]*mongo.IndexModel {
	keyName := fmt.Sprintf("%s_1_%s_1", SiteField, CompanyIdField)
	return map[string]*mongo.IndexModel{
		keyName: {
			Keys: bson.D{
				{Key: SiteField, Value: 1},
				{Key: CompanyIdField, Value: 1},
			},
			Options: options.Index().SetUnique(true),
		},
	}
}
