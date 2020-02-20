package characters

import (
	. "LODeditor/internal/app/storage"
)

func Miranda() Character { //TODO Map
	return Character{
		ID:        8,
		Name:      "Miranda",
		XP:        Attribute{0x68C,4,true},
		HP:        Attribute{0x694,2,true},
		Weapon:    Attribute{0x6A0,1,false},
		Helmet:    Attribute{0x6A1,1,false},
		Chest:     Attribute{0x6A2,1,false},
		Boots:     Attribute{0x6A3,1,false},
		Accessory: Attribute{0x6A4,1,false},
	}
}
