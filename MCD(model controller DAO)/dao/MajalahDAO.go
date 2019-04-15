package dao

import (
	. "github.com/belajar/models"
	"gopkg.in/mgo.v2/bson"
)

const (
	collection = "majalah"
)

func FindMajalah() ([]Majalah, error) {
	var majalah []Majalah
	err := db.C(collection).Find(bson.M{}).All(&majalah)
	return majalah, err
}

func FindMajalahById(id string) (Majalah, error) {
	var majalah Majalah
	err := db.C(collection).FindId(bson.ObjectIdHex(id)).One(&majalah)
	return majalah, err
}

func CreateMajalahDAO(majalah Majalah) error {
	err := db.C(collection).Insert(&majalah)
	return err
}

func UpdateMajalahDAO(id string, majalah Majalah) error {
	err := db.C(collection).Update(bson.ObjectIdHex(id), &majalah)
	return err
}

func DeleteMajalahDAO(id string) error {
	err := db.C(collection).Remove(bson.ObjectIdHex(id))
	return err
}
