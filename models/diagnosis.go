package models

import "time"

type Diagnosis struct {
	ID           string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	PatientID    string    `gorm:"type:uuid;not null;index"`
	Diagnosis    string    `gorm:"type:text;not null"`
	Prescription string    `gorm:"type:text;"`
	StartDate    time.Time `gorm:"not null"`
	CreatedAt    time.Time `gorm:"default:current_timestamp"`
}
