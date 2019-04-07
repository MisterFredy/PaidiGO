package dao

import (
	"log"

	. "github.com/belajar/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MajalahDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	collection = "majalah"
)

func (m *MajalahDAO) Koneksi() {
	koneksi, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = koneksi.DB(m.Database)
}

func (m *MajalahDAO) FindMajalah() ([]Majalah, error) {
	var majalah []Majalah
	err := db.C(collection).Find(bson.M{}).All(&majalah)
	return majalah, err
}

func (m *MajalahDAO) FindMajalahById(id string) (Majalah, error) {
	var majalah Majalah
	err := db.C(collection).FindId(bson.ObjectIdHex(id)).One(&majalah)
	return majalah, err
}

func (m *MajalahDAO) CreateMajalahDAO(majalah Majalah) error {
	err := db.C(collection).Insert(&majalah)
	return err
}

func (m *MajalahDAO) UpdateMajalahDAO(id string, majalah Majalah) error {
	err := db.C(collection).Update(bson.ObjectIdHex(id), &majalah)
	return err
}

func (m *MajalahDAO) DeleteMajalahDAO(id string) error {
	err := db.C(collection).Remove(bson.ObjectIdHex(id))
	return err
}
