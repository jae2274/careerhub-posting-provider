package jobposting

import (
	"context"

	"github.com/jae2274/goutils/enum"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type StateValues struct{}

type State = enum.Enum[StateValues]

const (
	Hiring = State("hiring")
	Closed = State("closed")
)

func (StateValues) Values() []string {
	return []string{
		string(Hiring),
		string(Closed),
	}
}

type JobPostingRepo struct {
	col *mongo.Collection
}

func NewJobPostingRepo(col *mongo.Collection) *JobPostingRepo {
	return &JobPostingRepo{
		col: col,
	}
}

func (jpr *JobPostingRepo) Get(id *JobPostingId) (*JobPosting, error) {
	var result JobPosting

	err := jpr.col.FindOne(context.Background(), bson.M{SiteField: id.Site, PostingIdField: id.PostingId}).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}

func (jpr *JobPostingRepo) Gets(ids []*JobPostingId) ([]*JobPosting, error) {
	var filters []bson.M
	for _, id := range ids {
		filter := bson.M{SiteField: id.Site, PostingIdField: id.PostingId}
		filters = append(filters, filter)
	}

	filter := bson.M{"$or": filters}

	cursor, err := jpr.col.Find(context.Background(), filter)

	if err != nil {
		return nil, err
	}

	var result []*JobPosting
	if err = cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (jpr *JobPostingRepo) Save(value *JobPosting) (*JobPosting, error) {
	_, err := jpr.col.InsertOne(context.Background(), value)

	if err != nil {
		return nil, err
	}

	return value, nil
}

func (jpr *JobPostingRepo) GetAllHiring(site string) ([]*JobPostingId, error) {
	filter := bson.M{SiteField: site}
	cursor, err := jpr.col.Find(context.Background(), filter)

	if err != nil {
		return nil, err
	}

	var result []*JobPostingId
	if err = cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (jpr *JobPostingRepo) DeleteAll(ids []*JobPostingId) error {
	var filters []bson.M
	for _, id := range ids {
		filter := bson.M{SiteField: id.Site, PostingIdField: id.PostingId}
		filters = append(filters, filter)
	}

	filter := bson.M{"$or": filters}

	_, err := jpr.col.DeleteMany(context.Background(), filter)

	return err
}
