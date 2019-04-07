package config

import (
	"log"
	"net/http"

	"github.com/belajar/controller"
	"github.com/gorilla/mux"
)

type Network struct {
	Port string
}

func (n *Network) Route() {
	r := mux.NewRouter()
	r.HandleFunc("/majalah", controller.AllMajalah).Methods("GET")
	r.HandleFunc("/findmajalah/{id}", controller.FindMajalahbyId).Methods("GET")
	r.HandleFunc("/createmajalah", controller.CreateMajalah).Methods("POST")
	r.HandleFunc("/delmajalah/{id}", controller.DeleteMajalah).Methods("DELETE")
	r.HandleFunc("/updatemajalah/{id}", controller.UpdateMajalh).Methods("PUT")
	if err := http.ListenAndServe(n.Port, r); err != nil {
		log.Fatal(err)
	}
	log.Fatal("running on server 80")
}
