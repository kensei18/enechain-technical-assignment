package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Company struct {
	ID        uuid.UUID `gorm:"<-:false;default:gen_random_uuid()"`
	Name      string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
	DeletedAt gorm.DeletedAt
}
