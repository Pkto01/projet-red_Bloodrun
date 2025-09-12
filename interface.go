package main

import (
	"github.com/common-nighthawk/go-figure"
)

func main() {
	myFigure := figure.NewColorFigure("Bloodrun", "", "red", true)
	myFigure.Print()
}
