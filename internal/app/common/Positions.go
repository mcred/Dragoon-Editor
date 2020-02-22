package common

import . "LODeditor/internal/app/storage"

// First : Location of first character slot in game
func First() Attribute {
	return Attribute{0x288, 4, true}
}

// FirstDisplay : Location of first character slot for the save file
func FirstDisplay() Attribute {
	return Attribute{0x188, 4, true}
}

// Second : Location of second character slot in game
func Second() Attribute {
	return Attribute{0x28C, 4, true}
}

// SecondDisplay : Location of second character slot for the save file
func SecondDisplay() Attribute {
	return Attribute{0x18C, 4, true}
}

// Third : Location of third character slot in game
func Third() Attribute {
	return Attribute{0x290, 4, true}
}

// ThirdDisplay : Location of third character slot for the save file
func ThirdDisplay() Attribute {
	return Attribute{0x190, 4, true}
}
