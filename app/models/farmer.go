package models

import (
	"time"

	"github.com/google/uuid"
)

type Farmer struct {
	UserID      uuid.UUID  `gorm:"type:uuid;primaryKey;column:user_id"`
	PhoneNumber string     `gorm:"type:varchar(20);column:phone_number"`
	CreatedAt   time.Time  `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time  `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt   *time.Time `gorm:"column:deleted_at"`
}
