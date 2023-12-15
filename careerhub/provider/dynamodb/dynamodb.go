package dynamodb

import (
	"careerhub-dataprovider/careerhub/provider/utils/terr"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func NewDbClient(cfg *aws.Config, endpoint *string) (*dynamodb.Client, error) {

	// Set the retrieved credentials to the AWS config

	// Create a new DynamoDB client with the updated config
	client := dynamodb.NewFromConfig(*cfg,
		func(options *dynamodb.Options) {
			if endpoint != nil {
				options.BaseEndpoint = endpoint
			}
		},
	)

	return client, nil
}

func CheckValidTable(dbClient *dynamodb.Client, model Model) error {
	tableDef := model.TableDef()
	desc, err := dbClient.DescribeTable(context.Background(), &dynamodb.DescribeTableInput{
		TableName: tableDef.TableName,
	})

	if err != nil {
		return terr.Wrap(err)
	}

	if !compareTableDef(tableDef, desc.Table) {
		fmt.Println("TableDef: ", tableDef)
		fmt.Println("Desc: ", desc.Table)
		return fmt.Errorf("table %s is not matched", *(tableDef.TableName))
	}
	return nil
}

func compareTableDef(a TableDefinition, b *types.TableDescription) bool {
	if !isAllAttributeEqual(a.AttributeDefinitions, b.AttributeDefinitions) {
		return false
	}
	if !isAllKeyEqual(a.KeySchema, b.KeySchema) {
		return false
	}
	return true
}

func isAllAttributeEqual(aAttrDef []types.AttributeDefinition, bAttrDef []types.AttributeDefinition) bool {
	if len(aAttrDef) != len(bAttrDef) {
		return false
	}
	for _, a := range aAttrDef {
		if !hasAttr(a, bAttrDef) {
			return false
		}
	}
	return true
}

func hasAttr(attr types.AttributeDefinition, attrDefs []types.AttributeDefinition) bool {
	for _, a := range attrDefs {
		if *attr.AttributeName == *a.AttributeName && attr.AttributeType == a.AttributeType {
			return true
		}
	}

	return false
}

func isAllKeyEqual(aKeySchema []types.KeySchemaElement, bKeySchema []types.KeySchemaElement) bool {
	if len(aKeySchema) != len(bKeySchema) {
		return false
	}
	for _, a := range aKeySchema {
		if !hasKey(a, bKeySchema) {
			return false
		}
	}
	return true
}

func hasKey(key types.KeySchemaElement, keySchemas []types.KeySchemaElement) bool {
	for _, a := range keySchemas {
		if *key.AttributeName == *a.AttributeName && key.KeyType == a.KeyType {
			return true
		}
	}
	return false
}
