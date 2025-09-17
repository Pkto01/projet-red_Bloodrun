package character

type Equipment struct {
	Weapon    string
	Armor     string
	Accessory string
}

type Character struct {
	Name      string
	Class     string
	Level     int
	Money     int
	Pvmax     int
	Pv        int
	skills    []string
	Inventory []string
	Equipped  Equipment
}
