package di

// implements store interface in di_test package
type Store struct{}

//implements Store interfaces
func (Store) Sell(id string) item {
	return bookStruct{}
}

type item = interface { // <-- alias
	Price() uint
}

// implements itemInterface
type bookStruct struct{}

func (bookStruct) Price() uint { return 999 }
