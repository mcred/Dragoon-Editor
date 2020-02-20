package characters

import (
	"LODeditor/internal/app/storage"
)

func Meru() Character { //TODO map
	return Character{
		ID:        06,
		Name:      "Meru",
		XP:        storage.Attribute{0x634,4,true},
		HP:        storage.Attribute{0x63C,2,true},
		Weapon:    storage.Attribute{0x648,1,false},
		Helmet:    storage.Attribute{0x649,1,false},
		Chest:     storage.Attribute{0x64A,1,false},
		Boots:     storage.Attribute{0x64B,1,false},
		Accessory: storage.Attribute{0x64C,1,false},
	}
}
