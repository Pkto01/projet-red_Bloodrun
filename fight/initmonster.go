package fight

func InitMonster(MName string, MPv int, MPvmax int, MAttack int) Monster {
	return Monster{
		MName:   MName,
		MPvmax:  MPvmax,
		MPv:     MPv,
		MAttack: MAttack,
	}
}
