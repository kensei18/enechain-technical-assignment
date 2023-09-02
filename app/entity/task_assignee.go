package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaskAssignee struct {
	ID         uuid.UUID `gorm:"<-:false;default:gen_random_uuid()"`
	TaskID     uuid.UUID `gorm:"not null"`
	AssigneeID uuid.UUID `gorm:"not null"`
	CreatedAt  time.Time `gorm:"not null"`
	UpdatedAt  time.Time `gorm:"not null"`
	DeletedAt  gorm.DeletedAt

	Task     Task
	Assignee User
}
