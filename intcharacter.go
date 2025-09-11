package main

func initCharacter(Name string, Class string, Level int, Pvmax int, Pv int, Inventory []string) Character {
	return Character{
		Name:      "Your name",
		Class:     "Elfe",
		Level:     1,
		Pvmax:     300,
		Pv:        30,
		Inventory: []string{"3 potions"},
	}
}
