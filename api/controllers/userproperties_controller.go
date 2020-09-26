package controllers

import (
	"encoding/json"
	"errors"
	"github.com/MHNassar1/Auth/api/auth"
	"github.com/MHNassar1/Auth/api/models"
	"github.com/MHNassar1/Auth/api/responses"
	"io/ioutil"
	"log"
	"net/http"
)

func AddProperties(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	var props []models.UserProperty
	err = json.Unmarshal(body, &props)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	for _, property := range props {
		property.UserId = tokenID
		err := property.SaveProperty()
		if err != nil {
			log.Fatal("This is the error:", err)
		}
	}
	response := responses.AddProperties{Message: "Add Done !!"}
	responses.JSON(w, http.StatusCreated, response)

}

func UpdateProperties(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	var props []models.UserProperty
	err = json.Unmarshal(body, &props)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	for _, property := range props {
		property.UserId = tokenID
		err := property.UpdateProperty()
		if err != nil {
			log.Fatal("This is the error:", err)
		}
	}
	response := responses.AddProperties{Message: "Add Done !!"}
	responses.JSON(w, http.StatusCreated, response)

}
