package character

// InitCharacter initialise un nouveau personnage avec toutes ses statistiques de base.
func InitCharacter(
	Name string,
	Class string,
	Level int,
	Money int,
	Pvmax int,
	Pv int,
	Attack int,
	Defense int,
	Experience int,
	NextLevelExp int,
	skills []string,
	Inventory []string,
	Initiative int,
	Mana int,
	Manamax int) Character {

	return Character{
		Name:              Name,
		Class:             Class,
		Level:             Level,
		Money:             Money,
		Pvmax:             Pvmax,
		Pv:                Pv,
		Attack:            Attack,
		Defense:           Defense,
		Experience:        Experience,
		NextLevelExp:      NextLevelExp,
		skills:            skills,
		Inventory:         Inventory,
		InventoryUpgrades: 0,
		Initiative:        Initiative,
		Mana:              Mana,
		Manamax:           Manamax,
		Equipped: Equipment{
			Weapon:    "Aucune",
			Armor:     "Aucune",
			Accessory: "Aucune",
		},
		Spells: []Spell{
			{Name: "Coup de poing", Damage: 8, Heal: 0, Mana: 0},
			{Name: "Boule de feu", Damage: 18, Heal: 0, Mana: 10},
			{Name: "Soin", Damage: 0, Heal: 15, Mana: 8},
		},
	}
}
