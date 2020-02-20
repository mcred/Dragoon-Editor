package characters

import (
	"LODeditor/internal/app/storage"
)

func Albert() Character { //TODO MAP
	return Character{
		ID:        05,
		Name:      "Albert",
		XP:        storage.Attribute{0x608,4,true},
		HP:        storage.Attribute{0x610,2,true},
		Weapon:    storage.Attribute{0x61C,1,false},
		Helmet:    storage.Attribute{0x61D,1,false},
		Chest:     storage.Attribute{0x61E,1,false},
		Boots:     storage.Attribute{0x61F,1,false},
		Accessory: storage.Attribute{0x620,1,false},
	}
}
