package characters

func Lavitz() Character {
	return Character{
		XP:			Attribute{0x584,4,true},
		HP: 		Attribute{0x534,2,true}, //TODO map this
		Weapon:		Attribute{0x56C,1,false},
		Helmet:    	Attribute{0x56D,1,false},
		Chest:     	Attribute{0x56E,1,false},
		Boots:     	Attribute{0x56F,1,false},
		Accessory: 	Attribute{0x570,1,false},
	}
}
