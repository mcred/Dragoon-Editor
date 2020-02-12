package app

import (
	"LODeditor/internal/app/characters"
	"LODeditor/internal/app/inventory"
	"LODeditor/internal/app/storage"
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/sqweek/dialog"
)

func makeForm(slot *storage.Slot, card *storage.Card) fyne.Widget {
	dart := characters.Dart()

	name := widget.NewEntry()
	name.SetPlaceHolder("Dart")
	weapon := widget.NewSelect(inventory.Weapons().GetVals(), func(string) {})
	weapon.SetSelected(inventory.Weapons().GetValByID(slot.GetValueByAttribute(dart.Weapon)))

	form := &widget.Form{
		OnCancel: func() {
			fmt.Println("Cancelled")
		},
		OnSubmit: func() {
			fmt.Println("Form submitted")
			slot.SetValueAtLocation(dart.Weapon, inventory.Weapons().GetIDByVal(weapon.Selected))
			card.SaveCard()
		},
	}
	form.Append("Name", name)
	form.Append("Weapon", weapon)
	return form
}

func Run() {
	var path string
	var err error
	var card storage.Card
	var slot storage.Slot

	if path == "" {
		path, err = dialog.File().Filter("Mednafen Saves", "mcr").Load()
		if err != nil {
			panic(err)
		}
		card = storage.LoadCard(path)
		slot = card.Slots[1]
	}

	a := app.New()
	w := a.NewWindow("Legend of Dragoon Editor")

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.FolderOpenIcon(), func() {
			path, err = dialog.File().Filter("Mednafen Saves", "mcr").Load()
			if err != nil {
				panic(err)
			}
			card = storage.LoadCard(path)
			slot = card.Slots[1]
		}),
		widget.NewToolbarAction(theme.CancelIcon(), a.Quit),
	)

	w.SetContent(widget.NewVBox(
		toolbar,
		makeForm(&slot, &card),
	))

	w.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("Open", func() {
				path, err = dialog.File().Filter("Mednafen Saves", "mcr").Load()
				if err != nil {
					panic(err)
				}
			}),
			fyne.NewMenuItem("Quit", a.Quit),
		),
	))

	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}