package controller

import (
	"encoding/json"
	"net/http"

	"github.com/belajar/dao"
	. "github.com/belajar/dao"
	. "github.com/belajar/models"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var majalahdao = MajalahDAO{}

func AllMajalah(response http.ResponseWriter, request *http.Request) {
	majalah, err := majalahdao.FindMajalah()
	if err != nil {
		dao.RespondWithError(response, http.StatusInternalServerError, err.Error())
	}
	dao.RespondWithJson(response, http.StatusOK, majalah)
}

func FindMajalahbyId(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	majalah, err := majalahdao.FindMajalahById(params["id"])
	if err != nil {
		dao.RespondWithError(response, http.StatusInternalServerError, "error function Majalhbyid")
	}
	dao.RespondWithJson(response, http.StatusAccepted, majalah)
}

func CreateMajalah(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	var majalah Majalah
	if err := json.NewDecoder(request.Body).Decode(&majalah); err != nil {
		dao.RespondWithError(response, http.StatusBadRequest, "Invalid request payload")
		return
	}
	majalah.ID = bson.NewObjectId()
	if err := majalahdao.CreateMajalahDAO(majalah); err != nil {
		dao.RespondWithError(response, http.StatusInternalServerError, "error insert data")
	}
	dao.RespondWithJson(response, http.StatusCreated, majalah)
}

func DeleteMajalah(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	err := majalahdao.DeleteMajalahDAO(params["id"])
	if err != nil {
		dao.RespondWithError(response, http.StatusInternalServerError, "error delete majalah")
	}
	dao.RespondWithJson(response, http.StatusAccepted, "delete sukses")
}

func UpdateMajalh(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	params := mux.Vars(request)
	var majalah Majalah
	if err := json.NewDecoder(request.Body).Decode(&majalah); err != nil {
		dao.RespondWithError(response, http.StatusInternalServerError, "error decode json")
	}
	err := majalahdao.UpdateMajalahDAO(params["id"], majalah)
	if err != nil {
		dao.RespondWithError(response, http.StatusInternalServerError, "error update majalah")
	}
	dao.RespondWithJson(response, http.StatusOK, majalah)
}
