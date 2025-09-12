package display

import "fmt"

func displayInfo(c Character) {
	fmt.Println("Name : ", c.Name)
	fmt.Println("Class : ", c.Class)
	fmt.Println("Level : ", c.Level)
	fmt.Println("Pvmax : ", c.Pvmax)
	fmt.Println("Pv : ", c.Pv)
	fmt.Println("Inventory : ", c.Inventory)
}
