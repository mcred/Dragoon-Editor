package common

import . "LODeditor/internal/app/storage"

func First() Attribute {
	return Attribute{0x288,4,true}
}
func FirstDisplay() Attribute {
	return Attribute{0x188,4,true}
}

func Second() Attribute {
	return Attribute{0x28C,4,true}
}
func SecondDisplay() Attribute {
	return Attribute{0x18C,4,true}
}

func Third() Attribute {
	return Attribute{0x290,4,true}
}
func ThirdDisplay() Attribute {
	return Attribute{0x190,4,true}
}