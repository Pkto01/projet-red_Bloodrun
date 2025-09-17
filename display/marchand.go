package display

import (
	"fmt"
	"projet-red_Bloodrun/character"
	"strconv"
	"strings"
)

// Item est utilisé pour les objets simples (achat/vente)
type Item struct {
	Nom  string
	Prix int
}

// CraftableItem est utilisé pour les recettes du forgeron
type CraftableItem struct {
	Nom    string
	Prix   int
	Requis map[string]int // Map[nom_materiau]quantité
}

// --- Fonctions d'aide à l'inventaire ---

// countItem compte combien de fois un objet spécifique apparaît dans l'inventaire.
func countItem(inventory []string, itemName string) int {
	count := 0
	for _, item := range inventory {
		if item == itemName {
			count++
		}
	}
	return count
}

// removeItems retire un certain nombre d'un objet spécifique de l'inventaire.
func removeItems(inventory []string, itemName string, quantity int) []string {
	var newInventory []string
	removedCount := 0
	for _, item := range inventory {
		if item == itemName && removedCount < quantity {
			removedCount++
		} else {
			newInventory = append(newInventory, item)
		}
	}
	return newInventory
}

func getInventoryLimit(class string) int {
	switch class {
	case "Doom Caster":
		return 7
	case "Doom Slayer":
		return 11
	case "Doom Bastion":
		return 15
	default:
		return 10
	}
}

func addInventory(j *character.Character, itemName string) {
	limit := getInventoryLimit(j.Class)
	if len(j.Inventory) >= limit {
		fmt.Println("\033[31mInventaire plein ! Impossible d'ajouter : " + itemName + "\033[0m")
		return
	}
	j.Inventory = append(j.Inventory, itemName)
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
			addInventory(j, item.Nom)
			fmt.Printf("\033[32m>> Vous avez acheté : %s\033[0m\n", item.Nom)
		} else {
			fmt.Println("\033[31mPas assez d'argent !\033[0m")
		}
	}
}

// --- Forgeron ---
func Forgeron(j *character.Character, recipes []CraftableItem) {
	for {
		fmt.Println("\n=== Forgeron ===")
		fmt.Printf("Argent disponible : %d pièces\n", j.Money)
		fmt.Println("Objets fabricables :")

		// --- BOUCLE D'AFFICHAGE ---
		// Affiche toutes les recettes et colore les matériaux requis.
		for i, recipe := range recipes {
			var reqs []string
			// La variable canCraft a été supprimée ici car elle était inutile.
			// La logique de couleur suffit pour l'affichage.
			for material, quantity := range recipe.Requis {
				playerHas := countItem(j.Inventory, material)
				color := "\033[32m" // Vert par défaut
				if playerHas < quantity {
					color = "\033[31m" // Rouge si matériaux manquants
				}
				reqs = append(reqs, fmt.Sprintf("%s%d %s (%d/%d)\033[0m", color, quantity, material, playerHas, quantity))
			}
			fmt.Printf("%d. %s (%d pièces) - Requis: %s\n", i+1, recipe.Nom, recipe.Prix, strings.Join(reqs, ", "))
		}
		fmt.Println("0. Retour")

		choix := LireEntree("Votre choix : ")
		num, err := strconv.Atoi(choix)
		if err != nil || num < 0 || num > len(recipes) {
			fmt.Println("Choix invalide.")
			continue
		}

		if num == 0 {
			return
		}

		recipe := recipes[num-1]

		// --- LOGIQUE DE FABRICATION ---
		// Ici, la variable canCraft est essentielle pour valider toutes les conditions.
		canCraft := true
		if j.Money < recipe.Prix {
			fmt.Println("\033[31mPas assez d'argent !\033[0m")
			canCraft = false
		}
		for material, quantity := range recipe.Requis {
			if countItem(j.Inventory, material) < quantity {
				fmt.Printf("\033[31mMatériaux manquants : %s\033[0m\n", material)
				canCraft = false
			}
		}

		// Si toutes les conditions sont remplies (canCraft est resté true)
		if canCraft {
			j.Money -= recipe.Prix
			for material, quantity := range recipe.Requis {
				j.Inventory = removeItems(j.Inventory, material, quantity)
			}
			addInventory(j, recipe.Nom)
			fmt.Printf("\033[32mVous avez fabriqué : %s !\033[0m\n", recipe.Nom)
		}
	}
}
