package characters

import (
	"LODeditor/internal/app/storage"
)

func Miranda() Character { //TODO Map
	return Character{
		ID:        8,
		Name:      "Miranda",
		XP:        storage.Attribute{0x68C,4,true},
		HP:        storage.Attribute{0x694,2,true},
		Weapon:    storage.Attribute{0x6A0,1,false},
		Helmet:    storage.Attribute{0x6A1,1,false},
		Chest:     storage.Attribute{0x6A2,1,false},
		Boots:     storage.Attribute{0x6A3,1,false},
		Accessory: storage.Attribute{0x6A4,1,false},
	}
}
