package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/w-ingsolutions/c/model"
	"net/http"
	"strconv"
)

type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var posts []Post

func (wc *WingCal) VrsteRadova(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	radovi := make(map[int]model.ElementMenu)
	for vr, rd := range wc.Radovi.PodvrsteRadova {
		radovi[vr] = model.ElementMenu{
			Id:    rd.Id,
			Title: rd.Naziv,
			Slug:  rd.Slug,
		}
	}
	json.NewEncoder(w).Encode(radovi)
}

func (wc *WingCal) PodvrsteRadova(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	radovi := make(map[int]model.ElementMenu)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
	}
	for vr, rd := range wc.Radovi.PodvrsteRadova[id+1].PodvrsteRadova {
		radovi[vr] = model.ElementMenu{
			Id:    rd.Id,
			Title: rd.Naziv,
			Slug:  rd.Slug,
		}
	}
	json.NewEncoder(w).Encode(radovi)
}

func (wc *WingCal) Elementi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	radovi := make(map[int]model.ElementMenu)
	elementi := wc.Db.DbRead(params["id"], params["el"])
	for vr, rd := range elementi.PodvrsteRadova {
		var m bool
		if rd.NeophodanMaterijal != nil {
			m = true
		}
		radovi[vr] = model.ElementMenu{
			Id:        rd.Id,
			Title:     rd.Naziv,
			Slug:      rd.Slug,
			Materijal: m,
		}
	}
	json.NewEncoder(w).Encode(radovi)
}

func (wc *WingCal) Element(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	//id, err := strconv.Atoi(params["id"])
	//if err != nil{}
	//el, err := strconv.Atoi(params["el"])
	//if err != nil{}
	e, err := strconv.Atoi(params["e"])
	if err != nil {
	}

	elementi := wc.Db.DbRead(params["id"], params["el"])

	json.NewEncoder(w).Encode(elementi.PodvrsteRadova[e-1])
}

//func (wc *WingCal)getPost(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	params := mux.Vars(r)
//	for _, item := range posts {
//		if item.ID == params["id"] {
//			json.NewEncoder(w).Encode(item)
//			return
//		}
//	}
//	json.NewEncoder(w).Encode(&Post{})
//}

func main() {
	wing := NewWingCal()

	router := mux.NewRouter()
	posts = append(posts, Post{ID: "1", Title: "My first post", Body: "This is the content of my first post"})
	router.HandleFunc("/radovi", wing.VrsteRadova).Methods("GET")
	router.HandleFunc("/radovi/{id}", wing.PodvrsteRadova).Methods("GET")
	router.HandleFunc("/radovi/{id}/{el}", wing.Elementi).Methods("GET")
	router.HandleFunc("/radovi/{id}/{el}/{e}", wing.Element).Methods("GET")

	//router.HandleFunc("/radovi/{id}", wing.getPost).Methods("GET")
	http.ListenAndServe(":9909", router)
}
