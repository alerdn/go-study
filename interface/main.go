package main

import (
	"study/interface/character"
	"study/interface/class"
	"study/interface/item"
)

// Potions

func main() {
	amon := &character.Character{
		Class: &class.Mage{
			HP:   100,
			Mana: 100,
			Spells: []class.Spell{
				{Name: "Fireball", Power: 30, ManaCost: 20},
				{Name: "Ice Spike", Power: 25, ManaCost: 45},
				{Name: "Lightning Bolt", Power: 35, ManaCost: 25},
			},
		},
	}

	m := amon.Class.(*class.Mage)
	m.Cast(m.Spells[0])
	m.Cast(m.Spells[2])
	m.Cast(m.Spells[2])
	m.Cast(m.Spells[1])

	healthPotion := item.HealthPotion{Amount: 50}
	healthPotion.RestoreHealth(m)

	manaPotion := item.ManaPotion{Amount: 50}
	manaPotion.RestoreMana(m)
}
