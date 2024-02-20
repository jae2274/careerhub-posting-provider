package tinit

import (
	"careerhub-dataprovider/careerhub/provider/domain/company"
	"careerhub-dataprovider/careerhub/provider/domain/jobposting"
	"careerhub-dataprovider/careerhub/provider/mongocfg"
	"careerhub-dataprovider/careerhub/provider/vars"
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DB(t *testing.T) *mongo.Database {
	envVars, err := vars.Variables()
	checkError(t, err)

	db, err := mongocfg.NewDatabase(envVars.MongoUri, envVars.DbName, envVars.DBUser)
	checkError(t, err)

	jpModel := &jobposting.JobPosting{}
	jpCol := db.Collection(jpModel.Collection())
	err = jpCol.Drop(context.TODO())
	checkError(t, err)
	createIndexes(t, jpCol, jpModel.IndexModels())

	companyModel := &company.Company{}
	companyCol := db.Collection(companyModel.Collection())
	err = companyCol.Drop(context.TODO())
	checkError(t, err)
	createIndexes(t, companyCol, companyModel.IndexModels())

	return db
}

func createIndexes(t *testing.T, col *mongo.Collection, indexModels map[string]*mongo.IndexModel) {
	for indexName, indexModel := range indexModels {
		if indexModel.Options == nil {
			indexModel.Options = options.Index()
		}
		indexModel.Options.Name = &indexName

		_, err := col.Indexes().CreateOne(context.TODO(), *indexModel, nil)
		checkError(t, err)
	}
}
