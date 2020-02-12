package storage

import (
	"fmt"
	"io/ioutil"
)

type Card struct {
	Path string
	Slots []Slot
}

func LoadCard(path string) Card {
 	f, err := ioutil.ReadFile(path)
 	var s []Slot
 	if err == nil {
		for i := 0;  i<=15; i++ {
			b := i * 0x2000
			t := b + 0x1FFF
			s = append(s, Slot{Head:b, Data:f[b:t]})
		}
	}
	return Card{
		Path:  path,
		Slots: s,
	}
}

func (c Card) SaveCard() {
	var b []byte
	for _, slot := range c.Slots {
		b = append(b, slot.Data...)
	}
	fmt.Println(c.Slots[1].Data[0x540])
	fmt.Println(b[0x2540]) //TODO this is shifted one byte too early


	//err := ioutil.WriteFile(c.Path, b, 0644)
	//if err != nil {
	//	panic(err)
	//}
}