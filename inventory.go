package main

import "fmt"

func accessInventory(inventory []string) {
	if len(inventory) == 0 {
		fmt.Println("L'inventaire est vide !")
		return
	}

	for i, item := range inventory {
		fmt.Printf("%d: %s\n", i+1, item)
	}
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
