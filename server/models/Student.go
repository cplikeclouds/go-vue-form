package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	ApplicantID    uint   `gorm:"not null" json:"-"`
	School         string `gorm:"size:100;not null" json:"school"`
	Major          string `gorm:"size:50;not null" json:"major"`
	EnrollmentYear string `gorm:"type:year" json:"enrollment_year,omitempty"`
	GenerateYear   string `gorm:"type:year" json:"generate_year,omitempty"`
}
