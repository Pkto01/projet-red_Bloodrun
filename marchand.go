package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Joueur struct {
	Nom        string
	Argent     int
	PV         int
	PVmax      int
	Inventaire []string
}

type Item struct {
	Nom  string
	Prix int
}

// --- Fonctions Inventaire ---
func addInventory(j *Joueur, item Item) {
	j.Inventaire = append(j.Inventaire, item.Nom)
	fmt.Printf(">> Vous avez acheté : %s\n", item.Nom)
}

func accessInventory(j *Joueur) {
	fmt.Println("\n--- Inventaire ---")
	if len(j.Inventaire) == 0 {
		fmt.Println("Inventaire vide !")
		return
	}
	for i, item := range j.Inventaire {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	fmt.Printf("\nArgent restant : %d pièces\n", j.Argent)
}

// --- Marchand ---
func marchand(j *Joueur, shop []Item) {
	for {
		fmt.Println("\n=== Marchand ===")
		fmt.Printf("Argent disponible : %d pièces\n", j.Argent)
		for i, item := range shop {
			fmt.Printf("%d. %s (%d pièces)\n", i+1, item.Nom, item.Prix)
		}
		fmt.Println("0. Retour")

		choix := lireEntree("Votre choix : ")
		num, err := strconv.Atoi(choix)
		if err != nil || num < 0 || num > len(shop) {
			fmt.Println("Choix invalide.")
			continue
		}

		if num == 0 {
			return
		}

		item := shop[num-1]
		if j.Argent >= item.Prix {
			j.Argent -= item.Prix
			addInventory(j, item)
		} else {
			fmt.Println("Pas assez d'argent !")
		}
	}
}

// --- Infos joueur ---
func afficherInfos(j *Joueur) {
	fmt.Println("\n--- Infos du personnage ---")
	fmt.Printf("Nom : %s\n", j.Nom)
	fmt.Printf("PV : %d / %d\n", j.PV, j.PVmax)
	fmt.Printf("Argent : %d pièces\n", j.Argent)
	accessInventory(j)
}

// --- Entrée utilisateur ---
func lire(texte string) string {
	fmt.Print(texte)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

// --- Programme principal ---
func argent() {
	// Joueur
	joueur := Joueur{
		Nom:        "Héros",
		Argent:     100, // Le joueur commence avec 100 pièces
		PV:         70,
		PVmax:      100,
		Inventaire: []string{},
	}

	// Boutique du marchand
	shop := []Item{
		{"Potion de vie", 20},
		{"Épée en fer", 50},
		{"Bouclier en bois", 30},
		{"Arc basique", 40},
	}

	quitter := false
	for !quitter {
		fmt.Println("\n=== MENU ===")
		fmt.Println("1. Afficher infos personnage")
		fmt.Println("2. Accéder à l'inventaire")
		fmt.Println("3. Marchand")
		fmt.Println("4. Quitter")

		choix := lireEntree("Votre choix : ")

		switch choix {
		case "1":
			afficherInfos(&joueur)
		case "2":
			accessInventory(&joueur)
		case "3":
			marchand(&joueur, shop)
		case "4":
			fmt.Println("Au revoir !")
			quitter = true
		default:
			fmt.Println("Choix invalide !")
		}
	}
}
