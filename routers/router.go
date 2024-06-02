package routers

import (
	"WEBApplication/configuration"
	"WEBApplication/user/student"
	"WEBApplication/user/teacher"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", nil)
}

func LoginHandlerStudentGet(context *gin.Context) {
	context.HTML(http.StatusOK, "loginStudent.html", nil)
}
func LoginHandlerStudentPost(context *gin.Context) {
	var form configuration.StudentDB
	if err := context.ShouldBind(&form); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if (student.UserId(form.ID) || student.Username(form.ID) || student.UserEmail(form.ID)) && student.UserPass(form.Password) {
		context.SetCookie("user", form.ID, 3600, "/", "localhost", false, true)
		context.Redirect(http.StatusFound, "/dashboard")
		context.HTML(http.StatusOK, "home.html", nil)
	} else {
		context.HTML(http.StatusOK, "error.html", nil)
		return

	}
}

func LoginHandlerTeacherGet(context *gin.Context) {
	context.HTML(http.StatusOK, "loginTeacher.html", nil)
}
func LoginHandlerTeacherPost(context *gin.Context) {
	var form configuration.StudentDB
	if err := context.ShouldBind(&form); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if (teacher.UserId(form.ID) || teacher.Username(form.ID) || teacher.UserEmail(form.ID)) && teacher.UserPass(form.Password) {
		context.Redirect(http.StatusFound, "/dashboard")
		context.HTML(http.StatusOK, "home.html", nil)
	} else {
		context.HTML(http.StatusFound, "error.html", nil)
		return

	}
}

func Help(c *gin.Context) {
	c.HTML(http.StatusOK, "help.html", nil)
}

func CreateAccountStudentGet(c *gin.Context) {
	c.HTML(http.StatusOK, "createAccount.html", nil)
}
func CreateAccountStudentPost(context *gin.Context) {
	var form configuration.StudentDB
	if err := context.ShouldBind(&form); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := student.InsertStudent(form.ID, form.Username, form.Password, form.Email, form.PhoneNumber); err == nil {
		context.Redirect(http.StatusFound, "/dashboard")
		context.HTML(http.StatusOK, "home.html", nil)
	} else {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	}
}
func CreateAccountTeacherGet(c *gin.Context) {
	c.HTML(http.StatusOK, "createAccount.html", nil)
}
func CreateAccountTeacherPost(context *gin.Context) {
	var form configuration.TeacherDB
	if err := context.ShouldBind(&form); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := teacher.InsertTeacher(form.ID, form.Username, form.Password, form.Email, form.PhoneNumber); err == nil {
		context.Redirect(http.StatusFound, "/dashboard")
		context.HTML(http.StatusOK, "home.html", nil)
	} else {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	}
}
