package characters

type Character struct {
	XP Attribute
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
