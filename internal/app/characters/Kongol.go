package characters

import (
	. "LODeditor/internal/app/storage"
)

func Kongol() Character { //TODO map
	return Character{
		ID:        	07,
		Name:      	"Kongol",
		XP:        	Attribute{0x660,4,true},
		HP:        	Attribute{0x668,2,true},
		MP:		   	Attribute{0x66A,2,true},
		Level:		Attribute{0x672,1,false},
		Weapon:    	Attribute{0x674,1,false},
		Helmet:    	Attribute{0x675,1,false},
		Chest:     	Attribute{0x676,1,false},
		Boots:     	Attribute{0x677,1,false},
		Accessory: 	Attribute{0x678,1,false},
	}
}
