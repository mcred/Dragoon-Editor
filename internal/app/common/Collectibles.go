package common

import . "LODeditor/internal/app/storage"

func Gold() Attribute {
	return Attribute{0x294, 4, true}
}
func GoldDisplay() Attribute {
	return Attribute{0x19C, 4, true}
}

func Stardust() Attribute {
	return Attribute{0x29C, 1, false}
}
func StardustDisplay() Attribute {
	return Attribute{0x1A8, 1, false}
}