package fight

import (
	"fmt"
)

// Fonction pour simuler le pattern de combat du gobelin
func goblinPattern(goblinName string, targetName string, maxHP int, damagePattern []int) {
	fmt.Printf("\n%s commence le combat contre %s.\n", goblinName, targetName)
	currentHP := maxHP
	for turn := 0; turn < 3; turn++ {
		damage := damagePattern[turn]
		currentHP -= damage
		if currentHP < 0 {
			currentHP = 0
		}
		fmt.Printf("Tour %d : %s inflige %d de dégâts à %s. Points de vie restants : %d/%d\n",
			turn+1, goblinName, damage, targetName, currentHP, maxHP)
	}
}

// Fonction pour simuler le tour du joueur
func characterTurn() {
	var choice string
	fmt.Println("\n--- Tour du joueur ---")
	fmt.Println("Menu :")
	fmt.Println("1. Attaquer")
	fmt.Println("2. Inventaire")
	fmt.Println("3. Niveaux")
	fmt.Print("Choisissez une action (1, 2 ou 3) : ")
	fmt.Scanln(&choice)

	switch choice {
	case "1":
		attackName := "Attaque basique"
		damage := 5
		targetMaxHP := 50
		targetCurrentHP := targetMaxHP - damage
		fmt.Printf("Vous utilisez '%s', infligez %d dégâts. Points de vie restants de l'adversaire : %d/%d\n",
			attackName, damage, targetCurrentHP, targetMaxHP)
	case "2":
		fmt.Println("Vous consultez votre inventaire.")
	case "3":
		SelectLevel()
	default:
		fmt.Println("Choix invalide.")
	}
}

// Fonction pour sélectionner un niveau
func SelectLevel() {
	var levelChoice string
	fmt.Print("Choisissez un niveau (1-5) : ")
	fmt.Scanln(&levelChoice)

	switch levelChoice {
	case "1":
		goblinPattern("Gobelin Facile", "Personnage 1", 100, []int{50, 50, 50})
	case "2":
		goblinPattern("Gobelin Moyen", "Personnage 2", 150, []int{100, 200, 100})
	case "3":
		goblinPattern("Gobelin Difficile", "Personnage 3", 200, []int{150, 300, 150})
	case "4":
		goblinPattern("Gobelin Très Difficile", "Personnage 4", 250, []int{200, 400, 200})
	case "5":
		goblinPattern("Gobelin Extrême", "Personnage 5", 300, []int{250, 500, 250})
	default:
		fmt.Println("Choix invalide.")
	}
}
