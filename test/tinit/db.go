package tinit

import (
	awsconfig "careerhub-dataprovider/careerhub/provider/awscfg"
	"careerhub-dataprovider/careerhub/provider/domain/jobposting"
	"careerhub-dataprovider/careerhub/provider/dynamo"
	"careerhub-dataprovider/careerhub/provider/vars"
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func endpoint(t *testing.T) *string {
	variables, err := vars.Variables()
	if err != nil {
		t.Errorf(err.Error())
		t.FailNow()
	}
	return variables.DbEndpoint
}

func DB(t *testing.T) *dynamodb.Client {

	var endpoint = endpoint(t)
	awsconfig, err := awsconfig.Config()
	if err != nil {
		t.Errorf("Error creating aws config: %v", err)
		t.FailNow()
	}
	dbClient, err := dynamo.NewDbClient(awsconfig, endpoint)
	if err != nil {
		t.Errorf("Error creating db client: %v", err)
		t.FailNow()
	}

	truncateTable(t, dbClient, &jobposting.JobPosting{})
	return dbClient
}

func truncateTable(t *testing.T, dbClient *dynamodb.Client, model ...dynamo.Model) {
	deleteTable(t, dbClient, model...)
	checkDeleted(t, dbClient, model...)
	createTable(t, dbClient, model...)
}

func deleteTable(t *testing.T, dbClient *dynamodb.Client, model ...dynamo.Model) {
	for _, m := range model {
		_, err := dbClient.DeleteTable(context.Background(), &dynamodb.DeleteTableInput{
			TableName: m.TableDef().TableName,
		})

		if err != nil && !checkResourceNotFound(err) {
			t.Errorf("Error deleting table: %v", err)
			t.FailNow()
		}
	}
}

func checkDeleted(t *testing.T, dbClient *dynamodb.Client, model ...dynamo.Model) {
	for _, m := range model {
	Inner:
		for {
			_, err := dbClient.DescribeTable(context.Background(), &dynamodb.DescribeTableInput{
				TableName: m.TableDef().TableName,
			})

			if checkResourceNotFound(err) {
				// fmt.Printf("Table %s is deleted. err: %v\n", *m.TableDef().TableName, err)
				break Inner
			}

			// fmt.Printf("Waiting for table %s, err:%v\n", *m.TableDef().TableName, err)
			time.Sleep(1 * time.Second)
		}
	}
}

func createTable(t *testing.T, dbClient *dynamodb.Client, model ...dynamo.Model) {
	errorList := make([]error, 0)

	for _, m := range model {
		tableDef := m.TableDef()

		_, err := dbClient.CreateTable(context.Background(), &dynamodb.CreateTableInput{
			AttributeDefinitions: tableDef.AttributeDefinitions,
			KeySchema:            tableDef.KeySchema,
			TableName:            tableDef.TableName,
			BillingMode:          types.BillingModePayPerRequest,
		})

		if err != nil {
			fmt.Printf("Error creating table: %s\n", *tableDef.TableName)
			t.Errorf("Error creating table: %v", err)
			t.FailNow()
		}

		waiter := dynamodb.NewTableExistsWaiter(dbClient)

		err = waiter.Wait(context.Background(), &dynamodb.DescribeTableInput{
			TableName: tableDef.TableName}, 5*time.Minute)

		if err != nil {
			errorList = append(errorList, err)
		}
	}

	if len(errorList) > 0 {
		t.Errorf("Error creating tables: %v", errorList)
		t.FailNow()
	}
}

func checkResourceNotFound(err error) bool {
	return err != nil && strings.Contains(err.Error(), "ResourceNotFoundException")
}
