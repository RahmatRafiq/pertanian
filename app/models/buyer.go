package models

import (
	"time"

	"github.com/google/uuid"
)

type Buyer struct {
	UserID         uuid.UUID  `gorm:"type:uuid;primaryKey;column:user_id"`
	CompanyName    string     `gorm:"type:varchar(255);not null;column:company_name"`
	ContactDetails string     `gorm:"type:text;column:contact_details"`
	CreatedAt      time.Time  `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt      time.Time  `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt      *time.Time `gorm:"column:deleted_at;index"`
}
