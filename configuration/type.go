package configuration

type StudentDB struct {
	ID          string `form:"id" `
	Username    string `form:"username" `
	Password    string `form:"password"`
	Email       string `form:"email"`
	PhoneNumber string `form:"phoneNumber"`
}
type TeacherDB struct {
	ID          string `form:"teacher_id" `
	Username    string `form:"teacher_username" `
	Password    string `form:"teacher_password"`
	Email       string `form:"teacher_email"`
	PhoneNumber string `form:"teacher_phoneNumber"`
}

type conDB struct {
	PublicHost string
	Port       string
	DBUser     string
	DBAddr     string
	DBPassword string
	DBName     string
}
