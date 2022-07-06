package foo

//go:generate mockgen -destination mock_foo.go -package=mocks -source foo.go
type Foo interface {
	Bar(x int) int
}

func SUT(f Foo) {
	f.Bar(99)
}
