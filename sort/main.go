package main

import (
	"fmt"
	"sort"
)

type Spell struct {
	Cost  int
	Power int
}

func (s Spell) String() string {
	return fmt.Sprintf("Spell{Cost: %d, Power: %d}", s.Cost, s.Power)
}

// Para ordernar é preciso implementar a interface sort.Interface
type ByCost []Spell

func (b ByCost) Len() int {
	return len(b)
}

func (b ByCost) Less(i, j int) bool {
	return b[i].Cost < b[j].Cost
}

func (b ByCost) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func main() {
	ii := []int{3, 8, 6, 0}
	sort.Ints(ii)

	fmt.Println("Sorted integers:", ii)

	spells := []Spell{
		{Cost: 3, Power: 10},
		{Cost: 1, Power: 5},
		{Cost: 2, Power: 8},
	}

	sort.Sort(ByCost(spells))
	fmt.Println("Sorted spells by cost:", spells)
}
