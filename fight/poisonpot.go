package fight

import (
	"fmt"
	"projet-red_Bloodrun/character"
	"time"
)

func poisonPot(m *character.Character) {

	fmt.Println("La potion de poison est activée!")
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		m.Pv -= 10
		if m.Pv < 0 {
			m.Pv = 0
		}
		fmt.Printf("Dégâts infligés! Points de vie actuels: %d/%d\n", m.Pv, m.Pvmax)
	}
}
