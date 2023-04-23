package elc_test

import (
	"math/rand"
	"testing"
	"time"
)

func Test(t *testing.T) {
	randN := rand.Intn(100)
	t.Log(randN)

	<-time.After(5 * time.Second * time.Duration(randN))

}
