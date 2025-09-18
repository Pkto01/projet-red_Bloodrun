package character

type Equipment struct {
	Weapon    string
	Armor     string
	Accessory string
}

type Character struct {
	Name              string
	Class             string
	Level             int
	Money             int
	Pvmax             int
	Pv                int
	Attack            int
	Defense           int
	Experience        int
	NextLevelExp      int
	skills            []string
	Inventory         []string
	InventoryUpgrades int
	Equipped          Equipment
	Initiative        int
}
