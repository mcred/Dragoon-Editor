package characters

type Character struct {
	HP Attribute
	Weapon Attribute
	Helmet Attribute
	Chest Attribute
	Boots Attribute
	Accessory Attribute
}

type Attribute struct {
	Location int
	Bytes int
	Reversed bool
}

func Dart() Character {
	return Character{
		HP: 		Attribute{0x534, 2,true},
		Weapon:		Attribute{0x540, 1,false},
		Helmet:    	Attribute{0x5411,1,false},
		Chest:     	Attribute{0x5421,1,false},
		Boots:     	Attribute{0x5431,1,false},
		Accessory: 	Attribute{0x5441,1,false},
	}
}



