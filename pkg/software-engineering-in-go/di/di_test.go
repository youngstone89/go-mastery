package di

import (
	"fmt"
	"testing"
)

var _ HighLevelModule = new(LowLevelModule)

type HighLevelModule interface {
	Do()
}

type LowLevelModule struct {
	HighLevelModule
}

func NewLowLevelModule() HighLevelModule {
	return &LowLevelModule{}
}

func (LowLevelModule) Do() {
	fmt.Println("I am a low level module implementing high level module")
}

func Test(t *testing.T) {
	highLevelModule := NewLowLevelModule()
	lowLevelModule := highLevelModule.(*LowLevelModule)
	lowLevelModule.HighLevelModule = NewLowLevelModule()
	lowLevelModule.HighLevelModule.Do()
	lowLevelModule.Do()

}
