package config

import (
	"fmt"
	"log"
	"net/http"

	"github.com/majalah/apps/controller"

	"github.com/gorilla/mux"
)

func GetRoute() {
	r := mux.NewRouter()
	r.HandleFunc("/movies", controller.AllMoviesEndPoint).Methods("GET")
	r.HandleFunc("/movies", controller.CreateMovieEndPoint).Methods("POST")
	r.HandleFunc("/movies", controller.UpdateMovieEndPoint).Methods("PUT")
	r.HandleFunc("/movies", controller.DeleteMovieEndPoint).Methods("DELETE")
	r.HandleFunc("/movies/{id}", controller.FindMovieEndpoint).Methods("GET")

	//running process from route
	//it's run at port 80
	err := http.ListenAndServe(":80", r)
	fmt.Println("running.....")
	if err != nil {
		log.Fatal(err)
	}
}
