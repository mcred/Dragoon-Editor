package characters

import (
	"LODeditor/internal/app/common"
)

func Meru() Character { //TODO map
	return Character{
		ID:        06,
		Name:      "Meru",
		XP:        common.Attribute{0x634,4,true},
		HP:        common.Attribute{0x63C,2,true},
		Weapon:    common.Attribute{0x648,1,false},
		Helmet:    common.Attribute{0x649,1,false},
		Chest:     common.Attribute{0x64A,1,false},
		Boots:     common.Attribute{0x64B,1,false},
		Accessory: common.Attribute{0x64C,1,false},
	}
}
