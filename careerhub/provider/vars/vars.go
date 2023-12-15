package vars

import (
	"fmt"
	"os"
)

type Vars struct {
	DbEndpoint  *string
	SqsEndpoint *string
	QueueName   string
}

type ErrNotExistedVar struct {
	VarName string
}

func NotExistedVar(varName string) *ErrNotExistedVar {
	return &ErrNotExistedVar{VarName: varName}
}

func (e *ErrNotExistedVar) Error() string {
	return fmt.Sprintf("%s is not existed", e.VarName)
}

func Variables() (*Vars, error) {

	dbEndpoint := getFromEnvPtr("DB_ENDPOINT")
	sqsEndpoint := getFromEnvPtr("SQS_ENDPOINT")

	queueName, err := getFromEnv("QUEUE_NAME")
	if err != nil {
		return nil, err
	}

	return &Vars{
		DbEndpoint:  dbEndpoint,
		SqsEndpoint: sqsEndpoint,
		QueueName:   queueName,
	}, nil
}

func getFromEnv(envVar string) (string, error) {
	ev := os.Getenv(envVar)

	if ev == "" {
		return "", fmt.Errorf("%s is not existed", envVar)
	}

	return ev, nil
}

func getFromEnvPtr(envVar string) *string {
	ev := os.Getenv(envVar)

	if ev == "" {
		return nil
	}

	return &ev
}
