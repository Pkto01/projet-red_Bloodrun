package character

// InitCharacter initialise un nouveau personnage avec toutes ses statistiques de base.
func InitCharacter(
	Name string,
	Class string,
	Level int,
	Money int,
	Pvmax int,
	Pv int,
	Manamax int,
	Mana int,
	Attack int,
	Defense int,
	Experience int,
	NextLevelExp int,
	Spells []Spell,
	Inventory []string,
	Initiative int) Character {

	return Character{
		Name:              Name,
		Class:             Class,
		Level:             Level,
		Money:             Money,
		Pvmax:             Pvmax,
		Pv:                Pv,
		Manamax:           Manamax,
		Mana:              Mana,
		Attack:            Attack,
		Defense:           Defense,
		Experience:        Experience,
		NextLevelExp:      NextLevelExp,
		Spells:            Spells,
		Inventory:         Inventory,
		InventoryUpgrades: 0,
		Initiative:        Initiative,
		Equipped: Equipment{
			Weapon:    "Aucune",
			Armor:     "Aucune",
			Accessory: "Aucune",
		},
	}
}
