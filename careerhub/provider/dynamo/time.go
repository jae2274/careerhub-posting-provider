package dynamo

import (
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DynamoTime time.Time

func (e DynamoTime) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	millis := timeAsMillis(time.Time(e))

	return &types.AttributeValueMemberN{
		Value: strconv.FormatInt(millis, 10),
	}, nil
}

func (e *DynamoTime) UnmarshalDynamoDBAttributeValue(av types.AttributeValue) error {
	avN, ok := av.(*types.AttributeValueMemberN)
	if !ok {
		return nil
	}

	n, err := strconv.ParseInt(avN.Value, 10, 64)
	if err != nil {
		return err
	}

	*e = DynamoTime(millisAsTime(n))
	return nil
}

func timeAsMillis(t time.Time) int64 {
	return t.UnixMilli()
}

func millisAsTime(millis int64) time.Time {

	return time.UnixMilli(millis)
}
