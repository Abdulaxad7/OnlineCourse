package teacher

import (
	"WEBApplication/configuration"
	"log"
)

func UserId(s string) bool {
	var id []string
	db, err := configuration.OpenTeacher()
	if err != nil {
		log.Fatal(err)
	}

	db.Model(&configuration.TeacherDB{}).Pluck("id", &id)
	for _, v := range id {
		if s == v {
			return true
		}
	}
	return false
}
func UserPass(s string) bool {
	var pass []string
	db, err := configuration.OpenTeacher()
	if err != nil {
		log.Fatal(err)
	}

	db.Model(&configuration.TeacherDB{}).Pluck("password", &pass)
	for _, v := range pass {
		if s == v {
			return true
		}
	}
	return false
}
func Username(s string) bool {
	var name []string
	db, err := configuration.OpenTeacher()
	if err != nil {
		log.Fatal(err)
	}

	db.Model(&configuration.TeacherDB{}).Pluck("username", &name)
	for _, v := range name {
		if s == v {
			return true
		}
	}
	return false
}
func UserEmail(s string) bool {
	var email []string
	db, err := configuration.OpenTeacher()
	if err != nil {
		log.Fatal(err)
	}

	db.Model(&configuration.TeacherDB{}).Pluck("email", &email)
	for _, v := range email {
		if s == v {
			return true
		}
	}
	return false
}
