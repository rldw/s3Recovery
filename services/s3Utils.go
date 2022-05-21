package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/sirupsen/logrus"
)

type IS3Utils interface {
	ListBuckets() ([]types.Bucket, error)
	GetBucketMetadata(bucketName string) map[string]string
}

type S3Utils struct{}

func (S3Utils) ListBuckets() ([]types.Bucket, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-central-1"))
	if err != nil {
		logrus.Error("unable to load SDK config")
		return nil, err
	}

	svc := s3.NewFromConfig(cfg)
	result, err := svc.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		return nil, err
	}

	return result.Buckets, nil
}

func (S3Utils) GetBucketMetadata(bucketName string) map[string]string {
	//TODO implement me
	panic("implement me")
}
