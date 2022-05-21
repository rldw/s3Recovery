package services

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/sirupsen/logrus"
	"time"
)

type ICloudwatchUtils interface {
	GetBucketMetrics(bucketName string) (map[string]string, error)
}

type CloudwatchUtils struct{}

func (CloudwatchUtils) GetBucketMetrics(bucketName string) (map[string]string, error) {
	session := session.Must(session.NewSession())
	cfg := aws.NewConfig().WithCredentialsChainVerboseErrors(true)
	svc := cloudwatch.New(session, cfg)

	// number of objects in bucket
	numOfObjectsInput := getMetricDataInput("NumberOfObjects", bucketName)
	numObjectsResult, err := svc.GetMetricData(&numOfObjectsInput)
	if err != nil {
		return nil, err
	}

	// TODO bucket size
	//bucketSizeInput := getMetricDataInput("BucketSizeBytes", bucketName)
	//bucketSizeResult, err := svc.GetMetricData(&bucketSizeInput)
	//if err != nil {
	//	return nil, err
	//}

	logrus.Infof("Number of objects in bucket: %v", *numObjectsResult.MetricDataResults[0].Values[0])
	//logrus.Infof("Bucket size in bytes: %v", bucketSizeResult)
	//logrus.Infof("Bucket size in bytes: %v", *bucketSizeResult.MetricDataResults[0].Values[0])
	return nil, nil
}

func getMetricDataInput(metricName string, bucketName string) cloudwatch.GetMetricDataInput {
	return cloudwatch.GetMetricDataInput{
		EndTime:       aws.Time(time.Now()),
		LabelOptions:  nil,
		MaxDatapoints: aws.Int64(1),
		MetricDataQueries: []*cloudwatch.MetricDataQuery{
			{
				Id: aws.String("foo"),
				MetricStat: &cloudwatch.MetricStat{
					Metric: &cloudwatch.Metric{
						Dimensions: []*cloudwatch.Dimension{
							{
								Name:  aws.String("StorageType"),
								Value: aws.String("AllStorageTypes"),
							},
							{
								Name:  aws.String("BucketName"),
								Value: aws.String(bucketName),
							},
						},
						MetricName: aws.String(metricName),
						Namespace:  aws.String("AWS/S3"),
					},
					Period: aws.Int64(60 * 60 * 24 * 2),
					Stat:   aws.String("Maximum"),
					//Unit:   nil,
				},
			},
		},
		ScanBy:    nil,
		StartTime: aws.Time(time.Now().AddDate(0, 0, -3)),
	}
}
