package characters

import (
	"LODeditor/internal/app/common"
)

func Miranda() Character { //TODO Map
	return Character{
		ID:        8,
		Name:      "Miranda",
		XP:        common.Attribute{0x5B0,4,true},
		HP:        common.Attribute{0x5B8,2,true},
		Weapon:    common.Attribute{0x5C4,1,false},
		Helmet:    common.Attribute{0x5C5,1,false},
		Chest:     common.Attribute{0x5C6,1,false},
		Boots:     common.Attribute{0x5C7,1,false},
		Accessory: common.Attribute{0x5C8,1,false},
	}
}
