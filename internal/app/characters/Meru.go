package characters

import (
	. "LODeditor/internal/app/storage"
)

func Meru() Character { //TODO map
	return Character{
		ID:        06,
		Name:      "Meru",
		XP:        Attribute{0x634,4,true},
		HP:        Attribute{0x63C,2,true},
		Weapon:    Attribute{0x648,1,false},
		Helmet:    Attribute{0x649,1,false},
		Chest:     Attribute{0x64A,1,false},
		Boots:     Attribute{0x64B,1,false},
		Accessory: Attribute{0x64C,1,false},
	}
}
