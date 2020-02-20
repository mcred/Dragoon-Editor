package characters

import (
	"LODeditor/internal/app/common"
)

func Miranda() Character { //TODO Map
	return Character{
		ID:        8,
		Name:      "Miranda",
		XP:        common.Attribute{0x68C,4,true},
		HP:        common.Attribute{0x694,2,true},
		Weapon:    common.Attribute{0x6A0,1,false},
		Helmet:    common.Attribute{0x6A1,1,false},
		Chest:     common.Attribute{0x6A2,1,false},
		Boots:     common.Attribute{0x6A3,1,false},
		Accessory: common.Attribute{0x6A4,1,false},
	}
}
