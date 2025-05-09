package class

import "fmt"

type Spell struct {
	Name     string
	Power    int
	ManaCost int
}

type Mage struct {
	HP     int
	Mana   int
	Spells []Spell
}

func (m *Mage) Name() string {
	return "Mage"
}

func (m *Mage) GetHP() int {
	return m.HP
}

func (m *Mage) SetHP(hp int) {
	m.HP = hp
}

func (m *Mage) GetMana() int {
	return m.Mana
}

func (m *Mage) SetMana(mana int) {
	m.Mana = mana
}

func (m *Mage) Cast(spell Spell) {
	if m.Mana < spell.ManaCost {
		fmt.Printf("Not enough mana to cast %s - %d/%d\n", spell.Name, spell.ManaCost, m.Mana)
		return
	}

	m.Mana -= spell.ManaCost
	fmt.Printf("Casting %s - dealing %d damage - mana left: %d\n", spell.Name, spell.Power, m.Mana)
}

func (m *Mage) String() string {
	return fmt.Sprintf("Name: %s, HP: %d, Mana: %d, Spells: %v", m.Name(), m.GetHP(), m.Mana, m.Spells)
}
