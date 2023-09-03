package entity

import (
	"database/sql/driver"
	"errors"
	"time"

	"github.com/google/uuid"
)

const (
	TaskStatusTodo    TaskStatus = "Todo"
	TaskStatusOnGoing TaskStatus = "OnGoing"
	TaskStatusDone    TaskStatus = "Done"
)

type TaskStatus string

var TaskStatusInvalidError = errors.New("invalid TaskStatus value")

func ParseTaskStatus(value string) (TaskStatus, error) {
	switch value {
	case string(TaskStatusTodo), string(TaskStatusOnGoing), string(TaskStatusDone):
		return TaskStatus(value), nil
	default:
		return TaskStatus(""), TaskStatusInvalidError
	}
}

func (t *TaskStatus) Scan(value interface{}) error {
	*t = TaskStatus(value.(string))
	return nil
}

func (t TaskStatus) Value() (driver.Value, error) {
	return string(t), nil
}

type Task struct {
	ID          uuid.UUID  `gorm:"<-:false;default:gen_random_uuid()"`
	CompanyID   uuid.UUID  `gorm:"not null"`
	AuthorID    uuid.UUID  `gorm:"not null"`
	Title       string     `gorm:"not null"`
	Description string     `gorm:"not null"`
	Status      TaskStatus `gorm:"not null" sql:"type:task_status"`
	IsPrivate   bool       `gorm:"not null"`
	CreatedAt   time.Time  `gorm:"not null"`
	UpdatedAt   time.Time  `gorm:"not null"`

	TaskAssignees []TaskAssignee
}
