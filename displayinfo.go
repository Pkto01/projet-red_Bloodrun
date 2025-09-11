package main

import "fmt"

func displayInfo() {
	Name, Class, Level, Pvmax, Pv := initCharacter()
	fmt.Println("Name : ", Name)
	fmt.Println("Class : ", Class)
	fmt.Println("Level : ", Level)
	fmt.Println("Pvmax : ", Pvmax)
	fmt.Println("Pv : ", Pv)

}
