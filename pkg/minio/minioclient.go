package minio

import (
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/minio/minio-go/v7/pkg/notification"
	"golang.org/x/net/context"
)

type MinioClient struct {
	MinioClient *minio.Client
}

func (m *MinioClient) GetBucketNotification(ctx context.Context, bucketName string) (notification.Configuration, error) {
	bucketNotification, err := m.MinioClient.GetBucketNotification(ctx, bucketName)
	if err != nil {
		fmt.Println("Unable to get the bucket notification: ", err)
		return bucketNotification, err
	}
	for _, queueConfig := range bucketNotification.LambdaConfigs {
		for _, e := range queueConfig.Events {
			fmt.Println(e + " event is enabled")
		}
	}
	for _, queueConfig := range bucketNotification.TopicConfigs {
		for _, e := range queueConfig.Events {
			fmt.Println(e + " event is enabled")
		}
	}

	for _, queueConfig := range bucketNotification.QueueConfigs {
		for _, e := range queueConfig.Events {
			fmt.Println(e + " event is enabled")
		}
	}
	return bucketNotification, err
}

func (m *MinioClient) SetBucketNotification(ctx context.Context, bucketName string, partition string, service string, region string, accountId string, resource string, events []notification.EventType) error {
	//arn:minio:sqs::1:webhook
	// partitio = minio
	// service = sqs
	// region = ?
	// accountid = 1
	// resource = webhook
	//"arn:" + arn.Partition + ":" + arn.Service + ":" + arn.Region + ":" + arn.AccountID + ":" + arn.Resource
	queueArn := notification.NewArn(partition, service, region, accountId, resource)
	fmt.Printf("%s \n", queueArn)
	//
	queueConfig := notification.NewConfig(queueArn)
	queueConfig.AddEvents(events...)
	fmt.Printf("%s \n", events)
	//
	config := notification.Configuration{}
	config.AddQueue(queueConfig)
	fmt.Printf("%v \n", queueConfig)
	return m.MinioClient.SetBucketNotification(ctx, bucketName, config)
}
func NewMinioClient(endpoint string, accessKeyId string, secretAccessKey string, useSSL bool) *MinioClient {
	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyId, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil
	}
	return &MinioClient{MinioClient: minioClient}
}

func DoMinioTest() {
	minioClient := NewMinioClient("localhost:9000", "minio", "minio123", false)
	if minioClient == nil {
		return
	}
	bucket := "audio-logs"
	_, err := minioClient.GetBucketNotification(context.Background(), bucket)
	if err != nil {
		fmt.Println("Failed to get bucket notification configurations for", bucket, err)
		return
	}
	err = minioClient.SetBucketNotification(context.Background(), bucket, "minio", "sqs", "us-east-1", "1004", "webhook", []notification.EventType{notification.ObjectCreatedPut})
	//err = minioClient.SetBucketNotification(context.Background(),bucket,"minio","sqs","us-east-1","2","kafka",[]notification.EventType{notification.ObjectCreatedPut})
	if err != nil {
		fmt.Println("Failed to get bucket notification configurations for", bucket, err)
		return
	}
	_, err = minioClient.GetBucketNotification(context.Background(), bucket)
	if err != nil {
		fmt.Println("Failed to get bucket notification configurations for", bucket, err)
		return
	}

}
