package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	ApplicantID    uint   `gorm:"not null" json:"-"`
	Company        string `gorm:"size:100;not null" json:"company"`
	Position       string `gorm:"size:50;not null" json:"position"`
	Industry       string `gorm:"type:ENUM('IT','finance','education','other');not null" json:"industry"`
	WorkExperience string `gorm:"size:50;" json:"work_experience"`
}
