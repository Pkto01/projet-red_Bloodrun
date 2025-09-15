package display

import (
	"fmt"
	"projet-red_Bloodrun/character"
)

func AccessInventory(j *character.Character) {
	fmt.Println("\n--- Inventaire ---")
	if len(j.Inventory) == 0 {
		fmt.Println("Inventaire vide !")
		return
	}
	for i, item := range j.Inventory {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	fmt.Printf("\nArgent restant : %d pièces\n", j.Money)

	quitter := false
	for !quitter {
		fmt.Println("\n=== Menu Inventaire ===")
		fmt.Println("1. Utiliser une potion de vie")
		fmt.Println("4. Quitter")

		choix := lireEntree("Votre choix : ")

		switch choix {
		case "1":
			takePot(j)
		case "4":
			fmt.Println("Retour au menu principal.")
			quitter = true
		default:
			fmt.Println("Choix invalide !")
		}
	}
}

// --- Utilisation d'une potion ---
func takePot(c *character.Character) {
	// Cherche une potion de vie
	found := -1
	for i, item := range c.Inventory {
		if item == "Potion de vie" {
			found = i
			break
		}
	}

	if found == -1 {
		fmt.Println("Pas de Potion de vie dans l'inventaire !")
		return
	}

	// Supprime UNE potion
	c.Inventory = append(c.Inventory[:found], c.Inventory[found+1:]...)

	// Soigne le perso
	heal := 50
	c.Pv += heal
	if c.Pv > c.Pvmax {
		c.Pv = c.Pvmax
	}

	fmt.Printf("%s utilise une potion de vie (+%d PV) ! PV : %d/%d\n", c.Name, heal, c.Pv, c.Pvmax)
}
