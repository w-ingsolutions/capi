package main

import (
	"github.com/gioapp/gel"
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
		Db:               db.DuoUIdbInit("./../BAZA"),
		PrikazaniElement: &model.WingVrstaRadova{},
		Suma: &model.WingIzabraniElementi{
			UkupanNeophodanMaterijal: map[int]model.WingNeophodanMaterijal{},
		},
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
		Id:    new(gel.Editor),
		Naziv: new(gel.Editor),
		Opis: &gel.Editor{
			SingleLine: false,
		},
		Obracun:  new(gel.Editor),
		Jedinica: new(gel.Editor),
		Cena:     new(gel.Editor),
		Slug:     new(gel.Editor),
		Omogucen: new(gel.CheckBox),
	}
}
