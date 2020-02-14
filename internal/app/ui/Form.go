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
	box1.Append(widget.NewLabel("Dart"))
	box1.Append(widget.NewLabel("Weapon"))
	box1.Append(createSelect(inventory.Swords(), dart.Weapon, slot))
	box1.Append(widget.NewLabel("Armor"))
	box1.Append(createSelect(inventory.Armor(), dart.Chest, slot))
	box1.Append(widget.NewLabel("Headgear"))
	box1.Append(createSelect(inventory.Helms(), dart.Helmet, slot))
	box1.Append(widget.NewLabel("Boots"))
	box1.Append(createSelect(inventory.Boots(), dart.Boots, slot))
	box1.Append(widget.NewLabel("Accessories"))
	box1.Append(createSelect(inventory.Accessories(), dart.Accessory, slot))

	shana := characters.Shana()
	box2 := widget.NewVBox()
	box2.Append(widget.NewLabel("Shana"))
	box2.Append(widget.NewLabel("Weapon"))
	box2.Append(createSelect(inventory.Bows(), shana.Weapon, slot))
	box2.Append(widget.NewLabel("Armor"))
	box2.Append(createSelect(inventory.Armor(), shana.Chest, slot))
	box2.Append(widget.NewLabel("Headgear"))
	box2.Append(createSelect(inventory.Helms(), shana.Helmet, slot))
	box2.Append(widget.NewLabel("Boots"))
	box2.Append(createSelect(inventory.Boots(), shana.Boots, slot))
	box2.Append(widget.NewLabel("Accessories"))
	box2.Append(createSelect(inventory.Accessories(), shana.Accessory, slot))

	lavitz := characters.Lavitz()
	box3 := widget.NewVBox()
	box3.Append(widget.NewLabel("Lavitz"))
	box3.Append(widget.NewLabel("Weapon"))
	box3.Append(createSelect(inventory.Spears(), lavitz.Weapon, slot))
	box3.Append(widget.NewLabel("Armor"))
	box3.Append(createSelect(inventory.Armor(), lavitz.Chest, slot))
	box3.Append(widget.NewLabel("Headgear"))
	box3.Append(createSelect(inventory.Helms(), lavitz.Helmet, slot))
	box3.Append(widget.NewLabel("Boots"))
	box3.Append(createSelect(inventory.Boots(), lavitz.Boots, slot))
	box3.Append(widget.NewLabel("Accessories"))
	box3.Append(createSelect(inventory.Accessories(), lavitz.Accessory, slot))

	box4 := widget.NewVBox()
	box4.Append(form)

	return fyne.NewContainerWithLayout(layout.NewGridLayout(3),
		box1, box2, box3, box4)
}
