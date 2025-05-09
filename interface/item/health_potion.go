package item

import (
	"fmt"
	"study/interface/class"
)

type HealthPotion struct {
	Amount int
}

func (h HealthPotion) RestoreHealth(c class.Class) {
	c.SetHP(c.GetHP() + h.Amount)
	fmt.Printf("Restored %d health. Current HP: %d\n", h.Amount, c.GetHP())
}
