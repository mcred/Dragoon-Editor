package ui

import (
	"LODeditor/internal/app/storage"
	"fyne.io/fyne"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/sqweek/dialog"
)

func Toolbar(slot *storage.Slot, a fyne.App) fyne.Widget {
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.FolderOpenIcon(), func() {
			path, err := dialog.File().Filter("Mednafen Saves", "mcr").Load()
			if err != nil {
				panic(err)
			}
			card := storage.LoadCard(path)
			*slot = card.Slots[1]
		}),
		widget.NewToolbarAction(theme.CancelIcon(), a.Quit),
	)
	return toolbar
}