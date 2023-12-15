package dynamo

import (
	"careerhub-dataprovider/careerhub/provider/dynamo"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestDynamoTime(t *testing.T) {
	now := time.Now()

	dynamoTime := dynamo.DynamoTime(now)

	av, err := dynamoTime.MarshalDynamoDBAttributeValue()
	require.NoError(t, err)

	var restoreDyTime dynamo.DynamoTime

	err = restoreDyTime.UnmarshalDynamoDBAttributeValue(av)
	require.NoError(t, err)

	require.Equal(t, now.UnixMilli(), time.Time(restoreDyTime).UnixMilli())
}
