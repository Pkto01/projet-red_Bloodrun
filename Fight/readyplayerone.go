package fight

import (
	"fmt"
	"projet-red_Bloodrun/character"
	"time"
)

// GoblinPattern gère un combat contre un gobelin suivant un schéma d'attaque précis.
func GoblinPattern(player *character.Character, goblin *Monster) {
	fmt.Println("\n💥💥💥 LE COMBAT COMMENCE ! 💥💥💥")
	turn := 1

	// Le combat continue tant que les deux combattants sont en vie.
	for player.Pv > 0 && goblin.Pv > 0 {
		fmt.Printf("\n---------- TOUR %d ----------\n", turn)

		// --- Tour du Joueur ---
		playerDamage := player.Attack - goblin.Defense
		if playerDamage < 0 {
			playerDamage = 0 // On ne peut pas infliger de dégâts négatifs
		}
		goblin.Pv -= playerDamage
		fmt.Printf("Vous infligez à %s %d de dégâts.\n", goblin.Name, playerDamage)
		fmt.Printf("Points de vie du Gobelin : %d/%d PV\n", goblin.Pv, goblin.Pvmax)

		if goblin.Pv <= 0 {
			break // Si le gobelin est vaincu, on sort de la boucle
		}

		time.Sleep(1 * time.Second) // Petite pause pour la lisibilité

		// --- Tour du Gobelin ---
		damageMultiplier := 1.0 // 100% par défaut

		// Tous les 3 tours, le multiplicateur passe à 200%
		if turn%3 == 0 {
			damageMultiplier = 2.0 // 200%
			fmt.Println("Le Gobelin concentre son énergie pour une attaque puissante !")
		}

		// Calcul des dégâts
		baseDamage := float64(goblin.Attack) * damageMultiplier
		inflictedDamage := int(baseDamage) - player.Defense
		if inflictedDamage < 0 {
			inflictedDamage = 0
		}

		// Application des dégâts au joueur
		player.Pv -= inflictedDamage

		// Affichage des informations de combat
		fmt.Printf("%s inflige à %s %d de dégâts.\n", goblin.Name, player.Name, inflictedDamage)
		fmt.Printf("Vos points de vie : %d/%d PV\n", player.Pv, player.Pvmax)

		turn++
		time.Sleep(1 * time.Second)
	}

	// --- Fin du Combat ---
	fmt.Println("\n---------- FIN DU COMBAT ----------")
	if player.Pv <= 0 {
		fmt.Println("Vous avez été vaincu... 💀")
		// Vous pouvez appeler votre fonction isDead ici si vous le souhaitez
	} else {
		fmt.Printf("Vous avez vaincu le %s !\n", goblin.Name)
		// Ici, vous pourriez ajouter une récompense (argent, expérience, etc.)
		player.Money += 25
		fmt.Println("Vous gagnez 25 pièces d'or.")
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
