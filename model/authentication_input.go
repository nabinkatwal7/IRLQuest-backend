package model

type AuthenticationInputRegister struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type AuthenticationInputLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}