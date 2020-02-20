package common

import "LODeditor/internal/app/storage"

func First() storage.Attribute {
	return storage.Attribute{0x288,1,false}
}
func FirstDisplay() storage.Attribute {
	return storage.Attribute{0x188,1,false}
}

func Second() storage.Attribute {
	return storage.Attribute{0x28C,1,false}
}
func SecondDisplay() storage.Attribute {
	return storage.Attribute{0x18C,1,false}
}

func Third() storage.Attribute {
	return storage.Attribute{0x290,1,false}
}
func ThirdDisplay() storage.Attribute {
	return storage.Attribute{0x190,1,false}
}