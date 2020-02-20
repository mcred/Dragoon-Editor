package characters

import (
	. "LODeditor/internal/app/storage"
)

func Haschel() Character { //TODO MAP
	return Character{
		ID:        04,
		Name:      "Haschel",
		XP:        Attribute{0x5DC,4,true},
		HP:        Attribute{0x5E4,2,true},
		Weapon:    Attribute{0x5F0,1,false},
		Helmet:    Attribute{0x5F1,1,false},
		Chest:     Attribute{0x5F2,1,false},
		Boots:     Attribute{0x5F3,1,false},
		Accessory: Attribute{0x5F4,1,false},
	}
}
