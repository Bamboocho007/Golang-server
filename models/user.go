package models

type User struct {
	Id       uint   `json:"id" validate:"required,min=1"`
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required,min=3"`
	Role     int    `json:"role" validate:"required,min=1"`
	Password string `json:"-" validate:"required,min=8"`
	Adress   Adress `json:"adress"`
}

type Adress struct {
	City   string `json:"city" validate:"required,min=10,max=100"`
	Street string `json:"street" validate:"required,min=10,max=100"`
	House  uint   `json:"house" validate:"required,min=1,max=999"`
}
