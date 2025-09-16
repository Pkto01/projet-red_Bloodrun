package main

import (
	"fmt"
	"projet-red_Bloodrun/character"
	"projet-red_Bloodrun/display"
	"time"

	"github.com/common-nighthawk/go-figure"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Cyan   = "\033[36m"
	Bold   = "\033[1m"
	Gray   = "\033[90m"
	White  = "\033[37m"
)

func AsciText() {
	myFigure := figure.NewColorFigure("Bloodrun", "", "red", true)
	myFigure.Print()
	fmt.Println(Gray + "══════════════════════════════════════════════════════════" + Reset)
	fmt.Println(Red + Bold + "    ⚔️  Bienvenue dans l'univers sanglant de Bloodrun ⚔️" + Reset)
	fmt.Println(Gray + "══════════════════════════════════════════════════════════" + Reset)
	fmt.Println()
	time.Sleep(800 * time.Millisecond)
}

func afficherSeparateur() {
	fmt.Println(Blue + "══════════════════════════════════════════════════════════" + Reset)
}

func afficherEnTete() {
	afficherSeparateur()
	fmt.Println(Blue + Bold + "║" + Reset + Yellow + Bold + "                 🎮  MENU PRINCIPAL  🎮          " + Reset + Blue + Bold + "       ║" + Reset)
	afficherSeparateur()
}

func afficherOption(numero int, texte string, icone string) {
	fmt.Printf(Blue+Bold+"║"+Reset+" %s%s%d.%s %s %s%-25s%s %s%s\n",
		Gray, Bold, numero, Reset, icone, Green+Bold, texte, Reset, Blue+Bold, "                       ║")
}

func loadingAnimation(msg string) {
	fmt.Print(Cyan + Bold + msg + Reset)
	for i := 0; i < 3; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Print(".")
	}
	fmt.Println()
}

func Menu(j *character.Character) {
	quitter := false
	for !quitter {
		afficherEnTete()
		afficherOption(1, "Afficher les infos", "🧙")
		afficherOption(2, "Accéder à l'inventaire", "🎒")
		afficherOption(3, "Quitter le jeu", "🚪")
		afficherSeparateur()

		choix := display.LireEntree("\n" + Gray + "👉 Votre choix [" + Cyan + "1-3" + Gray + "] : " + Reset)

		switch choix {
		case "1":
			loadingAnimation("Chargement des infos")
			fmt.Println(Yellow + Bold + ">> " + Reset + "Infos du personnage :")
			fmt.Printf("🧝 Nom : %s\n", j.Name)
			fmt.Printf("⚔️ Classe : %s | 🎚️ Niveau : %d\n", j.Class, j.Level)
			fmt.Printf("❤️ PV : %d/%d\n", j.Pv, j.Pvmax)
		case "2":
			loadingAnimation("Ouverture de l'inventaire")
			fmt.Println(Cyan + Bold + ">> " + Reset + "Inventaire :")
			if len(j.Inventory) == 0 {
				fmt.Println(Gray + "Inventaire vide... 🎒" + Reset)
			} else {
				for i, item := range j.Inventory {
					fmt.Printf("  %d. %s\n", i+1, item)
				}
			}
		case "3":
			fmt.Println(Red + Bold + ">> " + Reset + "Merci d'avoir joué à Bloodrun ! 💀")
			quitter = true
		default:
			fmt.Println(Red + Bold + ">> " + Reset + "Choix invalide ❌")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func isDead(j *character.Character) {
	if j.Pv <= 0 {
		fmt.Printf("%s est mort... 💀\n", j.Name)
		// Résurrection avec 50% des PV max
		j.Pv = j.Pvmax / 2
		fmt.Printf("%s est ressuscité avec %d/%d PV ! ✨\n", j.Name, j.Pv, j.Pvmax)
	}
}

func main() {
	AsciText()

	arthur := character.InitCharacter(
		"Arthur",
		"Barbare",
		1,
		40,
		100,
		40,
		[]string{"Potion de vie", "Épée rouillée"},
		[]string{}, // compétences
	)

	Menu(&arthur)
}
