package responses

import "github.com/MHNassar1/Auth/api/models"

type Login struct {
	Token string `json:"token"`
}

type CreateUser struct {
	Token string      `json:"token"`
	User  interface{} `json:"user"`
}

type AddProperties struct {
	Message interface{} `json:"message"`
}

type InitSettings struct {
	Cities   *[]models.City    `json:"cities"`
	Services *[]models.Service `json:"services"`
}
