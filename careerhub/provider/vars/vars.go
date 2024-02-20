package vars

import (
	"fmt"
	"os"
)

type DBUser struct {
	Username string
	Password string
}

type Vars struct {
	MongoUri     string
	DbName       string
	DBUser       *DBUser
	GrpcEndpoint string
	PostLogUrl   string
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

	mongoUri, err := getFromEnv("MONGO_URI")
	if err != nil {
		return nil, err
	}
	dbUsername := getFromEnvPtr("DB_USERNAME")
	dbPassword := getFromEnvPtr("DB_PASSWORD")

	dbName, err := getFromEnv("DB_NAME")
	if err != nil {
		return nil, err
	}

	var dbUser *DBUser
	if dbUsername != nil && dbPassword != nil {
		dbUser = &DBUser{
			Username: *dbUsername,
			Password: *dbPassword,
		}
	}

	grpcEndpoint, err := getFromEnv("GRPC_ENDPOINT")
	if err != nil {
		return nil, err
	}

	postLogUrl, err := getFromEnv("POST_LOG_URL")
	if err != nil {
		return nil, err
	}

	return &Vars{
		MongoUri:     mongoUri,
		DbName:       dbName,
		DBUser:       dbUser,
		GrpcEndpoint: grpcEndpoint,
		PostLogUrl:   postLogUrl,
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
