package times

import (
	"testing"
	"time"
)

func TestConvertTime(t *testing.T) {

	time := time.Now()
	t.Log(time)
	t.Log(time.Date())
	t.Log(time.Format("01-02-2006"))
	t.Log(time.Clock())
	t.Log(time.IsZero())

}
