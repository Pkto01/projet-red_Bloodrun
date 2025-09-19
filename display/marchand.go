package display

import (
	"fmt"
	"projet-red_Bloodrun/character"
	"strconv"
	"strings"
)

type ShopableItem interface {
	GetName() string
	GetPrice() int
	IsSpell() bool
}

// Item est utilisé pour les objets simples (achat/vente)
type Item struct {
	Nom  string
	Prix int
}

type ItemSpell struct {
	Nom  string
	Prix int
}

// CraftableItem est utilisé pour les recettes du forgeron
type CraftableItem struct {
	Nom    string
	Prix   int
	Requis map[string]int // Map[nom_materiau]quantité
}

func (i Item) GetName() string {
	return i.Nom
}

func (i Item) GetPrice() int {
	return i.Prix
}

func (i Item) IsSpell() bool {
	return false // Un Item n'est pas un sort
}

// --- Méthodes pour la struct ItemSpell ---

func (is ItemSpell) GetName() string {
	return is.Nom
}

func (is ItemSpell) GetPrice() int {
	return is.Prix
}

func (is ItemSpell) IsSpell() bool {
	return true // Un ItemSpell est un sort
}

// --- Fonctions d'aide à l'inventaire ---

// removeFromSliceByIndex retire un élément d'une tranche à un index donné.
func RemoveFromSliceByIndex(slice []string, index int) []string {
	if index < 0 || index >= len(slice) {
		return slice // Ne fait rien si l'index est invalide
	}
	return append(slice[:index], slice[index+1:]...)
}

// countItem compte combien de fois un objet spécifique apparaît dans l'inventaire.
func countItem(inventory []string, itemName string) int {
	count := 0
	for _, item := range inventory {
		if item == itemName {
			count++
		}
	}
	return count
}

// removeItems retire un certain nombre d'un objet spécifique de l'inventaire.
func removeItems(inventory []string, itemName string, quantity int) []string {
	var newInventory []string
	removedCount := 0
	for _, item := range inventory {
		if item == itemName && removedCount < quantity {
			removedCount++
		} else {
			newInventory = append(newInventory, item)
		}
	}
	return newInventory
}

func getInventoryLimit(class string, upgrades int) int {
	base := 0
	switch class {
	case "Doom Caster":
		base = 7
	case "Doom Slayer":
		base = 11
	case "Doom Bastion":
		base = 15
	default:
		base = 10
	}
	return base + (upgrades * 10)
}

func upgradeInventorySlot(j *character.Character) {
	if j.InventoryUpgrades >= 3 {
		fmt.Println("\033[31mVous avez déjà atteint la limite d'augmentations d'inventaire (3 fois)!\033[0m")
		return
	}
	j.InventoryUpgrades++
	fmt.Println("\033[32mVotre capacité d'inventaire a été augmentée de +10 !\033[0m")
}

func AddInventory(j *character.Character, itemName string) {
	limit := getInventoryLimit(j.Class, j.InventoryUpgrades)
	if len(j.Inventory) >= limit {
		fmt.Println("\033[31mInventaire plein ! Impossible d'ajouter : " + itemName + "\033[0m")
		return
	}
	j.Inventory = append(j.Inventory, itemName)

}

func AddSpell(c *character.Character, spellName string) {
	var allSpellsInGame = map[string]character.Spell{
		"Flammèche":          {Name: "Flammèche", Damage: 10, Mana: 10},
		"Frappe Puissante":   {Name: "Frappe Puissante", Damage: 12, Mana: 3},
		"Coup de Bouclier":   {Name: "Coup de Bouclier", Damage: 8, Mana: 5},
		"Entaille Sanglante": {Name: "Entaille Sanglante", Damage: 150, Mana: 7},
		"Ruée Des Enfers":    {Name: "Ruée Des Enfers", Damage: 250, Mana: 8},
		"Boule De Feu":       {Name: "Boule De Feu", Damage: 150, Mana: 17},
		"Blizzard Infernal":  {Name: "Blizzard Infernal", Damage: 250, Mana: 25},
		"Protection Du Bastion": {
			Name:        "Protection Du Bastion",
			Mana:        8,
			Effect:      "DefenseBuff", // Type d'effet
			EffectValue: 20,            // Augmente la défense de 20
			Duration:    4,             // Dure 3 tours après le tour où il est lancé
		},
		"Contre Attaque": {
			Name:        "Contre Attaque",
			Mana:        10,
			Effect:      "CounterAttack", // Type d'effet
			EffectValue: 50,              // Renvoie 50% des dégâts
			Duration:    2,               // Actif pour le prochain tour du monstre
		},
		// Ajoutez ici TOUS les sorts qui existent dans votre jeu
	}

	spellToAdd, exists := allSpellsInGame[spellName]
	if !exists {
		fmt.Printf("Erreur : Tentative d'ajout d'un sort inconnu : %s\n", spellName)
		return // Le sort n'existe pas, on ne fait rien.
	}

	// ÉTAPE B : Vérifier si le personnage connaît déjà ce sort.
	for _, existingSpell := range c.Spells {
		if existingSpell.Name == spellToAdd.Name {
			fmt.Printf("Vous connaissez déjà le sort '%s'.\n", spellToAdd.Name)
			return // On arrête, car le sort est déjà connu.
		}
	}

	// ÉTAPE C : Ajouter le sort à la liste des sorts du personnage.
	c.Spells = append(c.Spells, spellToAdd)
	fmt.Printf("Nouveau sort appris : %s !\n", spellToAdd.Name)
}

// getItemSalePrice retourne le prix de vente d'un objet (50% du prix de base).
func getItemSalePrice(itemName string) int {
	var basePrice int
	switch itemName {
	// Matériaux
	case "Fragments de Sang":
		basePrice = 50
	case "Os Fêlés":
		basePrice = 35
	case "Acier Noirci":
		basePrice = 100
	case "Étoffe Sanglante":
		basePrice = 60
	// Consommables
	case "Potion de vie":
		basePrice = 20
	case "Potion de mana":
		basePrice = 30
	case "Potion de poison":
		basePrice = 30
	// Équipements craftés
	case "Hache de Berserker", "Armure de Plaques":
		basePrice = 180 // Prix arbitraire élevé pour les items de haut niveau
	case "Bouclier en Acier", "Grimoire des Ombres":
		basePrice = 200
	default:
		basePrice = 10 // Prix par défaut pour les objets non listés
	}
	return basePrice / 2 // Le prix de vente est la moitié du prix de base
}

// sellItem gère le menu et la logique de vente d'objets.
func sellItem(j *character.Character) {
	if len(j.Inventory) == 0 {
		fmt.Println("\033[90mVotre inventaire est vide, vous n'avez rien à vendre.\033[0m")
		return
	}

	for {
		fmt.Println("\n=== Vendre un objet ===")
		fmt.Printf("Argent disponible : %d pièces\n", j.Money)
		fmt.Println("Quel objet souhaitez-vous vendre ?")

		// Affiche l'inventaire avec les prix de vente
		for i, item := range j.Inventory {
			salePrice := getItemSalePrice(item)
			fmt.Printf("%d. %s (Vendre pour %d pièces)\n", i+1, item, salePrice)
		}
		fmt.Println("0. Retour")

		choixStr := LireEntree("Votre choix : ")
		choix, err := strconv.Atoi(choixStr)
		if err != nil || choix < 0 || choix > len(j.Inventory) {
			fmt.Println("Choix invalide.")
			continue
		}

		if choix == 0 {
			return // Quitte la fonction de vente
		}

		itemIndex := choix - 1
		itemToSell := j.Inventory[itemIndex]
		salePrice := getItemSalePrice(itemToSell)

		// Met à jour l'argent et l'inventaire
		j.Money += salePrice
		j.Inventory = RemoveFromSliceByIndex(j.Inventory, itemIndex)

		fmt.Printf("\033[32mVous avez vendu %s pour %d pièces.\033[0m\n", itemToSell, salePrice)

		// Si l'inventaire devient vide, on quitte le menu de vente
		if len(j.Inventory) == 0 {
			fmt.Println("\033[90mVous n'avez plus rien à vendre.\033[0m")
			return
		}
	}
}

// Marchand gère l'achat et l'accès au menu de vente.
func Marchand(j *character.Character, shop []ShopableItem) {
	for {
		fmt.Println("\n=== Marchand ===")
		fmt.Printf("Argent disponible : %d pièces\n", j.Money)
		for i, item := range shop {
			fmt.Printf("%d. %s (%d pièces)\n", i+1, item.GetName(), item.GetPrice())
		}
		fmt.Printf("%d. Augmentation d'inventaire\n", len(shop)+1)
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
			item := shop[choix-1] // item est maintenant de type ShopableItem

			if j.Money >= item.GetPrice() {
				j.Money -= item.GetPrice()

				// Vérifie si l'objet est un sort
				if item.IsSpell() {
					// Si c'est un sort, on l'ajoute au grimoire du joueur
					AddSpell(j, item.GetName()) // Assurez-vous d'avoir la méthode AddSpell
					fmt.Printf("\033[36m>> Vous avez appris le sort : %s\033[0m\n", item.GetName())
				} else {
					// Sinon, c'est un objet normal, on l'ajoute à l'inventaire
					if item.GetName() == "Augmentation d'inventaire" {
						upgradeInventorySlot(j)
					} else {
						AddInventory(j, item.GetName())
						fmt.Printf("\033[32m>> Vous avez acheté : %s\033[0m\n", item.GetName())
					}
				}
			} else {
				fmt.Println("\033[31mPas assez d'argent !\033[0m")
			}
		} else {
			fmt.Println("Choix invalide.")
		}

	}
}

// --- Forgeron ---
func Forgeron(j *character.Character, recipes []CraftableItem) {
	for {
		fmt.Println("\n=== Forgeron ===")
		fmt.Printf("Argent disponible : %d pièces\n", j.Money)
		fmt.Println("Objets fabricables :")

		// --- BOUCLE D'AFFICHAGE ---
		// Affiche toutes les recettes et colore les matériaux requis.
		for i, recipe := range recipes {
			var reqs []string
			// La logique de couleur suffit pour l'affichage.
			for material, quantity := range recipe.Requis {
				playerHas := countItem(j.Inventory, material)
				color := "\033[32m" // Vert par défaut
				if playerHas < quantity {
					color = "\033[31m" // Rouge si matériaux manquants
				}
				reqs = append(reqs, fmt.Sprintf("%s%d %s (%d/%d)\033[0m", color, quantity, material, playerHas, quantity))
			}
			fmt.Printf("%d. %s (%d pièces) - Requis: %s\n", i+1, recipe.Nom, recipe.Prix, strings.Join(reqs, ", "))
		}
		fmt.Println("0. Retour")

		choix := LireEntree("Votre choix : ")
		num, err := strconv.Atoi(choix)
		if err != nil || num < 0 || num > len(recipes) {
			fmt.Println("Choix invalide.")
			continue
		}

		if num == 0 {
			return
		}

		recipe := recipes[num-1]

		// --- LOGIQUE DE FABRICATION ---
		canCraft := true
		if j.Money < recipe.Prix {
			fmt.Println("\033[31mPas assez d'argent !\033[0m")
			canCraft = false
		}
		for material, quantity := range recipe.Requis {
			if countItem(j.Inventory, material) < quantity {
				fmt.Printf("\033[31mMatériaux manquants : %s\033[0m\n", material)
				canCraft = false
			}
		}

		// Si toutes les conditions sont remplies (canCraft est resté true)
		if canCraft {
			j.Money -= recipe.Prix
			for material, quantity := range recipe.Requis {
				j.Inventory = removeItems(j.Inventory, material, quantity)
			}
			AddInventory(j, recipe.Nom)
			fmt.Printf("\033[32mVous avez fabriqué : %s !\033[0m\n", recipe.Nom)
		}
	}
}
