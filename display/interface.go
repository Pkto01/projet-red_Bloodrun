package display

import (
	"github.com/common-nighthawk/go-figure"
)

func AsciText() {
	myFigure := figure.NewColorFigure("Bloodrun", "", "red", true)
	myFigure.Print()
}
