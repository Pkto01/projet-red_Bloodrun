package display

import (
	"fmt"
	"projet-red_Bloodrun/character"
	"strconv"
)

// itemTypes définit le slot pour chaque objet équipable.
var itemTypes = map[string]string{
	// Armes
	"Hache de Berserker": "Weapon",
	"Bâton d'Apprenti":   "Weapon",
	"Marteau Lourd":      "Weapon",

	// Armures
	"Robe de Mage":      "Armor",
	"Armure de Plaques": "Armor",

	// Accessoires
	"Bouclier en Acier":   "Accessory",
	"Grimoire des Ombres": "Accessory",
	"Bottes de Célérité":  "Accessory",
	"Gantelets de Force":  "Accessory",
}

// equipItem gère la logique pour équiper un objet de l'inventaire.
func equipItem(j *character.Character) {
	var equippableItems []string
	var originalIndices []int

	// On ne liste que les objets qui peuvent être équipés
	for i, item := range j.Inventory {
		if _, isEquippable := itemTypes[item]; isEquippable {
			equippableItems = append(equippableItems, item)
			originalIndices = append(originalIndices, i)
		}
	}

	if len(equippableItems) == 0 {
		fmt.Println("\033[90mVous n'avez aucun objet équipable dans votre inventaire.\033[0m")
		return
	}

	fmt.Println("\nQuel objet souhaitez-vous équiper ?")
	for i, item := range equippableItems {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	fmt.Println("0. Annuler")

	choixStr := LireEntree("Votre choix : ")
	choix, err := strconv.Atoi(choixStr)
	if err != nil || choix < 0 || choix > len(equippableItems) {
		fmt.Println("Choix invalide.")
		return
	}

	if choix == 0 {
		return
	}

	selectedItemIndex := choix - 1
	itemName := equippableItems[selectedItemIndex]
	inventoryIndex := originalIndices[selectedItemIndex]
	itemSlot := itemTypes[itemName]

	// Logique de remplacement
	var oldItem string
	switch itemSlot {
	case "Weapon":
		oldItem = j.Equipped.Weapon
		j.Equipped.Weapon = itemName
	case "Armor":
		oldItem = j.Equipped.Armor
		j.Equipped.Armor = itemName
	case "Accessory":
		oldItem = j.Equipped.Accessory
		j.Equipped.Accessory = itemName
	}

	j.Inventory = RemoveFromSliceByIndex(j.Inventory, inventoryIndex)

	if oldItem != "Aucune" {
		addInventory(j, oldItem)
		fmt.Printf("\033[90m%s a été retourné à votre inventaire.\033[0m\n", oldItem)
	}

	fmt.Printf("\033[32mVous avez équipé : %s.\033[0m\n", itemName)
}

// AccessInventory mis à jour avec une pause pour une meilleure expérience utilisateur.
func AccessInventory(j *character.Character) {
	quitter := false
	for !quitter {
		fmt.Println("\n--- Inventaire & Équipement ---")
		fmt.Printf("%s\n", ShowInventory(j))
		// Affiche l'équipement actuel
		fmt.Println("\033[1mÉquipé :\033[0m")
		fmt.Printf("   Arme      : %s\n", j.Equipped.Weapon) // Assurez-vous d'utiliser "Equipped" si vous avez suivi ma correction précédente
		fmt.Printf("   Armure    : %s\n", j.Equipped.Armor)
		fmt.Printf("   Accessoire: %s\n", j.Equipped.Accessory)

		// Affichage de l'inventaire
		fmt.Println("\n\033[1mDans le sac :\033[0m")
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
			fmt.Println("\033[90mInventaire vide.\033[0m")
		} else {
			if potionCount > 0 {
				fmt.Printf("- %d x Potion(s) de vie\n", potionCount)
			}
			for _, item := range otherItems {
				fmt.Printf("- %s\n", item)
			}
		}

		// Menu d'actions
		fmt.Println("\n=== Menu Inventaire ===")
		fmt.Println("1. Utiliser une potion de vie")
		fmt.Println("2. Équiper un objet")
		fmt.Println("3. Quitter")

		choix := LireEntree("Votre choix : ")

		actionTaken := false // Notre nouvelle variable pour contrôler la pause
		switch choix {
		case "1":
			if potionCount > 0 {
				takePot(j)
				actionTaken = true // Une action a été effectuée
			} else {
				fmt.Println("Vous n'avez pas de potions.")
			}
		case "2":
			equipItem(j)
			actionTaken = true // Une action a été effectuée
		case "3":
			fmt.Println("Retour au menu principal.")
			quitter = true
		default:
			fmt.Println("Choix invalide !")
		}

		// Si une action a été effectuée (et qu'on ne quitte pas), on fait une pause.
		if actionTaken {
			LireEntree("\nAppuyez sur Entrée pour continuer...")
		}
	}
}

func ShowInventory(j *character.Character) string {
	limit := getInventoryLimit(j.Class, j.InventoryUpgrades)
	result := fmt.Sprintf("\033[34mInventaire (%d/%d)\033[0m\n", len(j.Inventory), limit)

	if len(j.Inventory) == 0 {
		result += " - (vide)\n"
		return result
	}

	for i, item := range j.Inventory {
		result += fmt.Sprintf(" %d. %s\n", i+1, item)
	}
	return result
}

// takePot - Utilisation d'une potion (INCHANGÉ)
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

	c.Inventory = RemoveFromSliceByIndex(c.Inventory, found) // Utilise notre nouvelle fonction utilitaire

	heal := 50
	c.Pv += heal
	if c.Pv > c.Pvmax {
		c.Pv = c.Pvmax
	}

	fmt.Printf("%s utilise une potion de vie (+%d PV) ! PV : %d/%d\n", c.Name, heal, c.Pv, c.Pvmax)
}
