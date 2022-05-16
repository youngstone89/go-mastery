package di_test

import (
	"go-mastery/pkg/software-engineering-in-go/di"
	"testing"
)

type Store interface {
	Sell(string) Item
}

type Item = interface {
	Price() uint
}

func TestStore(t *testing.T) {
	var store Store = di.Store{}
	var item Item = store.Sell("1")

	println(item.Price())
}
