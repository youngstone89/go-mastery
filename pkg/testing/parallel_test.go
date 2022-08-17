package testing

import (
	"strings"
	"testing"
)

func TestIndex(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		sentence  string
		substring string
		exp       int
	}{
		{"Gophers are amazing", "are", 8},
		{"Gophers are amazing", "Gophers", 0},
		{"Gophers are amazing", "rust", -1},
	}
	for _, tc := range tcs {
		tc := tc // rebind tc into this lexical scope
		t.Run(tc.substring, func(t *testing.T) {
			t.Parallel()
			t.Logf("testing indexing %q for %q", tc.sentence, tc.substring)
			got := strings.Index(tc.sentence, tc.substring)
			if got != tc.exp {
				t.Errorf("unexpected value for indexing %q for %q.  got %d, exp %d", tc.sentence, tc.substring, got, tc.exp)
			}
		})
	}
}
