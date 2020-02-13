package ui

import (
	"LODeditor/internal/app/characters"
	"LODeditor/internal/app/inventory"
	"LODeditor/internal/app/storage"
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func CharacterForm(slot *storage.Slot, card *storage.Card) fyne.Widget {
	dart := characters.Dart()

	name := widget.NewEntry()
	name.SetPlaceHolder("Dart")
	weapon := widget.NewSelect(inventory.Weapons().GetVals(), func(v string) {
		slot.SetValueAtLocation(dart.Weapon, inventory.Weapons().GetIDByVal(v))
	})
	weapon.SetSelected(inventory.Weapons().GetValByID(slot.GetValueByAttribute(dart.Weapon)))

	form := &widget.Form{
		OnCancel: func() {
			fmt.Println("Cancelled")
		},
		OnSubmit: func() {
			fmt.Println("Form submitted")
			card.SaveCard()
		},
	}
	form.Append("Name", name)
	form.Append("Weapon", weapon)
	return form
}
