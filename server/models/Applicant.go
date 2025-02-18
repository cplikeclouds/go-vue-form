package models

import (
	"gorm.io/gorm"
	"log"
	"server/global"
	"server/utils/errmsg"
)

type Applicant struct {
	gorm.Model
	ApplicantType string `gorm:"type:ENUM('student','employee');not null" json:"applicant_type"`
	Name          string `gorm:"size:50;not null" json:"name"`
	Phone         string `gorm:"size:20" json:"phone,omitempty"`
	Email         string `gorm:"size:100" json:"email,omitempty"`
	Purposes      string `gorm:"type:text" json:"purposes,omitempty"`
	Notes         string `gorm:"type:text" json:"notes,omitempty"`

	// 关联模型
	Employee Employee `gorm:"foreignKey:ApplicantID" json:"employee,omitempty"`
	Student  Student  `gorm:"foreignKey:ApplicantID" json:"student,omitempty"`
}

func CreateApplicant(data *Applicant) int {
	log.Println(data)
	createTable(Applicant{})

	// 开始事务
	tx := global.DB.Begin()
	if tx.Error != nil {
		return errmsg.ERROR
	}

	// 创建 Applicant
	if err := tx.Create(&data).Error; err != nil {
		tx.Rollback()
		return errmsg.ERROR
	}

	if data.Student.School != "" {
		data.Student.ApplicantID = data.ID
		if err := tx.Create(&data.Student).Error; err != nil {
			tx.Rollback()
			return errmsg.ERROR
		}
	}

	if data.Employee.Company != "" {
		data.Employee.ApplicantID = data.ID
		if err := tx.Create(&data.Employee).Error; err != nil {
			tx.Rollback()
			return errmsg.ERROR
		}
	}
	// 提交事务
	err := tx.Commit().Error
	if err != nil {
		return errmsg.ERROR
	}

	return errmsg.SUCCESS

}

func GetApplicantList(pageSize int, pageNum int) ([]Applicant, error) {

	var applicants []Applicant
	err := global.DB.Preload("Student").
		Preload("Employee").
		Limit(pageSize).
		Find(&applicants).Error

	return applicants, err
}

// ID查询用户
func CheckApplicantByID(applicantID uint) (Applicant, int) {
	var Applicant Applicant
	if err := global.DB.Where("id = ?", applicantID).Preload("Student").Preload("employee").First(&Applicant).Error; err == gorm.ErrRecordNotFound {
		return Applicant, errmsg.ERROR
	}
	return Applicant, errmsg.SUCCESS

}

// Name查询用户
func CheckApplicantByName(name string) (Applicant, int) {
	var Applicant Applicant
	if err := global.DB.Where("name = ?", name).Preload("Student").Preload("employee").First(&Applicant).Error; err == gorm.ErrRecordNotFound {
		return Applicant, errmsg.ERROR
	}
	return Applicant, errmsg.SUCCESS

}

// UpdateApplicant 更新报名者数据
func UpdateApplicant(db *gorm.DB, applicant Applicant) error {
	// 开始事务
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// 更新 Applicant 数据
	if err := tx.Save(&applicant).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 更新关联的 Student 数据
	if applicant.Student.ID != 0 {
		if err := tx.Save(&applicant.Student).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// 更新关联的 Employee 数据
	if applicant.Employee.ID != 0 {
		if err := tx.Save(&applicant.Employee).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// 提交事务
	return tx.Commit().Error
}
