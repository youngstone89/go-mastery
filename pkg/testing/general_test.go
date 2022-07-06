package testing

import (
	"errors"
	"fmt"
	"go-mastery/pkg/testing/mocks"
	"go-mastery/pkg/testing/slice"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestAssignableToTypeOf(t *testing.T) {

	fn := func() {}
	fn2 := func() {}
	match := gomock.AssignableToTypeOf(fn).Matches(fn2)
	if match {
		t.Log("match")
	} else {
		t.Log("no match")
	}
}

func TestDoAndReturn(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFoo := mocks.NewMockFoo(ctrl)
	var i int
	mockFoo.EXPECT().Bar(gomock.AssignableToTypeOf(i)).DoAndReturn(
		func(arg int) interface{} {
			return arg
		},
	)
	r := mockFoo.Bar(1)
	fmt.Println(r)
}

func TestDo(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDoer := mocks.NewMockDoer(ctrl)
	mockDoer.EXPECT().DoSomething(gomock.AssignableToTypeOf(1), gomock.AssignableToTypeOf("")).Do(
		func(arg1 int, arg2 string) {
			fmt.Println(arg1, arg2)
		},
	)
	mockDoer.DoSomething(1, "hi")
}

func TestDoTimes(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDoer := mocks.NewMockDoer(ctrl)
	mockDoer.EXPECT().DoSomething(gomock.AssignableToTypeOf(1), gomock.AssignableToTypeOf("")).Do(
		func(arg1 int, arg2 string) {
			fmt.Println(arg1, arg2)
		},
	).Times(1)
	mockDoer.DoSomething(1, "hi")
	mockDoer.EXPECT().DoSomething(gomock.Eq(2), gomock.Eq("hi")).Do(
		func(arg1 int, arg2 string) {
			fmt.Println(arg1, arg2)
		},
	).Times(1)
	mockDoer.DoSomething(2, "hi")
}

func TestDoMinTimes(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDoer := mocks.NewMockDoer(ctrl)
	mockDoer.EXPECT().DoSomething(gomock.AssignableToTypeOf(1), gomock.AssignableToTypeOf("")).Do(
		func(arg1 int, arg2 string) {
			fmt.Println(arg1, arg2)
		},
	).MinTimes(1)

}

func TestInOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDoer := mocks.NewMockDoer(ctrl)
	gomock.InOrder(
		mockDoer.EXPECT().DoSomething(gomock.Eq(1), gomock.Eq("1")).Do(
			func(arg1 int, arg2 string) {
				fmt.Println(arg1, arg2)
			},
		),
		mockDoer.EXPECT().DoSomething(gomock.Eq(2), gomock.Eq("2")).Do(
			func(arg1 int, arg2 string) {
				fmt.Println(arg1, arg2)
			},
		),
		mockDoer.EXPECT().DoSomething(gomock.Eq(3), gomock.Eq("3")).Do(
			func(arg1 int, arg2 string) {
				fmt.Println(arg1, arg2)
			},
		),
	)
	mockDoer.DoSomething(1, "1")
	mockDoer.DoSomething(2, "2")
	mockDoer.DoSomething(3, "3")

}

func TestAfter(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDoer := mocks.NewMockDoer(ctrl)
	first := mockDoer.EXPECT().DoSomething(gomock.Eq(1), gomock.Eq("1")).Do(
		func(arg1 int, arg2 string) {
			fmt.Println(arg1, arg2)
		}).Return(errors.New("error"))
	second := mockDoer.EXPECT().DoSomething(gomock.Eq(2), gomock.Eq("2")).After(first).Do(
		func(arg1 int, arg2 string) {
			fmt.Println(arg1, arg2)
		})

	mockDoer.EXPECT().DoSomething(gomock.Eq(3), gomock.Eq("3")).After(second).Do(
		func(arg1 int, arg2 string) {
			fmt.Println(arg1, arg2)
		})

	r := mockDoer.DoSomething(1, "1")
	fmt.Println(r)
	mockDoer.DoSomething(2, "2")
	mockDoer.DoSomething(3, "3")

}
func TestSetArg(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockSlice := mocks.NewMockSlice(ctrl)
	vals := make([]int, 1)
	var iSlice []int
	vals[0] = 0

	newVals := make([]int, 1)
	newVals[0] = 5
	mockSlice.EXPECT().Do(gomock.AssignableToTypeOf(iSlice)).SetArg(0, newVals).Do(
		func(arg1 []int) {
			for _, v := range arg1 {
				fmt.Println(v)
			}
		},
	)
	Run(mockSlice, vals)
}

func Run(m slice.Slice, vals []int) {
	m.Do(vals)
}
