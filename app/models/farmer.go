package models

import (
	"time"
)

type Farmer struct {
	ID          uint64     `gorm:"primaryKey;autoIncrement;column:id"`
	UserID      uint64     `gorm:"primaryKey;column:user_id"`
	Name        string     `gorm:"type:varchar(255);not null;column:name"`
	PhoneNumber string     `gorm:"type:varchar(20);column:phone_number"`
	NationalID  string     `gorm:"type:varchar(100);not null;column:national_id"`
	Address     string     `gorm:"type:text;column:address"`
	BirthDate   *time.Time `gorm:"type:date;column:birth_date"`
	Gender      string     `gorm:"type:enum('MALE','FEMALE','OTHER');column:gender"`
	CreatedAt   time.Time  `gorm:"column:created_at;autoCreateTime;default:now()" swaggerignore:"true"`
	UpdatedAt   time.Time  `gorm:"column:updated_at;autoUpdateTime;" swaggerignore:"true"`
	DeletedAt   *time.Time `gorm:"column:deleted_at" swaggerignore:"true"`
}
