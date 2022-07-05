package foo

//go:generate mockgen -destination mock_foo.go -package=foo -source foo.go
type Foo interface {
	Bar(x int) int
}

func SUT(f Foo) {
	f.Bar(99)
}
