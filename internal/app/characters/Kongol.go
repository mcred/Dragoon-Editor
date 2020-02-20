package characters

import (
	"LODeditor/internal/app/common"
)

func Kongol() Character { //TODO map
	return Character{
		ID:        07,
		Name:      "Kongol",
		XP:        common.Attribute{0x660,4,true},
		HP:        common.Attribute{0x668,2,true},
		Weapon:    common.Attribute{0x674,1,false},
		Helmet:    common.Attribute{0x675,1,false},
		Chest:     common.Attribute{0x676,1,false},
		Boots:     common.Attribute{0x677,1,false},
		Accessory: common.Attribute{0x678,1,false},
	}
}
