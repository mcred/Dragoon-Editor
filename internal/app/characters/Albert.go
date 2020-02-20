package characters

import (
	"LODeditor/internal/app/common"
)

func Albert() Character { //TODO MAP
	return Character{
		ID:        05,
		Name:      "Albert",
		XP:        common.Attribute{0x608,4,true},
		HP:        common.Attribute{0x610,2,true},
		Weapon:    common.Attribute{0x61C,1,false},
		Helmet:    common.Attribute{0x61D,1,false},
		Chest:     common.Attribute{0x61E,1,false},
		Boots:     common.Attribute{0x61F,1,false},
		Accessory: common.Attribute{0x620,1,false},
	}
}
