package display

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

// wrapText coupe un texte en plusieurs lignes sans couper les mots.
func wrapText(text string, maxWidth int) []string {
	var lines []string
	words := strings.Fields(text)
	if len(words) == 0 {
		return []string{}
	}

	currentLine := words[0]
	for _, word := range words[1:] {
		if utf8.RuneCountInString(currentLine)+1+utf8.RuneCountInString(word) > maxWidth {
			lines = append(lines, currentLine)
			currentLine = word
		} else {
			currentLine += " " + word
		}
	}
	lines = append(lines, currentLine)
	return lines
}

// padText ajoute des espaces à droite d'un texte pour atteindre une certaine largeur.
func padText(text string, width int) string {
	textWidth := utf8.RuneCountInString(text)
	if textWidth >= width {
		return text
	}
	return text + strings.Repeat(" ", width-textWidth)
}

// DisplayDungeonMap affiche la carte compacte avec les titres sur plusieurs lignes.
func DisplayDungeonMap() {
	// --- Configuration ---
	titles := []string{
		"Couloir Saignant",
		"Les Fosses de Chair",
		"Sanctuaire des Os Brisés",
		"La Forge Écarlate",
		"Trône du Seigneur Sanglant",
	}

	boxWidth := 18 // Largeur interne FIXE pour chaque colonne
	numBoxes := len(titles)

	// 1. Couper tous les titres en lignes et trouver la hauteur max
	wrappedTitles := [][]string{}
	maxHeight := 0
	for _, title := range titles {
		wrapped := wrapText(title, boxWidth)
		wrappedTitles = append(wrappedTitles, wrapped)
		if len(wrapped) > maxHeight {
			maxHeight = len(wrapped)
		}
	}

	// --- Dessin de la carte ---

	// Ligne supérieure
	fmt.Print(Blue + Bold + "╔" + strings.Repeat("═", boxWidth))
	for i := 1; i < numBoxes; i++ {
		fmt.Print("╦" + strings.Repeat("═", boxWidth))
	}
	fmt.Println("╗" + Reset)

	// Ligne des numéros (centrés)
	fmt.Print(Blue + Bold + "║" + Reset)
	for i := 0; i < numBoxes; i++ {
		numStr := strconv.Itoa(i + 1)
		padding := (boxWidth - len(numStr)) / 2
		rightPadding := boxWidth - len(numStr) - padding
		fmt.Print(strings.Repeat(" ", padding) + Cyan + Bold + numStr + Reset + strings.Repeat(" ", rightPadding))
		fmt.Print(Blue + Bold + "║" + Reset)
	}
	fmt.Println()

	// Ligne de séparation
	fmt.Print(Blue + Bold + "╠" + strings.Repeat("═", boxWidth))
	for i := 1; i < numBoxes; i++ {
		fmt.Print("╬" + strings.Repeat("═", boxWidth))
	}
	fmt.Println("╣" + Reset)

	// Lignes des titres (hauteur dynamique)
	for i := 0; i < maxHeight; i++ {
		fmt.Print(Blue + Bold + "║" + Reset) // Bordure gauche
		for j := 0; j < numBoxes; j++ {
			var lineContent string
			if i < len(wrappedTitles[j]) {
				lineContent = wrappedTitles[j][i] // On prend la ligne si elle existe
			} else {
				lineContent = "" // Sinon, c'est une ligne vide
			}
			fmt.Print(Gray + Bold + padText(lineContent, boxWidth) + Reset)
			fmt.Print(Blue + Bold + "║" + Reset) // Séparateur
		}
		fmt.Println()
	}

	// Ligne inférieure
	fmt.Print(Blue + Bold + "╚" + strings.Repeat("═", boxWidth))
	for i := 1; i < numBoxes; i++ {
		fmt.Print("╩" + strings.Repeat("═", boxWidth))
	}
	fmt.Println("╝" + Reset)
}

for {
		fmt.Println("\n=== Marchand ===")
		fmt.Printf("Argent disponible : %d pièces\n", j.Money)
		for i, item := range shop {
			fmt.Printf("%d. %s (%d pièces)\n", i+1, item.Nom, item.Prix)
		}
		fmt.Printf("%d. Augmentation d'inventaire", len(shop)+1)
		fmt.Printf("%d. Vendre un objet\n", len(shop)+2)
		fmt.Println("0. Retour")

		choixStr := LireEntree("Votre choix : ")
		choix, err := strconv.Atoi(choixStr)
		if err != nil {
			fmt.Println("Choix invalide.")
			continue
		}

		if choix == 0 {
			return
		}

		if choix == len(shop)+1 {
			upgradeInventorySlot(j)
			continue
		}

		if choix == len(shop)+2 { // Si l'utilisateur choisit l'option "Vendre"
			sellItem(j)
			continue // Revient au menu du marchand après la vente
		}

		// Logique d'achat
		if choix > 0 && choix <= len(shop) {
			item := shop[choix-1]
			if j.Money >= item.Prix {
				j.Money -= item.Prix
				addInventory(j, item.Nom)
				fmt.Printf("\033[32m>> Vous avez acheté : %s\033[0m\n", item.Nom)
			} else {
				fmt.Println("\033[31mPas assez d'argent !\033[0m")
			}
		} else {
			fmt.Println("Choix invalide.")
		}

		if choix > 0 && choix <= len(shop) {
			item := shop[choix-1]
			if j.Money >= item.Prix {
				j.Money -= item.Prix
				if item.Nom == "Augmentation d'inventaire" {
					upgradeInventorySlot(j)
				} else {
					addInventory(j, item.Nom)
				}
				fmt.Printf("\033[32m>> Vous avez acheté : %s\033[0m\n", item.Nom)
			} else {
				fmt.Printf("\033[31mPas assez d'argent !\033[0m")
			}
		}

	}