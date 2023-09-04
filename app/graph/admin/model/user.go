package model

import "github.com/kensei18/enechain-technical-assignment/app/entity"

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	CompanyID string
}

func NewUser(user entity.User) *User {
	return &User{
		ID:        user.ID.String(),
		Name:      user.Name,
		CompanyID: user.CompanyID.String(),
	}
}
