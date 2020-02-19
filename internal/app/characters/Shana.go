package characters

import (
	"LODeditor/internal/app/common"
)

func Shana() Character {
	return Character{
		ID:			2,
		Name:		"Shana",
		XP:			common.Attribute{0x558,4,true},
		HP: 		common.Attribute{0x560,2,true},
		Weapon:		common.Attribute{0x598,1,false},
		Helmet:    	common.Attribute{0x599,1,false},
		Chest:     	common.Attribute{0x59A,1,false},
		Boots:     	common.Attribute{0x59B,1,false},
		Accessory: 	common.Attribute{0x59C,1,false},
	}
}
