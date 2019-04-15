package dao

import "gopkg.in/mgo.v2/bson"

const (
	collection = "user"
)

func FindUser() ([]User, error) {
	var User []User
	err := db.C(collection).Find(bson.M{}).All(&User)
	return User, err
}

func FindUserById(id string) (User, error) {
	var User User
	err := db.C(collection).FindId(bson.ObjectIdHex(id)).One(&User)
	return User, err
}

func CreateUserDAO(User User) error {
	err := db.C(collection).Insert(&User)
	return err
}

func UpdateUserDAO(id string, User User) error {
	err := db.C(collection).Update(bson.ObjectIdHex(id), &User)
	return err
}

func DeleteUserDAO(id string) error {
	err := db.C(collection).Remove(bson.ObjectIdHex(id))
	return err
}
