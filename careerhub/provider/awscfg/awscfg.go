package awscfg

import (
	"careerhub-dataprovider/careerhub/provider/utils/terr"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func Config() (*aws.Config, error) {
	awsConfig, err := config.LoadDefaultConfig(context.Background())
	return &awsConfig, terr.Wrap(err)
}
