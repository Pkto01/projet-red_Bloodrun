package display

import (
	"fmt"
	"projet-red_Bloodrun/character"
	"strconv"
)

type Item struct {
	Nom  string
	Prix int
}

// --- Fonctions Inventaire ---
func getInventoryLimit(class string) int {
	switch class {
	case "Doom Caster":
		return 7
	case "Doom Slayer":
		return 11
	case "Doom Bastion":
		return 15
	default:
		return 0
	}
}

func addInventory(j *character.Character, item Item) {
	limit := getInventoryLimit(j.Class)
	if len(j.Inventory) >= limit {
		fmt.Println("Inventaire plein ! Impossible d'ajouter :", item.Nom)
		return
	}
	j.Inventory = append(j.Inventory, item.Nom)
	fmt.Printf(">> Vous avez acheté : %s\n", item.Nom)
}

// --- Marchand ---
func Marchand(j *character.Character, shop []Item) {
	for {
		fmt.Println("\n=== Marchand ===")
		fmt.Printf("Argent disponible : %d pièces\n", j.Money)
		for i, item := range shop {
			fmt.Printf("%d. %s (%d pièces)\n", i+1, item.Nom, item.Prix)
		}
		fmt.Println("0. Retour")

		choix := LireEntree("Votre choix : ")
		num, err := strconv.Atoi(choix)
		if err != nil || num < 0 || num > len(shop) {
			fmt.Println("Choix invalide.")
			continue
		}

		if num == 0 {
			return
		}

		item := shop[num-1]
		if j.Money >= item.Prix {
			j.Money -= item.Prix
			addInventory(j, item)
		} else {
			fmt.Println("Pas assez d'argent !")
		}
	}
}

func Forgeron(j *character.Character) {
	// Liste fixe des équipements
	equipements := []string{
		"Chapeau de l’aventurier",
		"Tunique de l’aventurier",
		"Bottes de l’aventurier",
	}

	for {
		fmt.Println("\n=== Forgeron ===")
		fmt.Printf("Argent disponible : %d pièces\n", j.Money)
		for i, eq := range equipements {
			fmt.Printf("%d. %s (5 pièces)\n", i+1, eq)
		}
		fmt.Println("0. Retour")

		choix := LireEntree("Votre choix : ")
		num, err := strconv.Atoi(choix)
		if err != nil || num < 0 || num > len(equipements) {
			fmt.Println("Choix invalide.")
			continue
		}

		if num == 0 {
			return
		}

		if j.Money >= 5 {
			j.Money -= 5
			// Vérifie limite d’inventaire avant ajout
			addInventory(j, Item{Nom: equipements[num-1], Prix: 5})
		} else {
			fmt.Println("Pas assez d’argent pour fabriquer cet équipement !")
		}
	}
}
