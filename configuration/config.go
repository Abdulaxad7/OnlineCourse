package configuration

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Open() (*gorm.DB, error) {
	s, err := Connect()
	if err != nil {
		return nil, err
	}
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	err = db.AutoMigrate(&StudentDB{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
func OpenTeacher() (*gorm.DB, error) {
	s, err := Connect()
	if err != nil {
		return nil, err
	}
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	err = db.AutoMigrate(&TeacherDB{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
