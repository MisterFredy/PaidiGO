package dao

import (
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"
)

type Core struct {
	Server   string
	Database string
}

var db *mgo.Database

func (m *Core) Koneksi() {
	koneksi, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = koneksi.DB(m.Database)
}

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondWithJson(w, code, map[string]string{"error": msg})
}

func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
