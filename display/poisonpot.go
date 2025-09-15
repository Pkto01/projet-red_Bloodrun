package main

import (
	"fmt"
	"time"
)

func (c *Character) poisonPot() {

	fmt.Println("La potion de poison est activée!")
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		c.PointsDeVieActuels -= 10
		if c.PointsDeVieActuels < 0 {
			c.PointsDeVieActuels = 0
		}
		fmt.Printf("Dégâts infligés! Points de vie actuels: %d/%d\n", c.PointsDeVieActuels, c.PointsDeVieMax)
	}
}
