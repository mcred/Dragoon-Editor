package ui

import (
	"LODeditor/internal/app/characters"
	"LODeditor/internal/app/common"
	"LODeditor/internal/app/inventory"
	. "LODeditor/internal/app/storage"
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"image/color"
	"strconv"
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

func createCharSelect(i inventory.Inventory, a Attribute, s *Slot) *widget.Select {
	r := widget.NewSelect(i.GetVals(), func(v string) {
		s.SetValueAtLocation(a, i.GetIDByVal(v))
	})
	r.SetSelected(i.GetValByID(s.GetValueByAttribute(a)))
	return r
}

func createPosSelect(n []string, a Attribute, ad Attribute, s *Slot) *widget.Select {
	r := widget.NewSelect(n, func(v string) {
		if v == "Empty" {
			s.SetValueAtLocation(a, 0xFFFFFFFF)
			s.SetValueAtLocation(ad, 0xFFFFFFFF)
		} else {
			s.SetValueAtLocation(a, characters.GetIDByName(v))
			s.SetValueAtLocation(ad, characters.GetIDByName(v))
		}
	})
	char := characters.GetNameByID(s.GetValueByAttribute(a))
	if char != "" {
		r.SetSelected(char)
	} else {
		r.SetSelected("Empty")
	}
	return r
}

func createCharEntry(a Attribute, s *Slot) *widget.Entry {
	e := widget.NewEntry()
	e.SetPlaceHolder(strconv.Itoa(s.GetValueByAttribute(a)))
	e.OnChanged = func(v string) {
		i, _ := strconv.Atoi(v)
		s.SetValueAtLocation(a, i)
	}
	return e
}

func createPartyForm(s *Slot) *fyne.Container {
	names := []string{"Empty"}
	names = append(names, characters.GetCharacterNames()...)
	s1 := createPosSelect(names, common.First(), common.FirstDisplay(), s)
	s2 := createPosSelect(names, common.Second(), common.SecondDisplay(), s)
	s3 := createPosSelect(names, common.Third(), common.ThirdDisplay(), s)
	e1 := widget.NewEntry()
	e1.SetPlaceHolder(strconv.Itoa(s.GetValueByAttribute(common.Gold())))
	e1.OnChanged = func(v string) {
		i, _ := strconv.Atoi(v)
		s.SetValueAtLocation(common.Gold(), i)
		s.SetValueAtLocation(common.GoldDisplay(), i)
	}
	e2 := widget.NewEntry()
	e2.SetPlaceHolder(strconv.Itoa(s.GetValueByAttribute(common.Stardust())))
	e2.OnChanged = func(v string) {
		i, _ := strconv.Atoi(v)
		s.SetValueAtLocation(common.Stardust(), i)
		s.SetValueAtLocation(common.StardustDisplay(), i)
	}

	spacer := canvas.NewRectangle(&color.RGBA{128, 128, 128, 255})
	spacer.SetMinSize(fyne.NewSize(190, 5))

	return fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
		widget.NewLabel("Party"), s1, s2, s3,
		spacer,
		widget.NewLabel("Gold"), e1,
		widget.NewLabel("Stardust"), e2,
	)
}

func createCharacterBox(c characters.Character, w inventory.Inventory, s *Slot) *fyne.Container {
	lb := widget.NewVBox()
	mb := widget.NewVBox()
	rb := widget.NewVBox()
	lb.Append(widget.NewLabel("Level"))
	lb.Append(createCharEntry(c.Level, s))
	lb.Append(widget.NewLabel("Dragoon Level"))
	lb.Append(createCharEntry(c.DLevel, s))
	lb.Append(widget.NewLabel("EXP"))
	lb.Append(createCharEntry(c.XP, s))
	lb.Append(widget.NewLabel("HP"))
	lb.Append(createCharEntry(c.HP, s))
	mb.Append(widget.NewLabel("MP"))
	mb.Append(createCharEntry(c.MP, s))
	mb.Append(widget.NewLabel("SP"))
	mb.Append(createCharEntry(c.SP, s))
	mb.Append(widget.NewLabel("SP Max"))
	mb.Append(createCharEntry(c.SPMax, s))
	rb.Append(widget.NewLabel("Weapon"))
	rb.Append(createCharSelect(w, c.Weapon, s))
	rb.Append(widget.NewLabel("Armor"))
	rb.Append(createCharSelect(inventory.Armor(), c.Chest, s))
	rb.Append(widget.NewLabel("Headgear"))
	rb.Append(createCharSelect(inventory.Helms(), c.Helmet, s))
	rb.Append(widget.NewLabel("Boots"))
	rb.Append(createCharSelect(inventory.Boots(), c.Boots, s))
	rb.Append(widget.NewLabel("Accessories"))
	rb.Append(createCharSelect(inventory.Accessories(), c.Accessory, s))
	container := fyne.NewContainerWithLayout(layout.NewFixedGridLayout(fyne.NewSize(190, 120)), lb, mb, rb)
	return container
}

// CreateForm : Returns the main form to edit the save slot
func CreateForm(slot *Slot, card *Card, w fyne.Window) *fyne.Container {
	form := &widget.Form{
		OnSubmit: func() {
			card.SaveCard()
			dialog.ShowInformation("Information", "Card Saved", w)
		},
	}

	box1 := createCharacterBox(dart, inventory.Swords(), slot)
	box2 := createCharacterBox(shana, inventory.Bows(), slot)
	box3 := createCharacterBox(lavitz, inventory.Spears(), slot)
	box4 := createCharacterBox(rose, inventory.Daggers(), slot)
	box5 := createCharacterBox(haschel, inventory.Knuckles(), slot)
	box6 := createCharacterBox(albert, inventory.Spears(), slot)
	box7 := createCharacterBox(meru, inventory.Maces(), slot)
	box8 := createCharacterBox(kongol, inventory.Axes(), slot)
	box9 := createCharacterBox(miranda, inventory.Bows(), slot)

	tabs := widget.NewTabContainer(
		widget.NewTabItem("Dart", box1),
		widget.NewTabItem("Shana", box2),
		widget.NewTabItem("Lavitz", box3),
		widget.NewTabItem("Rose", box4),
		widget.NewTabItem("Haschel", box5),
		widget.NewTabItem("Albert", box6),
		widget.NewTabItem("Meru", box7),
		widget.NewTabItem("Kongol", box8),
		widget.NewTabItem("Miranda", box9),
	)

	//chars := fyne.NewContainerWithLayout(layout.NewGridLayout(9),
	//	box1, box2, box3, box4, box5, box6, box7, box8, box9)

	submit := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		form)

	main := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		createPartyForm(slot), tabs)

	return fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
		main, submit)
}
