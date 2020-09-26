package responses

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
