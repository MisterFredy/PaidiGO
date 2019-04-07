package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Majalah struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	Description string        `bson:"description" json:"description"`
	Cover       string        `bson:"cover" json:"cover"`
}
