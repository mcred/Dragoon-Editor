package characters

import (
	"LODeditor/internal/app/storage"
)

func Haschel() Character { //TODO MAP
	return Character{
		ID:        04,
		Name:      "Haschel",
		XP:        storage.Attribute{0x5DC,4,true},
		HP:        storage.Attribute{0x5E4,2,true},
		Weapon:    storage.Attribute{0x5F0,1,false},
		Helmet:    storage.Attribute{0x5F1,1,false},
		Chest:     storage.Attribute{0x5F2,1,false},
		Boots:     storage.Attribute{0x5F3,1,false},
		Accessory: storage.Attribute{0x5F4,1,false},
	}
}
