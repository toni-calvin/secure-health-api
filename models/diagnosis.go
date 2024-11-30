package models

import "time"

type Diagnosis struct {
	ID        string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID    string    `gorm:"type:uuid;not null;index"`
	Diagnosis string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"default:current_timestamp"`
}
