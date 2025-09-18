package fight

import (
	"fmt"
	"math/rand"
	"projet-red_Bloodrun/character"
	"projet-red_Bloodrun/display"
	"time"
)

func playerTurn(player *character.Character, adversary *Monster) (combatOver bool) {
	turnOver := false
	for !turnOver {
		fmt.Println("\n--- C'est votre tour ---")
		fmt.Println("1. Attaquer")
		fmt.Println("2. Inventaire")
		fmt.Println("3. Lancer un sort")
		fmt.Println("4. Fuir")
		choix := display.LireEntree("Votre choix : ")

		switch choix {
		case "1": // ATTAQUER
			playerDamage := player.Attack - adversary.Defense
			if playerDamage < 1 {
				playerDamage = 1 // Toujours infliger au moins 1 de dégât
			}
			adversary.Pv -= playerDamage
			fmt.Printf("Vous attaquez et infligez à %s %d de dégâts.\n", adversary.Name, playerDamage)
			turnOver = true // Attaquer termine le tour

		case "2": // INVENTAIRE
			display.AccessInventory(player)
			// Après avoir utilisé l'inventaire (ex: une potion), le joueur peut encore agir.
			// On ne met PAS turnOver à true, la boucle du menu se relance.

		case "3":
			fmt.Println("📖 Choisissez un sort :")
			for i, spell := range player.Spells {
				fmt.Printf("%d. %s (Dégâts: %d | Soin: %d | Mana: %d)\n",
					i+1, spell.Name, spell.Damage, spell.Heal, spell.Mana)
			}

			var spellChoice int
			fmt.Scanln(&spellChoice)

			if spellChoice < 1 || spellChoice > len(player.Spells) {
				fmt.Println("❌ Choix invalide, tour perdu...")
				return
			}

			spell := player.Spells[spellChoice-1]

			// Vérification du mana
			if player.Mana < spell.Mana {
				fmt.Printf("❌ Pas assez de mana ! (%d/%d requis)\n", player.Mana, spell.Mana)
				return
			}

			// Consommation
			player.Mana -= spell.Mana

			if spell.Damage > 0 {
				fmt.Printf("%s lance %s et inflige %d dégâts à %s !\n",
					player.Name, spell.Name, spell.Damage, adversary.Name)
				adversary.Pv -= spell.Damage
			}

			if spell.Heal > 0 {
				player.Pv += spell.Heal
				if player.Pv > player.Pvmax {
					player.Pv = player.Pvmax
				}
				fmt.Printf("%s lance %s et récupère %d PV (PV: %d/%d) !\n",
					player.Name, spell.Name, spell.Heal, player.Pv, player.Pvmax)
			}

			fmt.Printf("💧 Mana restant : %d/%d\n", player.Mana, player.Manamax)
			turnOver = true

		case "4": // FUIR
			fmt.Println("Vous essayez de prendre la fuite...")
			time.Sleep(1 * time.Second)

			// Chance de fuite de 50%
			if rand.Intn(2) == 0 {
				fmt.Println("Vous avez réussi à fuir !")
				return true // Signale que le combat est terminé
			} else {
				fmt.Println("Votre tentative de fuite a échoué ! Le monstre vous attaque.")
				turnOver = true // La tentative de fuite ratée termine le tour
			}

		default:
			fmt.Println("Choix invalide.")
		}
	}
	return false // Le combat n'est pas terminé
}

func determineFirstAttacker(player *character.Character, adversary *Monster) bool {
	playerFirst := true
	if adversary.Initiative > player.Initiative {
		fmt.Printf("%s prend l'initiative et attaque en premier ! ⚡\n", adversary.Name)
		playerFirst = false
	} else if player.Initiative > adversary.Initiative {
		fmt.Printf("%s agit plus vite et attaque en premier ! ⚡\n", player.Name)
	} else {
		// En cas d'égalité, on peut décider aléatoirement ou donner la priorité au joueur
		if rand.Intn(2) == 0 { // 50% de chance pour le joueur d'attaquer en premier
			fmt.Printf("Les deux combattants sont aussi rapides, mais %s prend l'avantage ! ⚡\n", player.Name)
			playerFirst = true
		} else {
			fmt.Printf("Les deux combattants sont aussi rapides, mais %s prend l'avantage ! ⚡\n", adversary.Name)
			playerFirst = false
		}
	}
	time.Sleep(2 * time.Second) // Pause pour laisser le joueur lire le message d'initiative
	return playerFirst
}

// combatLoop est la boucle de base pour tous les combats.
func combatLoop(player *character.Character, adversary *Monster, monsterAction func(int, *Monster, *character.Character)) {
	DisplayMonsterArt(adversary.Name)
	fmt.Println("\n💥💥💥 LE COMBAT COMMENCE ! 💥💥💥")

	playerFirst := determineFirstAttacker(player, adversary)

	turn := 1
	hasFled := false

	for player.Pv > 0 && adversary.Pv > 0 {
		fmt.Printf("\n---------- TOUR %d ----------\n", turn)
		fmt.Printf("PV Joueur: %d/%d | PV %s: %d/%d\n", player.Pv, player.Pvmax, adversary.Name, adversary.Pv, adversary.Pvmax)

		// Gestion de l'ordre des tours selon l'initiative
		if playerFirst {
			// --- Tour du Joueur ---
			if playerTurn(player, adversary) {
				hasFled = true
				break // Si playerTurn retourne true, le joueur a fui, on arrête le combat.
			}
			// On vérifie si le monstre est mort après le tour du joueur
			if adversary.Pv <= 0 {
				break
			}
			time.Sleep(1 * time.Second)

			// --- Tour du Monstre ---
			monsterAction(turn, adversary, player)
		} else { // C'est le tour du monstre en premier
			// --- Tour du Monstre ---
			monsterAction(turn, adversary, player)
			// On vérifie si le joueur est mort après le tour du monstre
			if player.Pv <= 0 {
				break
			}
			time.Sleep(1 * time.Second)

			// --- Tour du Joueur ---
			if playerTurn(player, adversary) {
				hasFled = true
				break // Si playerTurn retourne true, le joueur a fui, on arrête le combat.
			}
		}

		turn++
		time.Sleep(1 * time.Second)
	}

	// --- Fin du Combat ---
	fmt.Println("\n---------- FIN DU COMBAT ----------")
	if hasFled {
		fmt.Println("Vous êtes retourné à l'entrée du donjon.")
	} else if player.Pv <= 0 {
		fmt.Println("Vous avez été vaincu... 💀")
	} else {
		fmt.Printf("Vous avez vaincu : %s !\n", adversary.Name)
		// player.GainExperience(adversary.ExperienceReward) // Décommentez quand vous aurez le système d'XP
	}

	// --- Fin du Combat ---
	fmt.Println("\n---------- FIN DU COMBAT ----------")
	if hasFled {
		fmt.Println("Vous êtes retourné à l'entrée du donjon.")
	} else if player.Pv <= 0 {
		fmt.Println("Vous avez été vaincu... 💀")
	} else {
		fmt.Printf("Vous avez vaincu : %s !\n", adversary.Name)

		// --- Attribution des récompenses ---
		fmt.Printf("Vous gagnez %d pièces d'or.\n", adversary.GoldReward)
		player.Money += adversary.GoldReward

		fmt.Printf("Vous obtenez : %s.\n", adversary.LootDrop)
		display.AddInventory(player, adversary.LootDrop)

		fmt.Printf("Vous gagnez %d points d'expérience.\n", adversary.ExperienceReward)
		player.GainExperience(adversary.ExperienceReward)
	}
}

// --- PATTERNS SPÉCIFIQUES POUR CHAQUE MONSTRE ---

// Pattern 1: Goule Sanguine - Attaques simples.
func GouleSanguinePattern(player *character.Character, goule *Monster) {
	combatLoop(player, goule, func(turn int, m *Monster, p *character.Character) {
		damage := m.Attack - p.Defense
		if damage < 1 {
			damage = 1
		}
		p.Pv -= damage
		fmt.Printf("%s vous mord et inflige %d de dégâts.\n", m.Name, damage)
		fmt.Printf("Vos points de vie : %d/%d PV\n", p.Pv, p.Pvmax)
	})
}

// Pattern 2: Abomination de Chair - Charge une attaque dévastatrice.
func AbominationPattern(player *character.Character, abomination *Monster) {
	combatLoop(player, abomination, func(turn int, m *Monster, p *character.Character) {
		var damage int
		if turn%3 == 0 {
			fmt.Println("L'Abomination de Chair lève son énorme poing et s'écrase sur vous !")
			damage = (m.Attack * 2) - p.Defense // Attaque puissante
		} else {
			damage = m.Attack - p.Defense
		}

		if damage < 1 {
			damage = 1
		}
		p.Pv -= damage
		fmt.Printf("%s vous inflige %d de dégâts.\n", m.Name, damage)
		fmt.Printf("Vos points de vie : %d/%d PV\n", p.Pv, p.Pvmax)
	})
}

// Pattern 3: Gardien Squelette - Se régénère.
func SquelettePattern(player *character.Character, squelette *Monster) {
	combatLoop(player, squelette, func(turn int, m *Monster, p *character.Character) {
		if turn%4 == 0 {
			heal := 15
			m.Pv += heal
			if m.Pv > m.Pvmax {
				m.Pv = m.Pvmax
			}
			fmt.Printf("Le %s utilise des fragments d'os pour se réparer et regagne %d PV !\n", m.Name, heal)
		} else {
			damage := m.Attack - p.Defense
			if damage < 1 {
				damage = 1
			}
			p.Pv -= damage
			fmt.Printf("%s vous frappe avec son épée rouillée et inflige %d de dégâts.\n", m.Name, damage)
			fmt.Printf("Vos points de vie : %d/%d PV\n", p.Pv, p.Pvmax)
		}
	})
}

// Pattern 4: Golem de la Forge - Vous inflige une brûlure.
func GolemPattern(player *character.Character, golem *Monster) {
	burnTurns := 0
	combatLoop(player, golem, func(turn int, m *Monster, p *character.Character) {
		if burnTurns > 0 {
			burnDamage := 5
			p.Pv -= burnDamage
			fmt.Printf("Vous subissez %d de dégâts de brûlure. Tours restants : %d\n", burnDamage, burnTurns-1)
			burnTurns--
		}

		if turn%3 == 0 {
			fmt.Printf("Le %s crache des scories en fusion !\n", m.Name)
			burnTurns = 3
			fmt.Println("Vous êtes en feu !")
		}

		damage := m.Attack - p.Defense
		if damage < 1 {
			damage = 1
		}
		p.Pv -= damage
		fmt.Printf("%s vous assène un coup brûlant et inflige %d de dégâts.\n", m.Name, damage)
		fmt.Printf("Vos points de vie : %d/%d PV\n", p.Pv, p.Pvmax)
	})
}

// Pattern 5: Seigneur Sanglant - Vole de la vie et s'énerve.
func SeigneurSanglantPattern(player *character.Character, boss *Monster) {
	isEnraged := false
	combatLoop(player, boss, func(turn int, m *Monster, p *character.Character) {
		if m.Pv < m.Pvmax/2 && !isEnraged {
			fmt.Println("Le Seigneur Sanglant entre dans une rage folle ! Sa puissance augmente !")
			m.Attack += 10
			isEnraged = true
		}

		if turn%4 == 0 {
			fmt.Printf("Le %s vous draine la vie !\n", m.Name)
			damage := m.Attack - p.Defense
			if damage < 1 {
				damage = 1
			}
			p.Pv -= damage
			heal := damage / 2
			m.Pv += heal
			if m.Pv > m.Pvmax {
				m.Pv = m.Pvmax
			}
			fmt.Printf("Il inflige %d de dégâts et se soigne de %d PV !\n", damage, heal)
		} else {
			damage := m.Attack - p.Defense
			if damage < 1 {
				damage = 1
			}
			p.Pv -= damage
			fmt.Printf("%s vous frappe violemment et inflige %d de dégâts.\n", m.Name, damage)
		}
		fmt.Printf("Vos points de vie : %d/%d PV\n", p.Pv, p.Pvmax)
	})
}

// --- POINT D'ENTRÉE POUR LES COMBATS ---

// StartDungeonCombat initialise le bon monstre et lance le combat correspondant.
func StartDungeonCombat(player *character.Character, level int) {
	switch level {
	case 1:
		adversary := InitGouleSanguine()
		GouleSanguinePattern(player, &adversary)
	case 2:
		adversary := InitAbominationDeChair()
		AbominationPattern(player, &adversary)
	case 3:
		adversary := InitGardienSquelette()
		SquelettePattern(player, &adversary)
	case 4:
		adversary := InitGolemDeLaForge()
		GolemPattern(player, &adversary)
	case 5:
		adversary := InitSeigneurSanglant()
		SeigneurSanglantPattern(player, &adversary)
	default:
		fmt.Println("Niveau de donjon invalide.")
	}
}
