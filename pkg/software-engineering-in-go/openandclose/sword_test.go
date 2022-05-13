package openandclose

import (
	"fmt"
	"testing"
)

type Sword struct {
	name string // Important tip for RPG players: always name your swords!
}

// Damage returns the damage dealt by this sword.
func (Sword) Damage() int {
	return 2
}

// String implements fmt.Stringer for the Sword type.
func (s Sword) String() string {
	return fmt.Sprintf(
		"%s is a sword that can deal %d points of damage to opponents",
		s.name, s.Damage(),
	)
}

type EnchantedSword struct {
	// Embed the Sword type
	Sword
}

// Damage returns the damage dealt by the enchanted sword.
func (EnchantedSword) Damage() int {
	return 42
}

func TestEnchanctedSword(t *testing.T) {
	specs := []struct {
		sword interface {
			Damage() int
		}
		exp int
	}{
		{
			sword: Sword{name: "Silver Saber"},
			exp:   2,
		},
		{
			sword: EnchantedSword{Sword{name: "Dragon's Greatsword"}},
			exp:   42,
		},
	}
	for specIndex, spec := range specs {
		if got := spec.sword.Damage(); got != spec.exp {
			t.Errorf("[spec %d] expected to get damage %d; got %d", specIndex, spec.exp, got)
		}
	}
}
