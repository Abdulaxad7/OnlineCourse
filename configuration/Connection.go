package configuration

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
)

func Connect() (string, error) {
	_, err := getConnection()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Configuration.DBUser, Configuration.DBPassword, Configuration.DBAddr, Configuration.DBName), err
}
func MYSQL(my mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", my.FormatDSN())
	if err != nil {
		return nil, err
	}
	return db, nil
}

func getConnection() (*sql.DB, error) {
	return MYSQL(mysql.Config{
		User:                 Configuration.DBUser,
		Passwd:               Configuration.DBPassword,
		Addr:                 Configuration.DBAddr,
		DBName:               Configuration.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
}
