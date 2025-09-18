package fight

// Monster définit la structure d'un adversaire.
type Monster struct {
	Name    string
	Pvmax   int
	Pv      int
	Attack  int
	Defense int
}

// InitGoblin crée et retourne un Gobelin d'entraînement avec des stats prédéfinies.
func InitGoblin() Monster {
	return Monster{
		Name:    "Gobelin d'entraînement",
		Pvmax:   60,
		Pv:      60,
		Attack:  5,
		Defense: 2,
	}
}
