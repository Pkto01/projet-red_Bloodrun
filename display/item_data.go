package display

// ItemStats contient les bonus qu'un objet peut donner.
type ItemStats struct {
	Damage  int
	Defense int
	Initiative int
	Slot       string
}

// ItemStatsDatabase associe un nom d'objet à ses statistiques.
var ItemStatsDatabase = map[string]ItemStats{
	// Armes
	"Hache de Berserker": {Damage: 20, Defense: 0, Initiative: 0, Slot: "Weapon"},
	"Bâton d'Apprenti":   {Damage: 14, Defense: 0, Initiative: 12, Slot: "Weapon"},
	"Marteau Lourd":      {Damage: 16, Defense: 0, Initiative: -5, Slot: "Weapon"},

	// Armures
	"Robe de Mage":      {Damage: 0, Defense: 22, Initiative: 12, Slot: "Armor"},
	"Armure de Plaques": {Damage: 0, Defense: 30, Initiative: -5, Slot: "Armor"},
	"Bottes de Célérité":  {Damage: 0, Defense: 10, Initiative: 20, Slot: "Armor"},

	// Accessoires
	"Bouclier en Acier":   {Damage: 0, Defense: 20, Initiative: -5, Slot: "Accessory"},
	"Grimoire des Ombres": {Damage: 20, Defense: 0, Initiative: 12, Slot: "Accessory"},
	"Gantelets de Force":  {Damage: 30, Defense: 15, Initiative: 5, Slot: "Accessory"},
}
