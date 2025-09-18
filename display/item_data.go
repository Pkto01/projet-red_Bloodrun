package display

// ItemStats contient les bonus qu'un objet peut donner.
type ItemStats struct {
	Damage  int
	Defense int
	Ability string // Une chaîne pour identifier une compétence spéciale
}

// ItemTypes définit le slot pour chaque objet équipable.
var ItemTypes = map[string]string{
	// Armes
	"Hache de Berserker": "Weapon",
	"Bâton d'Apprenti":   "Weapon",
	"Marteau Lourd":      "Weapon",

	// Armures
	"Robe de Mage":      "Armor",
	"Armure de Plaques": "Armor",

	// Accessoires (les bottes sont plus logiques ici)
	"Bottes de Célérité":  "Accessory",
	"Bouclier en Acier":   "Accessory",
	"Grimoire des Ombres": "Accessory",
	"Gantelets de Force":  "Accessory",
}

// ItemStatsDatabase associe un nom d'objet à ses statistiques.
var ItemStatsDatabase = map[string]ItemStats{
	// Armes
	"Hache de Berserker": {Damage: 15, Defense: 0, Ability: ""},
	"Bâton d'Apprenti":   {Damage: 8, Defense: 0, Ability: ""},
	"Marteau Lourd":      {Damage: 12, Defense: 0, Ability: ""},

	// Armures
	"Robe de Mage":      {Damage: 0, Defense: 5, Ability: ""},
	"Armure de Plaques": {Damage: 0, Defense: 12, Ability: ""},

	// Accessoires
	"Bottes de Célérité":  {Damage: 0, Defense: 1, Ability: "Initiative"}, // Donne peu de défense mais une compétence clé
	"Bouclier en Acier":   {Damage: 0, Defense: 8, Ability: ""},
	"Grimoire des Ombres": {Damage: 5, Defense: 0, Ability: "MagieNoire"}, // Un grimoire peut aussi augmenter un peu les dégâts
	"Gantelets de Force":  {Damage: 3, Defense: 2, Ability: ""},
}
