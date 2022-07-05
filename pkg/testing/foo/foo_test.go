package foo

import (
	"go-mastery/pkg/testing/mocks"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestFoo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockFoo(ctrl)

	m.EXPECT().Bar(gomock.Eq(99)).Return(101)

	SUT(m)
}
