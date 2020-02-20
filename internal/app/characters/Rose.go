package characters

import (
	"LODeditor/internal/app/storage"
)

func Rose() Character {
	return Character{
		ID:        3,
		Name:      "Rose",
		XP:        storage.Attribute{0x5B0,4,true},
		HP:        storage.Attribute{0x5B8,2,true},
		Weapon:    storage.Attribute{0x5C4,1,false},
		Helmet:    storage.Attribute{0x5C5,1,false},
		Chest:     storage.Attribute{0x5C6,1,false},
		Boots:     storage.Attribute{0x5C7,1,false},
		Accessory: storage.Attribute{0x5C8,1,false},
	}
}
