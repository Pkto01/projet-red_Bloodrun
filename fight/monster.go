package fight

// Monster définit la structure d'un adversaire.
type Monster struct {
	Name             string
	Pvmax            int
	Pv               int
	Attack           int
	Defense          int
	ExperienceReward int // Ajout de la récompense en expérience
}

// --- Monstre d'entraînement ---

// InitGoblin crée et retourne un Gobelin d'entraînement avec des stats prédéfinies.
func InitGoblin() Monster {
	return Monster{
		Name:             "Gobelin d'entraînement",
		Pvmax:            60,
		Pv:               60,
		Attack:           5,
		Defense:          2,
		ExperienceReward: 5, // Récompense de base pour le test
	}
}

// --- Monstres des Donjons ---

// Monstre 1: Couloir Saignant
func InitGouleSanguine() Monster {
	return Monster{
		Name:             "Goule Sanguine",
		Pvmax:            80,
		Pv:               80,
		Attack:           8,
		Defense:          3,
		ExperienceReward: 10,
	}
}

// Monstre 2: Les Fosses de Chair
func InitAbominationDeChair() Monster {
	return Monster{
		Name:             "Abomination de Chair",
		Pvmax:            150,
		Pv:               150,
		Attack:           12,
		Defense:          5,
		ExperienceReward: 25,
	}
}

// Monstre 3: Sanctuaire des Os Brisés
func InitGardienSquelette() Monster {
	return Monster{
		Name:             "Gardien Squelette",
		Pvmax:            120,
		Pv:               120,
		Attack:           10,
		Defense:          10,
		ExperienceReward: 40,
	}
}

// Monstre 4: La Forge Écarlate
func InitGolemDeLaForge() Monster {
	return Monster{
		Name:             "Golem de la Forge",
		Pvmax:            200,
		Pv:               200,
		Attack:           15,
		Defense:          15,
		ExperienceReward: 75,
	}
}

// Monstre 5: Trône du Seigneur Sanglant (BOSS)
func InitSeigneurSanglant() Monster {
	return Monster{
		Name:             "Seigneur Sanglant",
		Pvmax:            300,
		Pv:               300,
		Attack:           18,
		Defense:          12,
		ExperienceReward: 150,
	}
}
