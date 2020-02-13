package characters

func Lavitz() Character {
	return Character{
		XP:			Attribute{0x558,4,true},
		HP: 		Attribute{0x560,2,true},
		Weapon:		Attribute{0x598,1,false},
		Helmet:    	Attribute{0x599,1,false},
		Chest:     	Attribute{0x59A,1,false},
		Boots:     	Attribute{0x59B,1,false},
		Accessory: 	Attribute{0x59C,1,false},
	}
}
