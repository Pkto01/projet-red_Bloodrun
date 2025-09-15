package display

import (
	"fmt"
	"projet-red_Bloodrun/character"
)

func AccessInventory(j *character.Character) {
	fmt.Println("\n--- Inventaire ---")

	// Compter le nombre de potions
	potionCount := 0
	var otherItems []string

	for _, item := range j.Inventory {
		if item == "Potion de vie" {
			potionCount++
		} else {
			otherItems = append(otherItems, item)
		}
	}

	// Afficher les potions
	if potionCount > 0 {
		fmt.Printf("%d. %d Potions de vie\n", 1, potionCount)
	}

	// Afficher les autres items
	for i, item := range otherItems {
		fmt.Printf("%d. %s\n", potionCount+1+i, item)
	}

	if potionCount == 0 && len(otherItems) == 0 {
		fmt.Println("Inventaire vide !")
		return
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

	AccessInventory(c)
}
