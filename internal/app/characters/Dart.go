package characters

import (
	. "LODeditor/internal/app/storage"
)

func Dart() Character {
	return Character{
		ID:        	0,
		Name:      	"Dart",
		XP:        	Attribute{0x52C,4,true},
		HP:        	Attribute{0x534,2,true},
		MP:			Attribute{0x536,2,true},
		Level:		Attribute{0x5C2,1,false},
		Weapon:    	Attribute{0x540,1,false},
		Helmet:    	Attribute{0x541,1,false},
		Chest:     	Attribute{0x542,1,false},
		Boots:     	Attribute{0x543,1,false},
		Accessory: 	Attribute{0x544,1,false},
	}
}