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

// String implements fmt.Stringer for the Sword type.
func (s EnchantedSword) String() string {
	return fmt.Sprintf(
		"%s is a sword that can deal %d points of damage to opponents",
		s.name, s.Damage(),
	)
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

func TestSwordToString(t *testing.T) {
	specs := []struct {
		sword fmt.Stringer
		exp   string
	}{
		{
			sword: Sword{name: "Silver Saber"},
			exp:   "Silver Saber is a sword that can deal 2 points of damage to opponents",
		},
		{
			sword: EnchantedSword{Sword{name: "Dragon's Greatsword"}},
			exp:   "Dragon's Greatsword is a sword that can deal 42 points of damage to opponents",
		},
	}
	for specIndex, spec := range specs {
		if got := spec.sword.String(); got != spec.exp {
			t.Errorf("[spec %d] expected to get\n%q\ngot:\n%q",
				specIndex, spec.exp, got)
		}
	}
}
