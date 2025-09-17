package character

type Character struct {
	Name              string
	Class             string
	Level             int
	Money             int
	Pvmax             int
	Pv                int
	skills            []string
	Inventory         []string
	InventoryUpgrades int
}
