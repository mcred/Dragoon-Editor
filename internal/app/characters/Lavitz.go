package characters

import (
	"LODeditor/internal/app/common"
)

func Lavitz() Character {
	return Character{
		ID:        1,
		Name:      "Lavitz",
		XP:        common.Attribute{0x584,4,true},
		HP:        common.Attribute{0x534,2,true}, //TODO map this
		Weapon:    common.Attribute{0x56C,1,false},
		Helmet:    common.Attribute{0x56D,1,false},
		Chest:     common.Attribute{0x56E,1,false},
		Boots:     common.Attribute{0x56F,1,false},
		Accessory: common.Attribute{0x570,1,false},
	}
}
