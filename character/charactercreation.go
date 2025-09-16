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

func afficherSeparateur() {
	fmt.Println(Gray + "────────────────────────────────────────────────────────────────────────────────" + Reset)
}

func afficherTitre(titre string) {
	afficherSeparateur()
	// La logique ici pour construire les bordures avec des caractères.
	// Pour un affichage précis comme l'image, il faudrait une bibliothèque graphique ou terminale avancée.
	// Dans un terminal texte pur, on ne peut que simuler.
	fmt.Printf("%s%s %s %s%s\n", Blue+Bold, "╔", strings.Repeat("═", len(titre)+4), Bold, "╗")
	fmt.Printf("%s %s%s%s %s%s\n", Blue+Bold, "║", Cyan+Bold, titre, strings.Repeat(" ", len(titre)+1), Blue+Bold+"║")
	fmt.Printf("%s%s %s %s%s\n", Blue+Bold, "╚", strings.Repeat("═", len(titre)+4), Bold, "╝")
}

func afficherOption(numero int, texte string, details string) {
	fmt.Printf(Yellow+Bold+"  %d. %s"+Reset+" %s\n", numero, texte, details)
}

func CharacterCreation() Character {
	reader := bufio.NewReader(os.Stdin)
	var name string

	// --- Étape 1 : Choix du nom ---
	afficherTitre("CRÉATION DU PERSONNAGE") // <-- Ici s'afficherait la partie supérieure du cadre de l'image 1
	fmt.Println()

	// Cette section correspond visuellement à l'entrée du nom dans l'image 1
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
	fmt.Printf(Green+Bold+"\n  ✅ Bienvenue, %s ! ✨\n"+Reset, formattedName) // <-- Cette ligne correspond à la confirmation "Bienvenue, Darkvador !" dans l'image 1

	// --- Étape 2 : Choix de la classe ---
	afficherTitre("CHOIX DE LA CLASSE") // <-- Ici s'afficherait le cadre "CHOIX DE LA CLASSE" de l'image 1
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
			fmt.Println(Red + Bold + "  ❌ Choix invalide. Veuillez entrer 1, 2 ou 3." + Reset)
		}
	}

	// --- Étape 3 : Finalisation ---
	afficherTitre("FINALISATION") // <-- Ici s'afficherait le cadre "FINALISATION" de l'image 1
	fmt.Println()

	// Calcul des PV de départ (50% des PV max)
	pvCurrent := pvMax / 2
	fmt.Printf(Green+Bold+"  ✅ Vous avez choisi la classe %s.\n"+Reset, className)
	fmt.Printf(Green+Bold+"  ✅ Vous commencez l'aventure avec %d/%d PV.\n"+Reset, pvCurrent, pvMax)

	// Utilise la fonction InitCharacter pour créer le personnage avec les stats choisies
	return InitCharacter(
		formattedName,
		className,                 // Classe choisie
		1,                         // Niveau 1
		50,                        // Argent de départ
		pvMax,                     // PV max basés sur la classe
		pvCurrent,                 // PV actuels (50%)
		[]string{"Coup de Poing"}, // Compétence de base
		[]string{"Potion de vie"}, // Inventaire de départ
	)
}
