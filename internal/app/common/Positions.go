package common

import . "LODeditor/internal/app/storage"

func First() Attribute {
	return Attribute{0x288,1,false}
}
func FirstDisplay() Attribute {
	return Attribute{0x188,1,false}
}

func Second() Attribute {
	return Attribute{0x28C,1,false}
}
func SecondDisplay() Attribute {
	return Attribute{0x18C,1,false}
}

func Third() Attribute {
	return Attribute{0x290,1,false}
}
func ThirdDisplay() Attribute {
	return Attribute{0x190,1,false}
}