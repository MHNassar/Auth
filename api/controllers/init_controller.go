package controllers

import (
	"github.com/MHNassar1/Auth/api/models"
	"github.com/MHNassar1/Auth/api/responses"
	"net/http"
)

func GetInitSettings(w http.ResponseWriter, r *http.Request) {
	service := models.Service{}
	city := models.City{}
	services, err := service.FindAl()
	cities, err := city.FindAl()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	var response responses.InitSettings
	response.Cities = cities
	response.Services = services

	responses.JSON(w, http.StatusOK, response)
}
