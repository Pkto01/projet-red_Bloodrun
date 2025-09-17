package character

type Equipment struct {
	Weapon    string
	Armor     string
	Accessory string
}

type Character struct {
<<<<<<< HEAD
	Name              string
	Class             string
	Level             int
	Money             int
	Pvmax             int
	Pv                int
	skills            []string
	Inventory         []string
	InventoryUpgrades int
=======
	Name      string
	Class     string
	Level     int
	Money     int
	Pvmax     int
	Pv        int
	skills    []string
	Inventory []string
	Equipped  Equipment
>>>>>>> aa0ca1b3ee5889f19c6466368e0741e7ded837a6
}
