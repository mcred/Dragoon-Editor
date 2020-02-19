package ui

import (
	"LODeditor/internal/app/characters"
	"LODeditor/internal/app/common"
	"LODeditor/internal/app/inventory"
	"LODeditor/internal/app/storage"
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

var dart = characters.Dart()
var shana = characters.Shana()
var lavitz = characters.Lavitz()
var rose = characters.Rose()
var haschel = characters.Haschel()
var albert = characters.Albert()
var meru = characters.Meru()
var kongol = characters.Kongol()
var miranda = characters.Miranda()

var pos1 = common.Attribute{0x288,1, false}
var pos1Display = common.Attribute{0x188,1, false}
var pos2 = common.Attribute{0x28C,1, false}
var pos2Display = common.Attribute{0x18C,1, false}
var pos3 = common.Attribute{0x290,1, false}
var pos3Display = common.Attribute{0x190,1, false}

func createCharSelect(i inventory.Inventory, a common.Attribute, s *storage.Slot) *widget.Select {
	r := widget.NewSelect(i.GetVals(), func(v string) {
		s.SetValueAtLocation(a, i.GetIDByVal(v))
	})
	r.SetSelected(i.GetValByID(s.GetValueByAttribute(a)))
	return r
}

func createPosSelect(n []string, a common.Attribute, ad common.Attribute, s *storage.Slot) *widget.Select {
	r := widget.NewSelect(n, func(v string) {
		s.SetValueAtLocation(a, characters.GetIDByName(v))
		s.SetValueAtLocation(ad, characters.GetIDByName(v))
	})
	r.SetSelected(characters.GetNameByID(s.GetValueByAttribute(a)))
	return r
}

func createPartyForm(s *storage.Slot) *fyne.Container {
	names := []string{
		dart.Name, shana.Name, lavitz.Name, rose.Name, haschel.Name,
		albert.Name, meru.Name, kongol.Name, miranda.Name,
	}
	s1 := createPosSelect(names, pos1, pos1Display, s)
	s2 := createPosSelect(names, pos2, pos2Display, s)
	s3 := createPosSelect(names, pos3, pos3Display, s)
	return fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		widget.NewLabel("Party"), s1, s2, s3)
}

func createCharacterBox(b *widget.Box, c characters.Character, w inventory.Inventory, s *storage.Slot) {
	b.Append(widget.NewLabel(c.Name))
	b.Append(widget.NewLabel("Weapon"))
	b.Append(createCharSelect(w, c.Weapon, s))
	b.Append(widget.NewLabel("Armor"))
	b.Append(createCharSelect(inventory.Armor(), c.Chest, s))
	b.Append(widget.NewLabel("Headgear"))
	b.Append(createCharSelect(inventory.Helms(), c.Helmet, s))
	b.Append(widget.NewLabel("Boots"))
	b.Append(createCharSelect(inventory.Boots(), c.Boots, s))
	b.Append(widget.NewLabel("Accessories"))
	b.Append(createCharSelect(inventory.Accessories(), c.Accessory, s))
}

func CreateForm(slot *storage.Slot, card *storage.Card, w fyne.Window) *fyne.Container {
	form := &widget.Form{
		OnCancel: func() {
			fmt.Println("Cancelled")
		},
		OnSubmit: func() {
			card.SaveCard()
			dialog.ShowInformation("Information", "Card Saved", w)
		},
	}
	box1 := widget.NewVBox()
	createCharacterBox(box1, dart, inventory.Swords(), slot)
	box2 := widget.NewVBox()
	createCharacterBox(box2, shana, inventory.Bows(), slot)
	box3 := widget.NewVBox()
	createCharacterBox(box3, lavitz, inventory.Spears(), slot)
	box4 := widget.NewVBox()
	createCharacterBox(box4, rose, inventory.Daggers(), slot)

	chars := fyne.NewContainerWithLayout(layout.NewGridLayout(4),
		box1, box2, box3, box4)
	submit := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		form)

	return fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
		createPartyForm(slot), chars, submit)
}
