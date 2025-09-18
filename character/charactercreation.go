package character

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
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

// Constantes de couleur et de style pour l'affichage
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Cyan   = "\033[36m"
	Bold   = "\033[1m"
	Gray   = "\033[90m"
)

// afficherSeparateur affiche une ligne de séparation stylisée.
func afficherSeparateur() {
	fmt.Println(Gray + "────────────────────────────────────────────────────────────────────────────────" + Reset)
}

// afficherTitre affiche un titre stylisé dans une boîte.
func afficherTitre(titre string) {
	afficherSeparateur()

	// On calcule la largeur visuelle en comptant les runes, pas les bytes.
	contentWidth := utf8.RuneCountInString(titre) + 4

	// Ligne supérieure de la boîte
	fmt.Printf("%s╔%s╗%s\n", Blue+Bold, strings.Repeat("═", contentWidth), Reset)

	// Ligne du milieu avec le titre
	fmt.Printf("%s║  %s%s%s  ║%s\n", Blue+Bold, Cyan+Bold, titre, Blue+Bold, Reset)

	// Ligne inférieure de la boîte
	fmt.Printf("%s╚%s╝%s\n", Blue+Bold, strings.Repeat("═", contentWidth), Reset)
}

// afficherOption affiche une option numérotée avec son texte et des détails.
func afficherOption(numero int, texte string, details string) {
	fmt.Printf(Yellow+Bold+"  %d. %s"+Reset+" %s\n", numero, texte, details)
}

// CharacterCreation guide l'utilisateur pour créer un nouveau personnage.
func CharacterCreation() Character {
	reader := bufio.NewReader(os.Stdin)
	var name string
	var initiative int // Variable pour stocker l'initiative

	// --- Étape 1 : Choix du nom ---
	afficherTitre("CRÉATION DU PERSONNAGE")
	fmt.Println()

	for {
		fmt.Print(Green + Bold + "  ➤ Entrez le nom de votre personnage : " + Reset)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(Red + Bold + "  ❌ Erreur de lecture, veuillez réessayer." + Reset)
			continue
		}
		name = strings.TrimSpace(input)
		if isAlpha(name) {
			break
		} else {
			fmt.Println(Red + Bold + "  ❌ Erreur : Le nom ne doit contenir que des lettres et ne peut pas être vide." + Reset)
		}
	}

	// Formatage du nom
	name = strings.ToLower(name)
	runes := []rune(name)
	runes[0] = unicode.ToUpper(runes[0])
	formattedName := string(runes)
	fmt.Printf(Green+Bold+"\n  ✅ Bienvenue, %s ! ✨\n"+Reset, formattedName)

	// --- Étape 2 : Choix de la classe ---
	afficherTitre("CHOIX DE LA CLASSE")
	fmt.Println()

	var className string
	var pvMax int
	classIsValid := false
	for !classIsValid {
		fmt.Println(Cyan + Bold + "  Choisissez votre classe :\n" + Reset)
		afficherOption(1, "Doom Slayer", Red+"(DPS élevé,   80 PV max)"+Reset+" - Rapide, frappe fort, mais fragile. ⚡")
		afficherOption(2, "Doom Caster", Blue+"(DPS moyen,  100 PV max)"+Reset+" - Lance des sorts dévastateurs. 🔮")
		afficherOption(3, "Doom Bastion", Green+"(Peu de DPS, 120 PV max)"+Reset+" - Encaisse les coups et protège le groupe. 🛡️")

		fmt.Print(Green + Bold + "\n  ➤ Votre choix (1-3) : " + Reset)
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)
		switch choice {
		case "1":
			className = "Doom Slayer"
			pvMax = 80
			initiative = 12 // Assignation de l'initiative pour Doom Slayer
			classIsValid = true
		case "2":
			className = "Doom Caster"
			pvMax = 100
			initiative = 8 // Assignation de l'initiative pour Doom Caster
			classIsValid = true
		case "3":
			className = "Doom Bastion"
			pvMax = 120
			initiative = 6 // Assignation de l'initiative pour Doom Bastion
			classIsValid = true
		default:
			fmt.Println(Red + Bold + "  ❌ Choix invalide. Veuillez entrer 1, 2 ou 3." + Reset)
		}
	}

	// --- Étape 3 : Finalisation ---
	afficherTitre("FINALISATION")
	fmt.Println()

	// Calcul des PV de départ (50% des PV max)
	pvCurrent := pvMax / 2
	fmt.Printf(Green+Bold+"  ✅ Vous avez choisi la classe %s.\n"+Reset, className)
	fmt.Printf(Green+Bold+"  ✅ Vous commencez l'aventure avec %d/%d PV.\n"+Reset, pvCurrent, pvMax)
	fmt.Printf(Green+Bold+"  ✅ Votre initiative de base est de %d.\n"+Reset, initiative) // Confirmation de l'initiative

	// Initialisation des compétences et de l'inventaire de base (selon la première version)
	startingSkills := []string{"Coup de Poing"}
	startingInventory := []string{"Potion de vie"}
	if className == "Doom Caster" {
		startingSkills = append(startingSkills, "Éclair de Givre") // Exemple de compétence spécifique
		startingInventory = append(startingInventory, "Mana Potion")
	} else if className == "Doom Slayer" {
		startingSkills = append(startingSkills, "Attaque Rapide")
	} else if className == "Doom Bastion" {
		startingSkills = append(startingSkills, "Provocation")
	}

	// Retourne le personnage créé en utilisant les données collectées
	return InitCharacter(
		formattedName,
		className,
		1,          // Level 1
		100,        // Starting gold (adjust as needed, previous version had a very high number)
		pvMax,      // Max HP based on class
		pvCurrent,  // Current HP (50% of max)
		10,         // Base Attack (could be class-dependent too, using a placeholder for now)
		0,          // Exp
		100,        // ExpMax
		initiative, // Assignation de l'initiative finale
		startingSkills,
		startingInventory,
	)
}
