package controllers

import (
	"github.com/MHNassar1/Auth/api/responses"
	"net/http"
	)

func Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
