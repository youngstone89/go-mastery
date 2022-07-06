package slice

//go:generate mockgen -destination ../mocks/mock_slice.go -package=mocks -source slice.go
type Slice interface {
	Do([]int)
}
