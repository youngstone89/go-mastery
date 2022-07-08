package backoff

import (
	"errors"
	"fmt"
	"testing"
	"time"

	bff "github.com/cenkalti/backoff/v4"
)

func TestExponentialBackoffReturnNilError(t *testing.T) {
	exponentialBackOff := bff.NewExponentialBackOff()

	operation := func() error {
		fmt.Println("operating...")
		return nil
	}

	err := bff.Retry(operation, exponentialBackOff)
	if err != nil {
		t.Errorf(err.Error())
	}
}
func TestExponentialBackoffReturnPermanentError(t *testing.T) {
	exponentialBackOff := bff.NewExponentialBackOff()

	operation := func() error {
		fmt.Println("operating...")
		return bff.Permanent(errors.New("dummy error"))
	}

	err := bff.Retry(operation, exponentialBackOff)
	fmt.Println(err)
}

func TestExponentialBackoffReturnTransientErrorWithMaxElaspedTime5Seconds(t *testing.T) {
	exponentialBackOff := bff.NewExponentialBackOff()
	exponentialBackOff.MaxElapsedTime = 5 * time.Second

	operation := func() error {
		fmt.Println("operating...")
		return errors.New("transient error")
	}

	err := bff.Retry(operation, exponentialBackOff)
	fmt.Println(err)
}
func TestExponentialBackoffWithMaxRetriesReturnTransientErrorWithMaxElaspedTime2Seconds(t *testing.T) {
	exponentialBackOff := bff.NewExponentialBackOff()
	exponentialBackOff.MaxElapsedTime = 2 * time.Second
	backoffWithRetries := bff.WithMaxRetries(exponentialBackOff, 5)
	operation := func() error {
		fmt.Println("operating...")
		return errors.New("transient error")
	}

	err := bff.Retry(operation, backoffWithRetries)
	fmt.Println(err)
}

func TestExponentialBackoffWithMaxRetriesReturnTransientErrorWithMaxElaspedTime10Seconds(t *testing.T) {
	exponentialBackOff := bff.NewExponentialBackOff()
	exponentialBackOff.MaxElapsedTime = 10 * time.Second
	backoffWithRetries := bff.WithMaxRetries(exponentialBackOff, 5)
	operation := func() error {
		fmt.Println("operating...")
		return errors.New("transient error")
	}

	err := bff.Retry(operation, backoffWithRetries)
	fmt.Println(err)
}
func TestBackoffWithMaxRetries(t *testing.T) {
	backoff := bff.NewConstantBackOff(1 * time.Second)
	backoffWithRetries := bff.WithMaxRetries(backoff, 5)

	operation := func() error {
		fmt.Println("operating...")
		return errors.New("transient error")
	}

	err := bff.Retry(operation, backoffWithRetries)
	fmt.Println(err)
}
