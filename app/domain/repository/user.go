package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/kensei18/enechain-technical-assignment/app/entity"
)

type UserRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
}
