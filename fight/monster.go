package fight

import (
	"fmt"
)

type Monster struct {
	MName   string
	MPvmax  int
	MPv     int
	MAttack int
}

func (m *Monster) IsAlive() bool {
	return m.MPv > 0
}

func (m *Monster) TakeDamage(damage int) {
	m.MPv -= damage
	if m.MPv < 0 {
		m.MPv = 0
	}
}

func (m *Monster) String() string {
	return fmt.Sprintf("%s | PV: %d/%d | ATK: %d",
		m.MName, m.MPv, m.MPvmax, m.MAttack)
}
