package display

import (
	"fmt"
	"projet-red_Bloodrun/character"
)

func DisplayInfo(c character.Character) {
	fmt.Println("Name : ", c.Name)
	fmt.Println("Class : ", c.Class)
	fmt.Println("Level : ", c.Level)
	fmt.Println("Pvmax : ", c.Pvmax)
	fmt.Println("Pv : ", c.Pv)
}
