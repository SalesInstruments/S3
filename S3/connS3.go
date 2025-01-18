package S3

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type ConfigS3 struct {
	URL        string
	Key        string
	Secret_key string
	Bucket     string
	Region     string
}

type S3ClientFactory struct{}

func (f *S3ClientFactory) CreateS3Client(c ConfigS3) (*s3.Client, error) {
	customResolver := aws.EndpointResolverWithOptionsFunc(
		func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			if service == s3.ServiceID {
				return aws.Endpoint{
					URL: c.URL,
				}, nil
			}

			return aws.Endpoint{}, &aws.EndpointNotFoundError{}
		})

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(customResolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(c.Key, c.Secret_key, "")),
		config.WithRegion(c.Region), // Укажите регион, если нужно
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	return s3.NewFromConfig(cfg), nil
}
