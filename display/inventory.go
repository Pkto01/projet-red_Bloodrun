package display

import "fmt"

func accessInventory(j *Joueur) {
	fmt.Println("\n--- Inventaire ---")
	if len(j.Inventaire) == 0 {
		fmt.Println("Inventaire vide !")
		return
	}
	for i, item := range j.Inventaire {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	fmt.Printf("\nArgent restant : %d pièces\n", j.Argent)
}

func takePot(c *Character) {
	// Cherche une potion
	found := -1
	for i, item := range c.Inventory {
		if item == "potion" {
			found = i
			break
		}
	}

	if found == -1 {
		fmt.Println("Pas de potion dans l'inventaire !")
		return
	}

	// Supprime UNE potion
	c.Inventory = append(c.Inventory[:found], c.Inventory[found+1:]...)

	// Soigne le perso
	c.Pv += 50
	if c.Pv > c.Pvmax {
		c.Pv = c.Pvmax
	}

	fmt.Printf("%s utilise une potion ! PV : %d/%d\n", c.Name, c.Pv, c.Pvmax)
}
