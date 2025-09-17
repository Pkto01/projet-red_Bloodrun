package character

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// isAlpha vérifie si une chaîne de caractères ne contient que des lettres.
func isAlpha(s string) bool {
	if s == "" { // Un nom vide n'est pas valide
		return false
	}
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// CharacterCreation guide l'utilisateur pour créer un nouveau personnage.
func CharacterCreation() Character {
	reader := bufio.NewReader(os.Stdin)
	var name string

	// --- Étape 1 : Choix du nom ---
	for {
		fmt.Print("Entrez le nom de votre personnage : ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erreur de lecture, veuillez réessayer.")
			continue
		}
		name = strings.TrimSpace(input)

		if isAlpha(name) {
			break
		} else {
			fmt.Println("Erreur : Le nom ne doit contenir que des lettres et ne peut pas être vide. Réessayez.")
		}
	}

	// Formatage du nom
	name = strings.ToLower(name)
	runes := []rune(name)
	runes[0] = unicode.ToUpper(runes[0])
	formattedName := string(runes)

	fmt.Printf("\nBienvenue, %s !\n", formattedName)

	// --- Étape 2 : Choix de la classe ---
	var className string
	var pvMax int

	classIsValid := false
	for !classIsValid {
		fmt.Println("\nChoisissez votre classe :")
		fmt.Println("1. Doom Slayer  (DPS élevé,   80 PV max) - Rapide, frappe fort, mais fragile.")
		fmt.Println("2. Doom Caster  (DPS moyen,  100 PV max) - Lance des sorts dévastateurs.")
		fmt.Println("3. Doom Bastion (Peu de DPS, 120 PV max) - Encaisse les coups et protège le groupe.")
		fmt.Print("Votre choix (1-3) : ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			className = "Doom Slayer"
			pvMax = 80
			classIsValid = true
		case "2":
			className = "Doom Caster"
			pvMax = 100
			classIsValid = true
		case "3":
			className = "Doom Bastion"
			pvMax = 120
			classIsValid = true
		default:
			fmt.Println("Choix invalide. Veuillez entrer 1, 2 ou 3.")
		}
	}

	// --- Étape 3 : Finalisation ---
	// Calcul des PV de départ (50% des PV max)
	pvCurrent := pvMax / 2

	fmt.Printf("\nVous avez choisi la classe %s. Vous commencez l'aventure avec %d/%d PV.\n", className, pvCurrent, pvMax)

	// Utilise la fonction InitCharacter pour créer le personnage avec les stats choisies
	return InitCharacter(
		formattedName,
		className,                 // Classe choisie
		1,                         // Niveau 1
		10000000,                  // Argent de départ
		pvMax,                     // PV max basés sur la classe
		pvCurrent,                 // PV actuels (50%)
		[]string{"Coup de Poing"}, // Compétence de base
		[]string{"Potion de vie"}, // Inventaire de départ
	)
}
