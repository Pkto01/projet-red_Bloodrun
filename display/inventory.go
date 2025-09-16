package display

import (
	"fmt"
	"projet-red_Bloodrun/character"
)

func AccessInventory(j *character.Character) {
	quitter := false
	for !quitter { // La boucle principale qui gère le menu de l'inventaire

		// --- Affichage de l'inventaire (maintenant DANS la boucle) ---
		fmt.Println("\n--- Inventaire ---")

		potionCount := 0
		var otherItems []string
		for _, item := range j.Inventory {
			if item == "Potion de vie" {
				potionCount++
			} else {
				otherItems = append(otherItems, item)
			}
		}

		if potionCount == 0 && len(otherItems) == 0 {
			fmt.Println("Inventaire vide !")
			// Pas besoin de continuer, on peut quitter directement.
			quitter = true
			continue // Revient au début de la boucle, qui va se terminer.
		}

		// Afficher les potions et autres items
		if potionCount > 0 {
			fmt.Printf("Vous avez %d Potion(s) de vie.\n", potionCount)
		}
		for _, item := range otherItems {
			fmt.Printf("- %s\n", item)
		}

		fmt.Printf("\nArgent restant : %d pièces\n", j.Money)

		// --- Menu d'actions ---
		fmt.Println("\n=== Menu Inventaire ===")
		if potionCount > 0 {
			fmt.Println("1. Utiliser une potion de vie")
		}
		fmt.Println("4. Quitter")

		choix := LireEntree("Votre choix : ")
		switch choix {
		case "1":
			if potionCount > 0 {
				takePot(j) // Appelle la fonction qui ne fait QUE l'action
			} else {
				fmt.Println("Choix invalide !")
			}
		case "4":
			fmt.Println("Retour au menu principal.")
			quitter = true // Met fin à la boucle proprement
		default:
			fmt.Println("Choix invalide !")
		}
	}
}

// --- Utilisation d'une potion ---
// Cette fonction ne fait plus que son travail : utiliser la potion.
// Elle ne relance PAS le menu.
func takePot(c *character.Character) {
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

	// Soigne le personnage
	heal := 50
	c.Pv += heal
	if c.Pv > c.Pvmax {
		c.Pv = c.Pvmax
	}

	fmt.Printf("%s utilise une potion de vie (+%d PV) ! PV : %d/%d\n", c.Name, heal, c.Pv, c.Pvmax)
	// PAS d'appel à AccessInventory(c) ici !
}
