package dynamo

import (
	"careerhub-dataprovider/careerhub/provider/utils/terr"
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Repo[KEY any, VALUE Model] interface {
	Get(*KEY) (*VALUE, error)
	Save(*VALUE) (*VALUE, error)
	DbClient() *dynamodb.Client
}

func getFromRepo[KEY any, VALUE Model](r Repo[KEY, VALUE], context context.Context, value *VALUE) (*VALUE, error) {

	response, err := r.DbClient().GetItem(context, &dynamodb.GetItemInput{
		Key: (*value).GetKey(), TableName: (*value).TableDef().TableName,
	})

	if err != nil {
		return nil, terr.Wrap(err)
	} else if response.Item == nil {
		return nil, nil // Return nil when item is not found
	}

	err = attributevalue.UnmarshalMap(response.Item, value)
	if err != nil {
		return nil, err
	}

	return value, err
}

func saveFromRepo[KEY any, VALUE Model](r Repo[KEY, VALUE], context context.Context, value *VALUE) (*VALUE, error) {
	(*value).SetCreateAt(DynamoTime(time.Now()))

	item, err := attributevalue.MarshalMap(value)
	if err != nil {
		return nil, terr.Wrap(err)
	}

	_, err = r.DbClient().PutItem(context, &dynamodb.PutItemInput{
		TableName: (*value).TableDef().TableName, Item: item,
	})

	if err != nil {
		return nil, terr.Wrap(err)
	}

	return value, nil
}
