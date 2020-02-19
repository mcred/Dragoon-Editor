package characters

import (
	"LODeditor/internal/app/common"
)

func Dart() Character {
	return Character{
		ID:        0,
		Name:      "Dart",
		XP:        common.Attribute{0x52C,4,true},
		HP:        common.Attribute{0x534,2,true},
		Weapon:    common.Attribute{0x540,1,false},
		Helmet:    common.Attribute{0x541,1,false},
		Chest:     common.Attribute{0x542,1,false},
		Boots:     common.Attribute{0x543,1,false},
		Accessory: common.Attribute{0x544,1,false},
	}
}