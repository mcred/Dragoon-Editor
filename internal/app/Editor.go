package app

import (
	"LODeditor/internal/app/storage"
	"LODeditor/internal/app/ui"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"github.com/sqweek/dialog"
)

/**
Run : Main App entry point. Assembles and runs the main Fyne app.
 */
func Run() {
	var path string
	var err error
	var card storage.Card
	var slot storage.Slot

	//Load Save File on App Open
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

	w.SetContent(widget.NewVBox(
		ui.Toolbar(&slot, a),
		ui.CreateForm(&slot, &card, w),
	))

	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}
