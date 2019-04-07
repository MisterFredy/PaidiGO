package main

import (
	"encoding/json"
	"log"
	"net/http"

	. "github.com/belajar/config"
	. "github.com/belajar/dao"
	. "github.com/belajar/models"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var config = Config{}
var dao = MajalahDAO{}

func AllMajalah(response http.ResponseWriter, request *http.Request) {
	majalah, err := dao.FindMajalah()
	if err != nil {
		respondWithError(response, http.StatusInternalServerError, err.Error())
	}
	respondWithJson(response, http.StatusOK, majalah)
}

func FindMajalahbyId(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	majalah, err := dao.FindMajalahById(params["id"])
	if err != nil {
		respondWithError(response, http.StatusInternalServerError, "error function Majalhbyid")
	}
	respondWithJson(response, http.StatusAccepted, majalah)
}

func CreateMajalah(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	var majalah Majalah
	if err := json.NewDecoder(request.Body).Decode(&majalah); err != nil {
		respondWithError(response, http.StatusBadRequest, "Invalid request payload")
		return
	}
	majalah.ID = bson.NewObjectId()
	if err := dao.CreateMajalahDAO(majalah); err != nil {
		respondWithError(response, http.StatusInternalServerError, "error insert data")
	}
	respondWithJson(response, http.StatusCreated, majalah)
}

func DeleteMajalah(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	err := dao.DeleteMajalahDAO(params["id"])
	if err != nil {
		respondWithError(response, http.StatusInternalServerError, "error delete majalah")
	}
	respondWithJson(response, http.StatusAccepted, "delete sukses")
}

func UpdateMajalh(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	params := mux.Vars(request)
	var majalah Majalah
	if err := json.NewDecoder(request.Body).Decode(&majalah); err != nil {
		respondWithError(response, http.StatusInternalServerError, "error decode json")
	}
	err := dao.UpdateMajalahDAO(params["id"], majalah)
	if err != nil {
		respondWithError(response, http.StatusInternalServerError, "error update majalah")
	}
	respondWithJson(response, http.StatusOK, majalah)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func init() {
	config.Read()
	dao.Database = config.Database
	dao.Server = config.Server
	dao.Koneksi()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/majalah", AllMajalah).Methods("GET")
	r.HandleFunc("/findmajalah/{id}", FindMajalahbyId).Methods("GET")
	r.HandleFunc("/createmajalah", CreateMajalah).Methods("POST")
	r.HandleFunc("/delmajalah/{id}", DeleteMajalah).Methods("DELETE")
	r.HandleFunc("/updatemajalah/{id}", UpdateMajalh).Methods("PUT")
	if err := http.ListenAndServe(":80", r); err != nil {
		log.Fatal(err)
	}
	log.Fatal("running on server 80")
}
