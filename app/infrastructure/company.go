package infrastructure

import (
	"github.com/kensei18/enechain-technical-assignment/app/domain/repository"
	"gorm.io/gorm"
)

type CompanyRepository struct {
	conn *gorm.DB
}

func NewCompanyRepository(conn *gorm.DB) repository.CompanyRepository {
	return CompanyRepository{conn: conn}
}
