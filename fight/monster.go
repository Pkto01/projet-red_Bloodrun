package fight

// Monster définit la structure d'un adversaire avec ses récompenses.
type Monster struct {
	Name             string
	Pvmax            int
	Pv               int
	Attack           int
	Defense          int
	ExperienceReward int
	GoldReward       int    // Ajout de la récompense en Or
	LootDrop         string // Ajout du butin spécifique
	Initiative       int
}

// --- Monstre d'entraînement ---

// InitGoblin crée un Gobelin d'entraînement.
func InitGoblin() Monster {
	return Monster{
		Name:             "Gobelin d'entraînement",
		Pvmax:            60,
		Pv:               60,
		Attack:           10,
		Defense:          2,
		ExperienceReward: 5,
		GoldReward:       10,
		LootDrop:         "Os Fêlés", // Les gobelins peuvent aussi laisser des os
	}
}

// --- Monstres des Donjons ---

// Monstre 1: Couloir Saignant
func InitGouleSanguine() Monster {
	return Monster{
		Name:             "Goule Sanguine",
		Pvmax:            80,
		Pv:               80,
		Attack:           11,
		Defense:          3,
		ExperienceReward: 50,
		GoldReward:       30,
		LootDrop:         "Fragments de Sang",
		Initiative:       15,
	}
}

// Monstre 2: Les Fosses de Chair
func InitAbominationDeChair() Monster {
	return Monster{
		Name:             "Abomination de Chair",
		Pvmax:            150,
		Pv:               150,
		Attack:           48,
		Defense:          5,
		ExperienceReward: 150,
		GoldReward:       50,
		LootDrop:         "Étoffe Sanglante",
		Initiative:       20,
	}
}

// Monstre 3: Sanctuaire des Os Brisés
func InitGardienSquelette() Monster {
	return Monster{
		Name:             "Gardien Squelette",
		Pvmax:            120,
		Pv:               120,
		Attack:           83,
		Defense:          10,
		ExperienceReward: 500,
		GoldReward:       100,
		LootDrop:         "Os Fêlés",
		Initiative:       30,
	}
}

// Monstre 4: La Forge Écarlate
func InitGolemDeLaForge() Monster {
	return Monster{
		Name:             "Golem de la Forge",
		Pvmax:            200,
		Pv:               200,
		Attack:           140,
		Defense:          15,
		ExperienceReward: 1400,
		GoldReward:       600,
		LootDrop:         "Acier Noirci",
		Initiative:       45,
	}
}

// Monstre 5: Trône du Seigneur Sanglant (BOSS)
func InitSeigneurSanglant() Monster {
	return Monster{
		Name:             "Seigneur Sanglant",
		Pvmax:            300,
		Pv:               300,
		Attack:           200,
		Defense:          12,
		ExperienceReward: 5000000,
		GoldReward:       25000000000000000,  // Grosse récompense en or
		LootDrop:         "Coeur Démoniaque", // Un butin unique pour le boss
		Initiative:       70,
	}
}
