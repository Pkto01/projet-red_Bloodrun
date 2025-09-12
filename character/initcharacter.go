package character

func InitCharacter(Name string, Class string, Level int, Pvmax int, Pv int, Inventory []string) Character {
	return Character{
		Name:      Name,
		Class:     Class,
		Level:     Level,
		Pvmax:     Pvmax,
		Pv:        Pv,
		Inventory: Inventory,
	}
}
