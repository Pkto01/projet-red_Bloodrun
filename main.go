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

func spellBook(skills []string, newSpell string) []string {
	for _, spell := range skills {
		if spell == newSpell {
			return skills // le sort existe déjà, on ne fait rien
		}
	}
	return append(skills, newSpell) // sinon on l'ajoute
}

func shopitem(class string) []display.Item {
	// Objet commun à tous les magasins
	potionDeVie := display.Item{Nom: "Potion de vie", Prix: 20}

	switch class {
	case "Doom Slayer":
		fmt.Println(Yellow + "Le forgeron vous montre ses lames les plus affûtées." + Reset)
		return []display.Item{
			potionDeVie,
			{Nom: "Hache de Berserker", Prix: 150},
			{Nom: "Bottes de Célérité", Prix: 70},
			{Nom: "Gantelets de Force", Prix: 85},
		}
	case "Doom Caster":
		fmt.Println(Yellow + "L'enchanteur étale ses parchemins et artefacts mystiques." + Reset)
		return []display.Item{
			potionDeVie,
			{Nom: "Bâton d'Apprenti", Prix: 70},
			{Nom: "Robe de Mage", Prix: 60},
			{Nom: "Grimoire des Ombres", Prix: 200},
		}
	case "Doom Bastion":
		fmt.Println(Yellow + "L'armurier présente ses pièces les plus robustes." + Reset)
		return []display.Item{
			potionDeVie,
			{Nom: "Marteau Lourd", Prix: 90},
			{Nom: "Bouclier en Acier", Prix: 120},
			{Nom: "Armure de Plaques", Prix: 180},
		}
	default:
		// Un magasin par défaut si la classe n'est pas reconnue
		fmt.Println(Yellow + "Le marchand vous propose ses articles de base." + Reset)
		return []display.Item{
			potionDeVie,
			{Nom: "Dague Rouillée", Prix: 25},
			{Nom: "Tunique en Cuir", Prix: 30},
		}
	}
}

func Menu(j *character.Character) {
	quitter := false
	for !quitter {
		afficherEnTete()
		afficherOption(1, "Afficher les infos", "🧙")
		afficherOption(2, "Accéder à l'inventaire", "🎒")
		afficherOption(3, "Accéder au Marchant", "🛒")
		afficherOption(4, "Quitter le jeu", "🚪")
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
			loadingAnimation("Arrivée chez le Marchand")
			display.Marchand(j, shopitem(j.Class))
		case "4":
			loadingAnimation("Arrivée chez le Forgeron")
			display.Forgeron(j)
		case "5":
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

	player := character.CharacterCreation()

	Menu(&player)

	skills := []string{"Soin", "Téléportation"}

	// Ajouter "Boule de feu"
	skills = spellBook(skills, "Boule de feu")
	fmt.Println(skills)

	// Essayer de l'ajouter à nouveau
	skills = spellBook(skills, "Boule de feu")
	fmt.Println(skills)
}
