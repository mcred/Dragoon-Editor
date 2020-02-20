package characters

import (
	"LODeditor/internal/app/storage"
)

func Shana() Character {
	return Character{
		ID:        2,
		Name:      "Shana",
		XP:        storage.Attribute{0x584,4,true},
		HP:        storage.Attribute{0x58C,2,true},
		Weapon:    storage.Attribute{0x598,1,false},
		Helmet:    storage.Attribute{0x599,1,false},
		Chest:     storage.Attribute{0x59A,1,false},
		Boots:     storage.Attribute{0x59B,1,false},
		Accessory: storage.Attribute{0x59C,1,false},
	}
}
