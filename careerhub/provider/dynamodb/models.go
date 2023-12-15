package dynamodb

import "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

type TableDefinition struct {
	AttributeDefinitions []types.AttributeDefinition
	KeySchema            []types.KeySchemaElement
	TableName            *string
}

type Model interface {
	GetKey() map[string]types.AttributeValue
	TableDef() TableDefinition
}
