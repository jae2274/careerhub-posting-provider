package company

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CompanyRepo struct {
	col *mongo.Collection
}

func NewCompanyRepo(dbClient *mongo.Collection) *CompanyRepo {
	return &CompanyRepo{
		col: dbClient,
	}
}

func (cr *CompanyRepo) Get(companyId *CompanyId) (*Company, error) {
	var result Company
	err := cr.col.FindOne(context.TODO(), bson.M{SiteField: companyId.Site, CompanyIdField: companyId.CompanyId}).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &result, err
}

func (cr *CompanyRepo) Gets(companyIds []*CompanyId) ([]*Company, error) {
	if len(companyIds) == 0 {
		return make([]*Company, 0), nil
	}

	// 회사 ID 목록에서 site와 companyId를 추출하여 검색 조건으로 사용
	var filters []bson.M
	for _, id := range companyIds {
		filter := bson.M{"site": id.Site, "companyId": id.CompanyId}
		filters = append(filters, filter)
	}

	// MongoDB $or 연산자를 사용하여 여러 조건 중 하나라도 만족하는 문서 검색
	filter := bson.M{"$or": filters}

	cursor, err := cr.col.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var results []*Company
	if err = cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (cr *CompanyRepo) GetAll() ([]*Company, error) {
	cursor, err := cr.col.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var result []*Company
	if err = cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (cr *CompanyRepo) Save(company *Company) (*Company, error) {
	company.CreatedAt = time.Now()
	_, err := cr.col.InsertOne(context.TODO(), company)
	return company, err
}
