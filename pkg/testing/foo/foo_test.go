package foo

import (
	"go-mastery/pkg/testing/mocks"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

func TestFoo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockFoo(ctrl)

	m.EXPECT().Bar(gomock.Eq(99)).Return(101)

	SUT(m)
}

func TestFooStub(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockFoo(ctrl)

	// Does not make any assertions. Executes the anonymous functions and returns
	// its result when Bar is invoked with 99.
	m.
		EXPECT().
		Bar(gomock.Eq(99)).
		DoAndReturn(func(_ int) int {
			time.Sleep(1 * time.Second)
			return 101
		}).
		AnyTimes()

	// Does not make any assertions. Returns 103 when Bar is invoked with 101.
	m.
		EXPECT().
		Bar(gomock.Eq(101)).
		Return(103).
		AnyTimes()

	SUT(m)

}
