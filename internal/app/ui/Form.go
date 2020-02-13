package ui

import (
	"LODeditor/internal/app/characters"
	"LODeditor/internal/app/inventory"
	"LODeditor/internal/app/storage"
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)

func createSelect(i inventory.Inventory, a characters.Attribute, s *storage.Slot) *widget.Select {
	r := widget.NewSelect(i.GetVals(), func(v string) {
		s.SetValueAtLocation(a, i.GetIDByVal(v))
	})
	r.SetSelected(i.GetValByID(s.GetValueByAttribute(a)))
	return r
}

func CharacterForm(slot *storage.Slot, card *storage.Card, w fyne.Window) fyne.Widget {
	dart := characters.Dart()

	name := widget.NewEntry()
	name.SetPlaceHolder("Dart")

	form := &widget.Form{
		OnCancel: func() {
			fmt.Println("Cancelled")
		},
		OnSubmit: func() {
			card.SaveCard()
			dialog.ShowInformation("Information", "Card Saved", w)
		},
	}
	form.Append("Name", name)
	form.Append("Weapon", createSelect(inventory.Weapons(), dart.Weapon, slot))
	form.Append("Armor", createSelect(inventory.Armor(), dart.Chest, slot))
	form.Append("Headgear", createSelect(inventory.Helms(), dart.Helmet, slot))
	form.Append("Boots", createSelect(inventory.Boots(), dart.Boots, slot))
	form.Append("Accessories", createSelect(inventory.Accessories(), dart.Accessory, slot))
	return form
}
