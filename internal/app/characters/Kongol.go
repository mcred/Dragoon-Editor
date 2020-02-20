package characters

import (
	"LODeditor/internal/app/storage"
)

func Kongol() Character { //TODO map
	return Character{
		ID:        07,
		Name:      "Kongol",
		XP:        storage.Attribute{0x660,4,true},
		HP:        storage.Attribute{0x668,2,true},
		Weapon:    storage.Attribute{0x674,1,false},
		Helmet:    storage.Attribute{0x675,1,false},
		Chest:     storage.Attribute{0x676,1,false},
		Boots:     storage.Attribute{0x677,1,false},
		Accessory: storage.Attribute{0x678,1,false},
	}
}
