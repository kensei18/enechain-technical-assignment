package infrastructure

import (
	"context"

	"github.com/google/uuid"
	"github.com/kensei18/enechain-technical-assignment/app/domain/repository"
	"github.com/kensei18/enechain-technical-assignment/app/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) repository.UserRepository {
	return &UserRepository{conn: conn}
}

func (r *UserRepository) FindByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	user := &entity.User{ID: id}
	if err := r.conn.First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
