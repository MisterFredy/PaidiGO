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
	/*
	* route Majalah
	 */
	r.HandleFunc("/majalah", controller.AllMajalah).Methods("GET")
	r.HandleFunc("/majalah/{id}", controller.FindMajalahbyId).Methods("GET")
	r.HandleFunc("/majalah", controller.CreateMajalah).Methods("POST")
	r.HandleFunc("/majalah/{id}", controller.DeleteMajalah).Methods("DELETE")
	r.HandleFunc("/majalah/{id}", controller.UpdateMajalh).Methods("PUT")

	/*
	* route User
	 */
	r.HandleFunc("/user", controller.TambahUser).Methods("POST")
	r.HandleFunc("/auth", controller.Login).Methods("POST")
	r.HandleFunc("/user/{id}", controller.InfoUser).Methods("GET")
	r.HandleFunc("/user/{id}", controller.UpdateUser).Methods("POST")

	if err := http.ListenAndServe(n.Port, r); err != nil {
		log.Fatal(err)
	}
	log.Fatal("running on server 80")
}
