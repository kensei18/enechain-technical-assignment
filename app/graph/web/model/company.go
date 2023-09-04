package model

import "github.com/kensei18/enechain-technical-assignment/app/entity"

type Company struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewCompany(company entity.Company) *Company {
	return &Company{
		ID:   company.ID.String(),
		Name: company.Name,
	}
}
