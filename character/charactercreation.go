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
	var baseAttack int
	var baseDefense int
	var initiative int
	var mana int

	classIsValid := false
	for !classIsValid {
		// --- L'AFFICHAGE QUE J'AVAIS OUBLIÉ, MAINTENANT RÉINTÉGRÉ ---
		fmt.Println(Cyan + Bold + "  Choisissez votre classe :\n" + Reset)
		afficherOption(1, "Doom Slayer", Red+"(DPS élevé,   80 PV max)"+Reset+" - Rapide, frappe fort, mais fragile. ⚡")
		afficherOption(2, "Doom Caster", Blue+"(DPS moyen,  100 PV max)"+Reset+" - Lance des sorts dévastateurs. 🔮")
		afficherOption(3, "Doom Bastion", Green+"(Peu de DPS, 120 PV max)"+Reset+" - Encaisse les coups et protège le groupe. 🛡️")
		// -----------------------------------------------------------------

		fmt.Print(Green + Bold + "\n  ➤ Votre choix (1-3) : " + Reset)
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)
		switch choice {
		case "1":
			className = "Doom Slayer"
			pvMax = 80
			baseAttack = 12
			baseDefense = 2
			initiative = 12
			classIsValid = true
			mana = 0
		case "2":
			className = "Doom Caster"
			pvMax = 100
			baseAttack = 8
			baseDefense = 4
			initiative = 8
			classIsValid = true
			mana = 20
		case "3":
			className = "Doom Bastion"
			pvMax = 120
			baseAttack = 6
			baseDefense = 8
			initiative = 6
			classIsValid = true
			mana = 0
		default:
			fmt.Println(Red + Bold + "  ❌ Choix invalide. Veuillez entrer 1, 2 ou 3." + Reset)
		}
	}

	// --- Étape 3 : Finalisation ---
	afficherTitre("FINALISATION")
	fmt.Println()

	// Le personnage commence avec 100% de ses PV max
	pvCurrent := pvMax
	fmt.Printf(Green+Bold+"  ✅ Vous avez choisi la classe %s.\n"+Reset, className)
	fmt.Printf(Green+Bold+"  ✅ Vous commencez l'aventure avec %d/%d PV.\n"+Reset, pvCurrent, pvMax)
	fmt.Printf(Green+Bold+"  ✅ Votre initiative de base est de %d.\n"+Reset, initiative)

	// Utilise InitCharacter avec TOUS les arguments requis dans le bon ordre
	return InitCharacter(
		formattedName,
		className,
		1,                         // Level
		50,                        // Money
		pvMax,                     // Pvmax (basé sur la classe)
		pvCurrent,                 // Pv (commence avec 100% de vie)
		baseAttack,                // Attack (basé sur la classe)
		baseDefense,               // Defense (basé sur la classe)
		0,                         // Experience
		100,                       // NextLevelExp (exp nécessaire pour le niveau 2)
		[]string{"Coup de Poing"}, // skills de base
		[]string{"Potion de vie"}, // inventaire de base
		initiative,                // Initiative (basé sur la classe)
	)
}
