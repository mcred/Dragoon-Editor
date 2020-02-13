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

	dartWeapon := widget.NewSelect(inventory.Weapons().GetVals(), func(v string) {
		slot.SetValueAtLocation(dart.Weapon, inventory.Weapons().GetIDByVal(v))
	})
	dartWeapon.SetSelected(inventory.Weapons().GetValByID(slot.GetValueByAttribute(dart.Weapon)))
	dartArmor := widget.NewSelect(inventory.Armor().GetVals(), func(v string) {
		slot.SetValueAtLocation(dart.Chest, inventory.Armor().GetIDByVal(v))
	})
	dartArmor.SetSelected(inventory.Armor().GetValByID(slot.GetValueByAttribute(dart.Chest)))
	dartHelmet := widget.NewSelect(inventory.Helms().GetVals(), func(v string) {
		slot.SetValueAtLocation(dart.Helmet, inventory.Helms().GetIDByVal(v))
	})
	dartHelmet.SetSelected(inventory.Helms().GetValByID(slot.GetValueByAttribute(dart.Helmet)))
	dartBoots := widget.NewSelect(inventory.Boots().GetVals(), func(v string) {
		slot.SetValueAtLocation(dart.Boots, inventory.Boots().GetIDByVal(v))
	})
	dartBoots.SetSelected(inventory.Boots().GetValByID(slot.GetValueByAttribute(dart.Boots)))
	dartAcc := widget.NewSelect(inventory.Accessories().GetVals(), func(v string) {
		slot.SetValueAtLocation(dart.Accessory, inventory.Accessories().GetIDByVal(v))
	})
	dartAcc.SetSelected(inventory.Accessories().GetValByID(slot.GetValueByAttribute(dart.Accessory)))

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
	form.Append("Weapon", dartWeapon)
	form.Append("Armor", dartArmor)
	form.Append("Headgear", dartHelmet)
	form.Append("Boots", dartBoots)
	form.Append("Accessories", dartAcc)
	return form
}
