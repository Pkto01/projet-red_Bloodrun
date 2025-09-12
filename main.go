package main

import (
	"fmt"
	"projet-red_Bloodrun/character"
	"projet-red_Bloodrun/display"
)

func main() {
	arthur := character.InitCharacter("Arthur", "Barbare", 1, 40, 100, 40, []string{"potion", "potion", "potion"})
	display.Money(&arthur)
}

func isDead(j *Joueur) {
	if j.PV <= 0 {
		fmt.Printf("%s est mort...\n", j.Nom)
		// Résurrection avec 50% des PV max
		j.PV = j.PVmax / 2
		fmt.Printf("%s est ressuscité avec %d/%d PV !\n", j.Nom, j.PV, j.PVmax)
	}
}
