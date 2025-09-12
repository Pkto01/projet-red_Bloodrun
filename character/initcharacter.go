package character

func InitCharacter(Name string, Class string, Level int, Money int, Pvmax int, Pv int, Inventory []string) Character {
	return Character{
		Name:      Name,
		Class:     Class,
		Level:     Level,
		Money:     Money,
		Pvmax:     Pvmax,
		Pv:        Pv,
		Inventory: Inventory,
	}
}
