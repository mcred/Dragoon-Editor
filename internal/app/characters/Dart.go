package characters

import (
	"LODeditor/internal/app/storage"
)

func Dart() Character {
	return Character{
		ID:        0,
		Name:      "Dart",
		XP:        storage.Attribute{0x52C,4,true},
		HP:        storage.Attribute{0x534,2,true},
		Weapon:    storage.Attribute{0x540,1,false},
		Helmet:    storage.Attribute{0x541,1,false},
		Chest:     storage.Attribute{0x542,1,false},
		Boots:     storage.Attribute{0x543,1,false},
		Accessory: storage.Attribute{0x544,1,false},
	}
}