package config

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	c "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func LocalDyno() (*dynamodb.Client, error) {
	cfg, err := c.LoadDefaultConfig(context.Background(),
		c.WithRegion("us-east-1"),
		c.WithEndpointResolver(aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
			return aws.Endpoint{URL: "http://localhost:8000"}, nil
		})),
		c.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     "",
				SessionToken:    "",
				SecretAccessKey: "",
				Source:          "",
			},
		}),
	)

	if err != nil {
		return nil, err
	}

	return dynamodb.NewFromConfig(cfg), nil
}
