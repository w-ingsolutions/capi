package main

import (
	"gioui.org/widget"
	"github.com/w-ingsolutions/c/model"
	"github.com/w-ingsolutions/capi/db"
)

type WingCal struct {
	Naziv            string
	Strana           string
	Edit             bool
	Materijal        map[int]*model.WingMaterijal
	Radovi           model.WingVrstaRadova
	IzbornikRadova   *model.WingVrstaRadova
	Transfered       model.WingCalGrupaRadova
	Db               *db.DuoUIdb
	Client           *model.Client
	PrikazaniElement *model.WingVrstaRadova
	Suma             *model.WingIzabraniElementi
}

func NewWingCal() *WingCal {
	wing := &WingCal{
		Naziv:            "W-ing Solutions - Kalkulator",
		Db:               db.DuoUIdbInit("/home/marcetin/wing/BAZA"),
		PrikazaniElement: &model.WingVrstaRadova{},
		//Suma: &model.WingIzabraniElementi{
		//	UkupanNeophodanMaterijal: map[int]model.WingNeophodanMaterijal{},
		//},
	}
	//wing.NewMaterijal()
	wing.Radovi = model.WingVrstaRadova{
		Id:             0,
		Naziv:          "Radovi",
		Slug:           "radovi",
		Omogucen:       false,
		Baza:           false,
		Element:        false,
		PodvrsteRadova: wing.Db.DbReadAll("radovi"),
	}

	return wing
}

func (w *WingCal) GenerisanjeEdita() (edit *model.EditabilnaPoljaVrsteRadova) {
	//w.EditabilnaPoljaVrsteRadova = make(map[int]*model.EditabilnaPoljaVrsteRadova)
	//for rad, _ := range radovi {
	//	w.EditabilnaPoljaVrsteRadova[rad] =
	return &model.EditabilnaPoljaVrsteRadova{
		Id:    new(widget.Editor),
		Naziv: new(widget.Editor),
		Opis: &widget.Editor{
			SingleLine: false,
		},
		Obracun:  new(widget.Editor),
		Jedinica: new(widget.Editor),
		Cena:     new(widget.Editor),
		Slug:     new(widget.Editor),
		Omogucen: new(widget.Bool),
	}
}
