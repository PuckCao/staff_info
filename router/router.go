package router

import (
	"github.com/gin-gonic/gin"
	"gorm/controller"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.GET("/user",controller.GetAllUsers)
	r.GET("/user/:id",controller.GetUser)
	r.POST("/user",controller.AddUser)
	//usergroup.PUT("user/:id")
	r.DELETE("/user",controller.DeleteUser)

	return r



}
