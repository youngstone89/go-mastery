package rate_test

import (
	"go-mastery/pkg/rate"
	"testing"
)

func TestDoRateLimit(t *testing.T) {

	rate.DoRateLimit()

}

func TestDoBurstyLimit(t *testing.T) {

	rate.DoBurstyLimit()

}

func TestDoRateLimitWithGOPackage(t *testing.T) {
	rate.DoRateLimitWithGOPackage()

}
