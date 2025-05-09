package character

import (
	"fmt"
	"study/interface/class"
)

type Character struct {
	class.Class
}

func (c *Character) String() string {
	return fmt.Sprintf("Class: %v,", c.Class)
}
