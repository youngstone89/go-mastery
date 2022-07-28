package embedding_test

import (
	"fmt"
	"testing"
)

type Embedded struct {
}

func (e *Embedded) print() {
	fmt.Printf("%p\n", e)
}

type Embedding struct {
	Embedded
}

type ExplicitEmbedding struct {
	embedded Embedded
}

func (e *ExplicitEmbedding) print() {
	e.embedded.print()
}

func TestEmbedding(t *testing.T) {
	embedding := new(Embedding)
	embedding.print()
}

func TestExplicitEmbedding(t *testing.T) {
	embedding := new(ExplicitEmbedding)
	embedding.embedded.print()

}
