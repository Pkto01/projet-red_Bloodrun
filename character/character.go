package character

type Equipment struct {
	Weapon    string
	Armor     string
	Accessory string
}

type Spell struct {
	Name   string
	Damage int
	Heal   int
	Mana   int // Coût en mana du sort
}

type Character struct {
	Name string
	Class string
	Level int
	Money int
	Pvmax int
	Pv int
	Manamax int
	Mana int
	// Stats totales (celles qui changent avec l'équipement)
	Attack int
	Defense int
	Initiative int
	// Stats de base (celles qui augmentent avec les niveaux)
	BaseAttack int
	BaseDefense int
	BaseInitiative int
	Experience int
	NextLevelExp int
	Spells []Spell
	Inventory []string
	InventoryUpgrades int
	Equipped Equipment
}