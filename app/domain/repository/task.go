package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/kensei18/enechain-technical-assignment/app/entity"
)

type TaskRepository interface {
	Create(ctx context.Context, params CreateTaskParams) (*entity.Task, error)
	Update(ctx context.Context, params UpdateTaskParams) (*entity.Task, error)
	Delete(ctx context.Context, id uuid.UUID) error
	GetCompanyPublicTasks(ctx context.Context, companyID uuid.UUID) ([]*entity.Task, error)
	GetUserPrivateTasks(ctx context.Context, userID uuid.UUID) ([]*entity.Task, error)
	GetCompanyTasks(ctx context.Context, companyID uuid.UUID) ([]*entity.Task, error)
}

type CreateTaskParams struct {
	CompanyID   uuid.UUID
	AuthorID    uuid.UUID
	Title       string
	Description string
	Status      entity.TaskStatus
	IsPrivate   bool
	AssigneeIDs []uuid.UUID
}

type UpdateTaskParams struct {
	TaskID      uuid.UUID
	Title       *string
	Description *string
	Status      *entity.TaskStatus
	IsPrivate   *bool
	AssigneeIDs []uuid.UUID
}
