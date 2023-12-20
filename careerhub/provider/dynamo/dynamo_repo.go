package dynamo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/jae2274/goutils/terr"
)

type Repo[KEY any, VALUE Model] interface {
	Get(*KEY) (*VALUE, error)
	Gets([]*KEY) ([]*VALUE, error)
	Save(*VALUE) (*VALUE, error)
	DbClient() *dynamodb.Client
}

func Get[KEY any, VALUE Model](r Repo[KEY, VALUE], context context.Context, key map[string]types.AttributeValue) (*VALUE, error) {

	model := new(VALUE)
	response, err := r.DbClient().GetItem(context, &dynamodb.GetItemInput{
		Key: key, TableName: (*model).TableDef().TableName,
	})

	if err != nil {
		return nil, terr.Wrap(err)
	} else if response.Item == nil {
		return nil, nil // Return nil when item is not found
	}

	err = attributevalue.UnmarshalMap(response.Item, model)
	if err != nil {
		return nil, terr.Wrap(err)
	}

	return model, err
}

func Save[KEY any, VALUE Model](r Repo[KEY, VALUE], context context.Context, value *VALUE) (*VALUE, error) {
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

func Gets[KEY any, VALUE Model](r Repo[KEY, VALUE], context context.Context, keys []map[string]types.AttributeValue) ([]*VALUE, error) {
	model := new(VALUE)
	tableName := *(*model).TableDef().TableName

	response, err := r.DbClient().BatchGetItem(context, &dynamodb.BatchGetItemInput{
		RequestItems: map[string]types.KeysAndAttributes{
			tableName: {
				Keys: keys,
			},
		},
	})

	if err != nil {
		return nil, terr.Wrap(err)
	}

	result := make([]*VALUE, len(response.Responses[tableName]))

	for i, item := range response.Responses[tableName] {
		value := new(VALUE)
		err = attributevalue.UnmarshalMap(item, value)
		if err != nil {
			return nil, terr.Wrap(err)
		}

		result[i] = value
	}

	return result, nil
}

func GetAll[KEY any, VALUE Model](r Repo[KEY, VALUE], context context.Context) ([]*VALUE, error) {
	model := new(VALUE)

	response, err := r.DbClient().Scan(context, &dynamodb.ScanInput{
		TableName: (*model).TableDef().TableName,
	})

	if err != nil {
		return nil, terr.Wrap(err)
	}

	result := make([]*VALUE, len(response.Items))

	for i, item := range response.Items {
		value := new(VALUE)
		err = attributevalue.UnmarshalMap(item, value)
		if err != nil {
			return nil, terr.Wrap(err)
		}

		result[i] = value
	}

	return result, nil
}
