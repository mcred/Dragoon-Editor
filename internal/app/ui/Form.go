package ui

import (
	"LODeditor/internal/app/characters"
	"LODeditor/internal/app/common"
	"LODeditor/internal/app/inventory"
	. "LODeditor/internal/app/storage"
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
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

	return fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		widget.NewLabel("Party"), s1, s2, s3,
		widget.NewLabel("Gold"), e1,
		widget.NewLabel("Stardust"), e2,
	)
}

func createCharacterBox(b *widget.Box, c characters.Character, w inventory.Inventory, s *Slot) {
	b.Append(widget.NewLabel(c.Name))
	b.Append(widget.NewLabel("Level"))
	b.Append(createCharEntry(c.Level, s))
	b.Append(widget.NewLabel("Dragoon Level"))
	b.Append(createCharEntry(c.DLevel, s))
	b.Append(widget.NewLabel("EXP"))
	b.Append(createCharEntry(c.XP, s))
	b.Append(widget.NewLabel("HP"))
	b.Append(createCharEntry(c.HP, s))
	b.Append(widget.NewLabel("MP"))
	b.Append(createCharEntry(c.MP, s))
	b.Append(widget.NewLabel("SP"))
	b.Append(createCharEntry(c.SP, s))
	b.Append(widget.NewLabel("SP Max"))
	b.Append(createCharEntry(c.SPMax, s))
	//b.Append(widget.NewLabel("Weapon"))
	b.Append(createCharSelect(w, c.Weapon, s))
	//b.Append(widget.NewLabel("Armor"))
	b.Append(createCharSelect(inventory.Armor(), c.Chest, s))
	//b.Append(widget.NewLabel("Headgear"))
	b.Append(createCharSelect(inventory.Helms(), c.Helmet, s))
	//b.Append(widget.NewLabel("Boots"))
	b.Append(createCharSelect(inventory.Boots(), c.Boots, s))
	//b.Append(widget.NewLabel("Accessories"))
	b.Append(createCharSelect(inventory.Accessories(), c.Accessory, s))
}

// CreateForm : Returns the main form to edit the save slot
func CreateForm(slot *Slot, card *Card, w fyne.Window) *fyne.Container {
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
	box5 := widget.NewVBox()
	createCharacterBox(box5, haschel, inventory.Knuckles(), slot)
	box6 := widget.NewVBox()
	createCharacterBox(box6, albert, inventory.Spears(), slot)
	box7 := widget.NewVBox()
	createCharacterBox(box7, meru, inventory.Maces(), slot)
	box8 := widget.NewVBox()
	createCharacterBox(box8, kongol, inventory.Axes(), slot)
	box9 := widget.NewVBox()
	createCharacterBox(box9, miranda, inventory.Bows(), slot)

	chars := fyne.NewContainerWithLayout(layout.NewGridLayout(9),
		box1, box2, box3, box4, box5, box6, box7, box8, box9)
	submit := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		form)

	return fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
		createPartyForm(slot), chars, submit)
}
