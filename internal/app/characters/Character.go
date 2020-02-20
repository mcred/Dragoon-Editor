package characters

import (
	"LODeditor/internal/app/common"
)

type Character struct {
	ID        int
	Name      string
	XP        common.Attribute
	HP        common.Attribute
	Weapon    common.Attribute
	Helmet    common.Attribute
	Chest     common.Attribute
	Boots     common.Attribute
	Accessory common.Attribute
}

func GetCharacters() []Character {
	return []Character {
		Dart(), Shana(), Lavitz(), Rose(), Haschel(), Albert(), Miranda(), Meru(), Kongol(),
	}
}

func GetCharacterNames() []string {
	var r []string
	for _, c := range GetCharacters() {
		r = append(r, c.Name)
	}
	return r
}

func GetNameByID(i int) string {
	r := ""
	for _, c := range GetCharacters() {
		if i == c.ID {
			r = c.Name
		}
	}
	return r
}

func GetIDByName(n string) int {
	r := 0
	for _, c := range GetCharacters() {
		if n == c.Name {
			r = c.ID
		}
	}
	return r
}