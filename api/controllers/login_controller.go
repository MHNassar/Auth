package controllers

import (
	"encoding/json"
	"github.com/MHNassar1/Auth/api/auth"
	"github.com/MHNassar1/Auth/api/models"
	"github.com/MHNassar1/Auth/api/responses"
	"github.com/MHNassar1/Auth/api/utils/formaterror"
	"golang.org/x/crypto/bcrypt"
	"github.com/MHNassar1/Auth/api/core"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := SignIn(user.Email, user.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}

	loginResponse := responses.Login{}
	loginResponse.Token = token

	responses.JSON(w, http.StatusOK, loginResponse)
}

func  SignIn(email, password string) (string, error) {

	var err error
    
	user := models.User{}
    
	err = core.AppInstance.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(user.ID)
}
