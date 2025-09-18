package fight

import (
	"fmt"
	"projet-red_Bloodrun/character"
	"time"
)

// combatLoop est la boucle de base pour tous les combats.
// Elle gère le déroulement des tours et appelle une fonction pour l'action du monstre.
func combatLoop(player *character.Character, adversary *Monster, monsterAction func(turn int, m *Monster, p *character.Character)) {
	// Affiche l'art du monstre avant le début du combat
	DisplayMonsterArt(adversary.Name)

	fmt.Println("\n💥💥💥 LE COMBAT COMMENCE ! 💥💥💥")
	turn := 1

	for player.Pv > 0 && adversary.Pv > 0 {
		fmt.Printf("\n---------- TOUR %d ----------\n", turn)

		// --- Tour du Joueur ---
		// TODO: Implémenter ici un menu pour le joueur (Attaquer, Inventaire, Compétences)
		playerDamage := player.Attack - adversary.Defense
		if playerDamage < 1 {
			playerDamage = 1 // Toujours infliger au moins 1 de dégât
		}
		adversary.Pv -= playerDamage
		fmt.Printf("Vous infligez à %s %d de dégâts.\n", adversary.Name, playerDamage)
		fmt.Printf("Points de vie de l'ennemi : %d/%d PV\n", adversary.Pv, adversary.Pvmax)

		if adversary.Pv <= 0 {
			break
		}
		time.Sleep(1 * time.Second)

		// --- Tour du Monstre (action spécifique définie par la fonction monsterAction) ---
		monsterAction(turn, adversary, player)

		turn++
		time.Sleep(1 * time.Second)
	}

	// --- Fin du Combat ---
	fmt.Println("\n---------- FIN DU COMBAT ----------")
	if player.Pv <= 0 {
		fmt.Println("Vous avez été vaincu... 💀")
	} else {
		fmt.Printf("Vous avez vaincu : %s !\n", adversary.Name)
		// player.GainExperience(adversary.ExperienceReward) // Le joueur gagne de l'expérience
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
