package common

import . "LODeditor/internal/app/storage"

// Gold : Location of in game Gold value
func Gold() Attribute {
	return Attribute{0x294, 4, true}
}

// GoldDisplay : Location of Gold displayed on the save file
func GoldDisplay() Attribute {
	return Attribute{0x19C, 4, true}
}

// Stardust : Location of in game Stardust value
func Stardust() Attribute {
	return Attribute{0x29C, 1, false}
}

// StardustDisplay : Location of Stardust displayed on the save file
func StardustDisplay() Attribute {
	return Attribute{0x1A8, 1, false}
}
