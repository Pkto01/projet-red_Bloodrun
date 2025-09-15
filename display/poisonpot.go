package display

import (
	"fmt"
	"projet-red_Bloodrun/character"
	"time"
)

func poisonPot(c *character.Character) {

	fmt.Println("La potion de poison est activée!")
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		c.Pv -= 10
		if c.Pv < 0 {
			c.Pv = 0
		}
		fmt.Printf("Dégâts infligés! Points de vie actuels: %d/%d\n", c.Pv, c.Pvmax)
	}
}
