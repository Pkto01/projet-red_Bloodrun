package main

import "fmt"

func main() {
	displayInfo()
	arthur := initCharacter("Arthur", "Barbare", 1, 100, 40, []string{"potion", "potion", "potion"})
	displayInfo(arthur)
}

func isDead(j *Joueur) {
	if j.PV <= 0 {
		fmt.Printf("%s est mort...\n", j.Nom)
		// Résurrection avec 50% des PV max
		j.PV = j.PVmax / 2
		fmt.Printf("%s est ressuscité avec %d/%d PV !\n", j.Nom, j.PV, j.PVmax)
	}
}
