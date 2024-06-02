package main

import (
	"WEBApplication/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.StaticFile("static/", "./static")
	router.LoadHTMLGlob("static/*")
	router.GET("/", routers.HomeHandler)
	router.GET("/loginStudent", routers.LoginHandlerStudentGet)
	router.POST("/loginStudent", routers.LoginHandlerStudentPost)
	router.GET("/loginTeacher", routers.LoginHandlerTeacherGet)
	router.POST("/loginTeacher", routers.LoginHandlerTeacherPost)
	router.GET("/help", routers.Help)
	router.GET("/createAccountStudent", routers.CreateAccountStudentGet)
	router.POST("/createAccountStudent", routers.CreateAccountStudentPost)
	router.GET("/createAccountTeacher", routers.CreateAccountTeacherGet)
	router.POST("/createAccountTeacher", routers.CreateAccountTeacherPost)

	router.Run(":8080")
}
