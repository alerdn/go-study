package class

type Caster interface {
	GetMana() int
	SetMana(int)
	Cast(Spell)
}
