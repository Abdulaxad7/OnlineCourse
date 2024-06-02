package teacher

import (
	"WEBApplication/configuration"
	"gorm.io/gorm"
)

func InsertTeacher(id, username, password, email, phoneNumber string) error {
	db, err := configuration.OpenTeacher()
	if err != nil {
		return err
	}
	err = teacher(db, id, username, password, email, phoneNumber)
	if err != nil {
		return err
	}
	return nil
}

func teacher(db *gorm.DB, id, username, password, email, phoneNumber string) error {
	err := db.Create(&configuration.TeacherDB{
		ID:          id,
		Username:    username,
		Password:    password,
		Email:       email,
		PhoneNumber: phoneNumber,
	})
	if err.Error != nil {
		return err.Error
	}
	return nil
}
