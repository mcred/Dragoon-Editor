package characters

import (
	. "LODeditor/internal/app/storage"
)

func Rose() Character {
	return Character{
		ID:        	3,
		Name:      	"Rose",
		XP:        	Attribute{0x5B0,4,true},
		HP:        	Attribute{0x5B8,2,true},
		MP:		   	Attribute{0x5BA,2,true},
		Level:		Attribute{0x5C2,1,false},
		Weapon:    	Attribute{0x5C4,1,false},
		Helmet:    	Attribute{0x5C5,1,false},
		Chest:     	Attribute{0x5C6,1,false},
		Boots:     	Attribute{0x5C7,1,false},
		Accessory: 	Attribute{0x5C8,1,false},
	}
}
