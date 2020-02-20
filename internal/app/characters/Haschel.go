package characters

import (
	"LODeditor/internal/app/common"
)

func Haschel() Character { //TODO MAP
	return Character{
		ID:        04,
		Name:      "Haschel",
		XP:        common.Attribute{0x5DC,4,true},
		HP:        common.Attribute{0x5E4,2,true},
		Weapon:    common.Attribute{0x5F0,1,false},
		Helmet:    common.Attribute{0x5F1,1,false},
		Chest:     common.Attribute{0x5F2,1,false},
		Boots:     common.Attribute{0x5F3,1,false},
		Accessory: common.Attribute{0x5F4,1,false},
	}
}
