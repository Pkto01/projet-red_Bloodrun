package display

import (
	"fmt"
	"projet-red_Bloodrun/character"
)

func DisplayInfo(j character.Character) {
	fmt.Println(Yellow + Bold + ">> " + Reset + "--- Fiche du Personnage ---")

	// --- Infos de Base ---
	fmt.Printf("🧝 Nom : %s\n", j.Name)
	fmt.Printf("⚔️ Classe : %s | 🎚️ Niveau : %d\n", j.Class, j.Level)

	// --- Vitals & Progression ---
	fmt.Printf("❤️ PV : %d/%d\n", j.Pv, j.Pvmax)
	fmt.Printf("🔮 Mana : %d/%d\n", j.Mana, j.Manamax)
	fmt.Printf("⭐ Exp : %d/%d\n", j.Experience, j.NextLevelExp)

	// --- Statistiques de Combat ---
	fmt.Println("\n--- Statistiques de Combat ---")
	fmt.Printf("💥 Attaque : %d\n", j.Attack)
	fmt.Printf("🛡️ Défense : %d\n", j.Defense)
	fmt.Printf("⚡ Initiative : %d\n", j.Initiative)

	// --- Équipement ---
	fmt.Println("\n--- Équipement ---")
	fmt.Printf("   Arme      : %s\n", j.Equipped.Weapon)
	fmt.Printf("   Armure    : %s\n", j.Equipped.Armor)
	fmt.Printf("   Accessoire: %s\n", j.Equipped.Accessory)

	// --- Ressources ---
	fmt.Printf("\n💰 Argent : %d pièces\n", j.Money)
}
