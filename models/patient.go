package models

import "time"

type Patient struct {
	ID        string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string    `gorm:"type:varchar(255);uniqueIndex;not null"`
	NIF       string    `gorm:"type:varchar(255);uniqueIndex;not null"`
	Email     string    `gorm:"type:varchar(255);not null"`
	Phone     string    `gorm:"type:varchar(255);"`
	Address   string    `gorm:"type:varchar(255);"`
	CreatedAt time.Time `gorm:"default:current_timestamp"`
}
