package characters

import (
	. "LODeditor/internal/app/storage"
)

func Albert() Character { //TODO MAP
	return Character{
		ID:        05,
		Name:      "Albert",
		XP:        Attribute{0x608,4,true},
		HP:        Attribute{0x610,2,true},
		Weapon:    Attribute{0x61C,1,false},
		Helmet:    Attribute{0x61D,1,false},
		Chest:     Attribute{0x61E,1,false},
		Boots:     Attribute{0x61F,1,false},
		Accessory: Attribute{0x620,1,false},
	}
}
