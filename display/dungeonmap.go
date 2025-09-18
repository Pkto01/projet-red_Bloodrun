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
