package student

import (
	"WEBApplication/configuration"
	"gorm.io/gorm"
)

func InsertStudent(id, username, password, email, phoneNumber string) error {
	db, err := configuration.Open()
	if err != nil {
		return err
	}
	err = student(db, id, username, password, email, phoneNumber)
	if err != nil {
		return err
	}
	return nil
}

func student(db *gorm.DB, id, username, password, email, phoneNumber string) error {
	err := db.Create(&configuration.StudentDB{
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
