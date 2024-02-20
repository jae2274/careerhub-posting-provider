package jobposting

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type JobPostingId struct {
	Site      string `bson:"site"`
	PostingId string `bson:"postingId"`
}
type JobPosting struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Site      string             `bson:"site"`
	PostingId string             `bson:"postingId"`
	CreatedAt time.Time          `bson:"createdAt"`
}

func NewJobPosting(site string, postingId string) *JobPosting {
	return &JobPosting{
		Site:      site,
		PostingId: postingId,
	}
}

const (
	IdField        = "_id"
	SiteField      = "site"
	PostingIdField = "postingId"
)

func (*JobPosting) Collection() string {
	return "JobPosting"
}

func (*JobPosting) IndexModels() map[string]*mongo.IndexModel {
	keyName := fmt.Sprintf("%s_1_%s_1", SiteField, PostingIdField)
	return map[string]*mongo.IndexModel{
		keyName: {
			Keys: bson.D{
				{Key: SiteField, Value: 1},
				{Key: PostingIdField, Value: 1},
			},
			Options: options.Index().SetUnique(true),
		},
	}
}
