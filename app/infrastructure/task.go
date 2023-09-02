package infrastructure

import (
	"github.com/kensei18/enechain-technical-assignment/app/domain/repository"
	"gorm.io/gorm"
)

type TaskRepository struct {
	conn *gorm.DB
}

func NewTaskRepository(conn *gorm.DB) repository.TaskRepository {
	return TaskRepository{conn: conn}
}
