package ui

import (
	"LODeditor/internal/app/characters"
	"LODeditor/internal/app/inventory"
	"LODeditor/internal/app/storage"
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func createSelect(i inventory.Inventory, a characters.Attribute, s *storage.Slot) *widget.Select {
	r := widget.NewSelect(i.GetVals(), func(v string) {
		s.SetValueAtLocation(a, i.GetIDByVal(v))
	})
	r.SetSelected(i.GetValByID(s.GetValueByAttribute(a)))
	return r
}

func createCharacterBox(b *widget.Box, c characters.Character, w inventory.Inventory, s *storage.Slot) {
	b.Append(widget.NewLabel(c.Name))
	b.Append(widget.NewLabel("Weapon"))
	b.Append(createSelect(w, c.Weapon, s))
	b.Append(widget.NewLabel("Armor"))
	b.Append(createSelect(inventory.Armor(), c.Chest, s))
	b.Append(widget.NewLabel("Headgear"))
	b.Append(createSelect(inventory.Helms(), c.Helmet, s))
	b.Append(widget.NewLabel("Boots"))
	b.Append(createSelect(inventory.Boots(), c.Boots, s))
	b.Append(widget.NewLabel("Accessories"))
	b.Append(createSelect(inventory.Accessories(), c.Accessory, s))
}

func CharacterForm(slot *storage.Slot, card *storage.Card, w fyne.Window) *fyne.Container {
	form := &widget.Form{
		OnCancel: func() {
			fmt.Println("Cancelled")
		},
		OnSubmit: func() {
			card.SaveCard()
			dialog.ShowInformation("Information", "Card Saved", w)
		},
	}
	dart := characters.Dart()
	box1 := widget.NewVBox()
	createCharacterBox(box1, dart, inventory.Swords(), slot)

	shana := characters.Shana()
	box2 := widget.NewVBox()
	createCharacterBox(box2, shana, inventory.Bows(), slot)

	lavitz := characters.Lavitz()
	box3 := widget.NewVBox()
	createCharacterBox(box3, lavitz, inventory.Spears(), slot)

	rose := characters.Rose()
	box4 := widget.NewVBox()
	createCharacterBox(box4, rose, inventory.Daggers(), slot)

	box5 := widget.NewVBox()
	box5.Append(form)

	return fyne.NewContainerWithLayout(layout.NewGridLayout(4),
		box1, box2, box3, box4, box5)
}
