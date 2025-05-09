package item

import (
	"fmt"
	"study/interface/class"
)

type ManaPotion struct {
	Amount int
}

func (m ManaPotion) RestoreMana(c class.Caster) {
	c.SetMana(c.GetMana() + m.Amount)
	fmt.Printf("Restored %d mana. Current mana: %d\n", m.Amount, c.GetMana())
}
