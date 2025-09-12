package main

import (
	"fmt"
	"projet-red_Bloodrun/character"
	"projet-red_Bloodrun/display"
)

func main() {
	arthur := character.InitCharacter("Arthur", "Barbare", 1, 100, 40, []string{"potion", "potion", "potion"})
	display.DisplayInfo(arthur)
}

func isDead(j character.Character) {
	if j.Pv <= 0 {
		fmt.Printf("%s est mort...\n", j.Name)
		// Résurrection avec 50% des PV max
		j.Pv = j.Pvmax / 2
		fmt.Printf("%s est ressuscité avec %d/%d PV !\n", j.Name, j.Pv, j.Pvmax)
	}
}
