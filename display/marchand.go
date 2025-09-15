package display

import (
	"bufio"
	"fmt"
	"os"
	"projet-red_Bloodrun/character"
	"strconv"
	"strings"
)

type Item struct {
	Nom  string
	Prix int
}

// --- Fonctions Inventaire ---
func addInventory(j character.Character, item Item) {
	j.Inventory = append(j.Inventory, item.Nom)
	fmt.Printf(">> Vous avez acheté : %s\n", item.Nom)
}

// --- Marchand ---
func marchand(j *character.Character, shop []Item) {
	for {
		fmt.Println("\n=== Marchand ===")
		fmt.Printf("Argent disponible : %d pièces\n", j.Money)
		for i, item := range shop {
			fmt.Printf("%d. %s (%d pièces)\n", i+1, item.Nom, item.Prix)
		}
		fmt.Println("0. Retour")

		choix := lireEntree("Votre choix : ")
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
			addInventory(*j, item)
		} else {
			fmt.Println("Pas assez d'argent !")
		}
	}
}

// --- Infos joueur ---
func afficherInfos(j character.Character) {
	fmt.Println("\n--- Infos du personnage ---")
	fmt.Printf("Nom : %s\n", j.Name)
	fmt.Printf("PV : %d / %d\n", j.Pv, j.Pvmax)
	fmt.Printf("Argent : %d pièces\n", j.Money)
	AccessInventory(&j)
}

// --- Entrée utilisateur ---
func lire(texte string) string {
	fmt.Print(texte)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

// --- Programme principal ---
func Money(j *character.Character) {

	// Boutique du marchand
	shop := []Item{
		{"Potion de vie", 20},
		{"Épée en fer", 50},
		{"Bouclier en bois", 30},
		{"Arc basique", 40},
	}

	quitter := false
	for !quitter {
		fmt.Println("\n=== MENU ===")
		fmt.Println("1. Afficher infos personnage")
		fmt.Println("2. Accéder à l'inventaire")
		fmt.Println("3. Marchand")
		fmt.Println("4. Quitter")

		choix := lireEntree("Votre choix : ")

		switch choix {
		case "1":
			DisplayInfo(*j)
		case "2":
			AccessInventory(j)
		case "3":
			marchand(j, shop)
		case "4":
			fmt.Println("Au revoir !")
			quitter = true
		default:
			fmt.Println("Choix invalide !")
		}
	}
}
