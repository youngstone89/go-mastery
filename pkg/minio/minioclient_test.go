package minio_test

import (
	"fmt"
	"go-mastery/pkg/minio"
	"testing"

	"github.com/minio/minio-go/v7/pkg/notification"
	"golang.org/x/net/context"
)

func TestGetBucketNotification(t *testing.T) {

	minioClient := minio.NewMinioClient("localhost:9000", "20ee7722-444a-4ff4-afc2-a242bcad1a2b", "d10ddd48-3136-4045-8292-c8ee3543957e", false)
	if minioClient == nil {
		return
	}
	bucket := "audio-logs"
	config, err := minioClient.GetBucketNotification(context.Background(), bucket)
	if err != nil {
		fmt.Println("Failed to get bucket notification configurations for", bucket, err)
		return

	}
	fmt.Printf("bucket notification config: %v \n", config)

}

func TestSetBucketNotification(t *testing.T) {
	minioClient := minio.NewMinioClient("localhost:9000", "20ee7722-444a-4ff4-afc2-a242bcad1a2b", "d10ddd48-3136-4045-8292-c8ee3543957e", false)
	if minioClient == nil {
		return
	}
	bucket := "audio-logs"
	err := minioClient.SetBucketNotification(context.Background(), bucket, "minio", "sqs", "us-east-1", "804605494417", "webhook", []notification.EventType{notification.ObjectCreatedPut})
	if err != nil {
		fmt.Println("Failed to set bucket notification configurations for", bucket, err)
		return

	}
	config, err := minioClient.GetBucketNotification(context.Background(), bucket)
	if err != nil {
		fmt.Println("Failed to get bucket notification configurations for", bucket, err)
		return
	}
	fmt.Printf("bucket notification config: %v", config)
}

func TestSetBucketNotificationNone(t *testing.T) {
	minioClient := minio.NewMinioClient("localhost:9000", "20ee7722-444a-4ff4-afc2-a242bcad1a2b", "d10ddd48-3136-4045-8292-c8ee3543957e", false)
	if minioClient == nil {
		return
	}
	bucket := "audio-logs"
	err := minioClient.SetBucketNotification(context.Background(), bucket, "minio", "sqs", "", "4", "webhook", []notification.EventType{notification.ObjectCreatedPut})
	if err != nil {
		fmt.Println("Failed to set bucket notification configurations for", bucket, err)
		return

	}
	config, err := minioClient.GetBucketNotification(context.Background(), bucket)
	if err != nil {
		fmt.Println("Failed to get bucket notification configurations for", bucket, err)
		return
	}
	fmt.Printf("bucket notification config: %v", config)
}

func Test(t *testing.T) {

	minioClient := minio.NewMinioClient("localhost:9000", "20ee7722-444a-4ff4-afc2-a242bcad1a2b", "d10ddd48-3136-4045-8292-c8ee3543957e", false)
	if minioClient == nil {
		return
	}

}
