package main

import (
	"bufio"
	"fmt"
	"os"
	"projet-red_Bloodrun/character"
	"projet-red_Bloodrun/display"
	"projet-red_Bloodrun/fight"
	"strconv"
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

func AsciAccueil() {
	// Gros titre en ASCII rouge
	myFigure := figure.NewColorFigure("Bloodrun", "", "red", true)
	myFigure.Print()

	// Fond ASCII décoratif
	fmt.Println(Gray + "══════════════════════════════════════════════════════════" + Reset)
	fmt.Println(Red + Bold + "       ⚔️⚔️⚔️   LE SANG APPELLE LE SANG   ⚔️⚔️⚔️" + Reset)
	fmt.Println(Gray + "══════════════════════════════════════════════════════════" + Reset)
	fmt.Println()

	// Message pour entrer
	fmt.Println(Yellow + Bold + ">>> Appuyez sur [Entrée] pour commencer l’aventure..." + Reset)

	// Attente touche entrée
	bufio.NewReader(os.Stdin).ReadBytes('\n')
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

func handleDungeonSelection(player *character.Character) {
	display.DisplayDungeonMap() // Affiche la carte

	choixStr := display.LireEntree("Choisissez un donjon (1-5) ou 0 pour retourner : ")
	level, err := strconv.Atoi(choixStr) // Convertit le choix en nombre
	if err != nil || level < 0 || level > 5 {
		fmt.Println("Choix invalide.")
		return
	}

	if level == 0 {
		return
	}

	// Appel direct à la nouvelle fonction centralisée
	fight.StartDungeonCombat(player, level)

	isDead(player) // Vérifie si le joueur est mort après le combat
}

func spellBook(skills []string, newSpell string) []string {
	for _, spell := range skills {
		if spell == newSpell {
			return skills // le sort existe déjà, on ne fait rien
		}
	}
	return append(skills, newSpell) // sinon on l'ajoute
}

// forgeitem définit maintenant les RECETTES de fabrication pour chaque classe.
func forgeitem(class string) []display.CraftableItem {
	switch class {
	case "Doom Slayer":
		return []display.CraftableItem{
			{Nom: "Hache de Berserker", Prix: 150, Requis: map[string]int{"Acier Noirci": 4, "Fragments de Sang": 3}},
			{Nom: "Bottes de Célérité", Prix: 70, Requis: map[string]int{"Étoffe Sanglante": 3, "Os Fêlés": 2}},
			{Nom: "Gantelets de Force", Prix: 85, Requis: map[string]int{"Acier Noirci": 2, "Os Fêlés": 5}},
		}
	case "Doom Caster":
		return []display.CraftableItem{
			{Nom: "Bâton d'Apprenti", Prix: 70, Requis: map[string]int{"Os Fêlés": 4, "Fragments de Sang": 1}},
			{Nom: "Robe de Mage", Prix: 60, Requis: map[string]int{"Étoffe Sanglante": 5}},
			{Nom: "Grimoire des Ombres", Prix: 200, Requis: map[string]int{"Étoffe Sanglante": 2, "Fragments de Sang": 3}},
		}
	case "Doom Bastion":
		return []display.CraftableItem{
			{Nom: "Marteau Lourd", Prix: 90, Requis: map[string]int{"Acier Noirci": 3, "Os Fêlés": 3}},
			{Nom: "Bouclier en Acier", Prix: 120, Requis: map[string]int{"Acier Noirci": 7}},
			{Nom: "Armure de Plaques", Prix: 180, Requis: map[string]int{"Acier Noirci": 5, "Étoffe Sanglante": 3}},
		}
	default:
		// Recettes de base si la classe n'est pas reconnue
		return []display.CraftableItem{}
	}
}

// shopitem définit ce que le MARCHAND vend (potions et matériaux de fabrication).
func shopitem() []display.Item {
	return []display.Item{
		{Nom: "Potion de vie", Prix: 20},
		{Nom: "Os Fêlés", Prix: 35},
		{Nom: "Acier Noirci", Prix: 100},
		{Nom: "Étoffe Sanglante", Prix: 60},
	}
}
func Menu(j *character.Character) {
	quitter := false
	for !quitter {
		afficherEnTete()
		afficherOption(1, "Afficher les infos", "🧙")
		afficherOption(2, "Accéder à l'inventaire", "🎒")
		afficherOption(3, "Accéder au Marchant", "🛒")
		afficherOption(4, "Accéder au Forgeron", "⚒️")
		afficherOption(5, "Accéder aux Abysses", "👾")
		afficherOption(6, "Quitter le jeu", "🚪")
		afficherSeparateur()

		choix := display.LireEntree("\n" + Gray + "👉 Votre choix [" + Cyan + "1-5" + Gray + "] : " + Reset)

		switch choix {
		case "1":
			loadingAnimation("Chargement des infos")
			fmt.Println(Yellow + Bold + ">> " + Reset + "Infos du personnage :")
			fmt.Printf("🧝 Nom : %s\n", j.Name)
			fmt.Printf("⚔️ Classe : %s | 🎚️ Niveau : %d\n", j.Class, j.Level)
			fmt.Printf("❤️ PV : %d/%d\n", j.Pv, j.Pvmax)
		case "2":
			loadingAnimation("Ouverture de l'inventaire")
			display.AccessInventory(j)
		case "3":
			loadingAnimation("Arrivée chez le Marchand")
			display.Marchand(j, shopitem())
		case "4":
			loadingAnimation("Arrivée chez le Forgeron")
			display.Forgeron(j, forgeitem(j.Class))
		case "5":
			loadingAnimation("Arrivée dans les prodondeurs des abysses")
			handleDungeonSelection(j)
		case "6":
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
	// character.Boss()

	AsciAccueil()

	player := character.CharacterCreation()

	Menu(&player)
}
