package entity

import (
	"database/sql/driver"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	TaskStatusTodo    TaskStatus = "Todo"
	TaskStatusOnGoing TaskStatus = "OnGoing"
	TaskStatusDone    TaskStatus = "Done"
)

type TaskStatus string

func (t *TaskStatus) Scan(value interface{}) error {
	*t = TaskStatus(value.([]byte))
	return nil
}

func (t TaskStatus) Value() (driver.Value, error) {
	return string(t), nil
}

type Task struct {
	ID          uuid.UUID  `gorm:"<-:false;default:gen_random_uuid()"`
	AuthorID    uuid.UUID  `gorm:"not null"`
	Title       string     `gorm:"not null"`
	Description string     `gorm:"not null"`
	Status      TaskStatus `gorm:"not null" sql:"type:task_status"`
	IsPrivate   bool       `gorm:"not null"`
	CreatedAt   time.Time  `gorm:"not null"`
	UpdatedAt   time.Time  `gorm:"not null"`
	DeletedAt   gorm.DeletedAt
}
