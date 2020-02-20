package characters

import (
	"LODeditor/internal/app/storage"
)

func Lavitz() Character {
	return Character{
		ID:        1,
		Name:      "Lavitz",
		XP:        storage.Attribute{0x558,4,true},
		HP:        storage.Attribute{0x560,2,true},
		Weapon:    storage.Attribute{0x56C,1,false},
		Helmet:    storage.Attribute{0x56D,1,false},
		Chest:     storage.Attribute{0x56E,1,false},
		Boots:     storage.Attribute{0x56F,1,false},
		Accessory: storage.Attribute{0x570,1,false},
	}
}
