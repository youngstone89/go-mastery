package user_test

import (
	"errors"
	"go-mastery/pkg/testing/mocks"
	"go-mastery/pkg/testing/user"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestUse(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDoer := mocks.NewMockDoer(mockCtrl)
	testUser := &user.User{Doer: mockDoer}

	// Expect Do to be called once with 123 and "Hello GoMock" as parameters, and return nil from the mocked call.
	mockDoer.EXPECT().DoSomething(123, "Hello GoMock").Return(nil).Times(1)

	testUser.Use()

}

func TestUseReturnsErrorFromDo(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	dummyError := errors.New("dummy error")
	mockDoer := mocks.NewMockDoer(mockCtrl)
	testUser := &user.User{Doer: mockDoer}

	// Expect Do to be called once with 123 and "Hello GoMock" as parameters, and return dummyError from the mocked call.
	mockDoer.EXPECT().DoSomething(123, "Hello GoMock").Return(dummyError).Times(1)

	err := testUser.Use()

	if err != dummyError {
		t.Fail()
	}
}
