package characters

func Dart() Character {
	return Character{
		XP: 		Attribute{0x52C,4,true},
		HP: 		Attribute{0x534,2,true},
		Weapon:		Attribute{0x540,1,false},
		Helmet:    	Attribute{0x541,1,false},
		Chest:     	Attribute{0x542,1,false},
		Boots:     	Attribute{0x543,1,false},
		Accessory: 	Attribute{0x544,1,false},
	}
}