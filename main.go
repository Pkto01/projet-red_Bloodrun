package main

import (
	"fmt"
	"projet-red_Bloodrun/character"
	"projet-red_Bloodrun/display"
)

func main() {
	// Initialisation du personnage
	arthur := character.InitCharacter("Arthur", "Barbare", 1, 40, 100, 40, []string{"Coup de poing"}, []string{"Potion de vie", "Potion de vie", "Potion de vie"})

	// Boutique du marchand
	shop := []display.Item{
		{"Potion de vie", 20},
		{"Épée en fer", 50},
		{"Bouclier en bois", 30},
		{"Arc basique", 40},
	}

	// Boucle de jeu principale
	quitter := false
	for !quitter {
		fmt.Println("\n=== MENU ===")
		fmt.Println("1. Afficher infos personnage")
		fmt.Println("2. Accéder à l'inventaire")
		fmt.Println("3. Marchand")
		fmt.Println("4. Quitter")

		choix := display.LireEntree("Votre choix : ")

		switch choix {
		case "1":
			display.DisplayInfo(arthur)
		case "2":
			display.AccessInventory(&arthur)
		case "3":
			display.Marchand(&arthur, shop)
		case "4":
			fmt.Println("Au revoir !")
			quitter = true
		default:
			fmt.Println("Choix invalide !")
		}
	}
}

func isDead(j character.Character) {
	if j.Pv <= 0 {
		fmt.Printf("%s est mort...\n", j.Name)
		// Résurrection avec 50% des PV max
		j.Pv = j.Pvmax / 2
		fmt.Printf("%s est ressuscité avec %d/%d PV !\n", j.Name, j.Pv, j.Pvmax)
	}
}
