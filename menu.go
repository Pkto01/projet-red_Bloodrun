package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Cyan   = "\033[36m"
	Bold   = "\033[1m"
)

func lireEntree(texte string) string {
	fmt.Print(Cyan + texte + Reset)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func Menu() {
	quitter := false

	for !quitter {

		fmt.Println(Bold + Blue + "\n======================" + Reset)
		fmt.Println(Bold + Yellow + "      MENU PRINCIPAL  " + Reset)
		fmt.Println(Bold + Blue + "======================" + Reset)

		fmt.Println(Green + "1." + Reset + " Afficher les infos du personnage")
		fmt.Println(Green + "2." + Reset + " Accéder à l'inventaire")
		fmt.Println(Green + "3." + Reset + " Quitter")

		choix := lireEntree("\nVotre choix : ")

		switch choix {
		case "1":
			fmt.Println(Yellow + ">> Infos du personnage affichées (exemple)" + Reset)
		case "2":
			fmt.Println(Cyan + ">> Inventaire affiché (exemple)" + Reset)
		case "3":
			fmt.Println(Red + ">> Au revoir !" + Reset)
			quitter = true
		default:
			fmt.Println(Red + "Choix invalide, réessayez !" + Reset)
		}
	}
}
