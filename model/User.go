package model

import "github.com/google/uuid"

type User struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role string `json:"role"`
}

func (user *User) BeforeCreate() {
	user.ID = uuid.New()
}
