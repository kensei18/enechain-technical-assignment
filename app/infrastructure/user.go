package infrastructure

import (
	"github.com/kensei18/enechain-technical-assignment/app/domain/repository"
	"gorm.io/gorm"
)

type UserRepository struct {
	conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) repository.UserRepository {
	return UserRepository{conn: conn}
}
