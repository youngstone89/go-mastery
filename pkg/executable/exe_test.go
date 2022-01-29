package executable_test

import (
	"go-mastery/pkg/executable"
	"testing"
)

func TestAdminGetConfigNotifyWebhook(t *testing.T) {
	executable.AdminGetConfigNotifyWebhook()
}

func TestSetNotifyWebhook(t *testing.T) {
	executable.SetNotifyWebhook("4", "0", "http://77f5-125-176-145-160.ngrok.io/api/v1/minio/events", "")
}

func TestRestartMinioService(t *testing.T) {
	executable.RestartMinioService()
}
