package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/kensei18/enechain-technical-assignment/app/domain/repository"
	"github.com/kensei18/enechain-technical-assignment/app/entity"
	"github.com/kensei18/enechain-technical-assignment/app/infrastructure"
	"gorm.io/gorm"
)

type taskService struct {
	DB *gorm.DB
}

func NewTaskService(db *gorm.DB) *taskService {
	return &taskService{DB: db}
}

type CreateTaskParams struct {
	AuthorID    uuid.UUID
	Title       string
	Description string
	Status      entity.TaskStatus
	IsPrivate   bool
	AssigneeIDs []uuid.UUID
}

func (s *taskService) CreateTask(ctx context.Context, params CreateTaskParams) (*entity.Task, error) {
	var task *entity.Task

	err := s.DB.Transaction(func(tx *gorm.DB) error {
		userRepository := infrastructure.NewUserRepository(tx)
		taskRepository := infrastructure.NewTaskRepository(tx)

		user, err := userRepository.FindByID(ctx, params.AuthorID)
		if err != nil {
			return err
		}
		t, err := taskRepository.Create(ctx, repository.CreateTaskParams{
			CompanyID:   user.CompanyID,
			AuthorID:    user.ID,
			Title:       params.Title,
			Description: params.Description,
			Status:      params.Status,
			IsPrivate:   params.IsPrivate,
			AssigneeIDs: params.AssigneeIDs,
		})
		task = t
		return err
	})

	return task, err
}

func (s *taskService) UpdateTask(ctx context.Context, params repository.UpdateTaskParams) (*entity.Task, error) {
	var task *entity.Task

	err := s.DB.Transaction(func(tx *gorm.DB) error {
		taskRepository := infrastructure.NewTaskRepository(tx)
		t, err := taskRepository.Update(ctx, params)
		task = t
		return err
	})

	return task, err
}

func (s *taskService) DeleteTask(ctx context.Context, taskID uuid.UUID) error {
	return s.DB.Transaction(func(tx *gorm.DB) error {
		taskRepository := infrastructure.NewTaskRepository(tx)
		return taskRepository.Delete(ctx, taskID)
	})
}
