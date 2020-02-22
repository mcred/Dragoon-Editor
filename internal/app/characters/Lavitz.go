package characters

import (
	. "LODeditor/internal/app/storage"
)

func Lavitz() Character {
	return Character{
		ID:        	1,
		Name:      	"Lavitz",
		XP:        	Attribute{0x558,4,true},
		HP:        	Attribute{0x560,2,true},
		MP:		   	Attribute{0x562,2,true},
		Weapon:    	Attribute{0x56C,1,false},
		Helmet:    	Attribute{0x56D,1,false},
		Chest:     	Attribute{0x56E,1,false},
		Boots:     	Attribute{0x56F,1,false},
		Accessory: 	Attribute{0x570,1,false},
	}
}
