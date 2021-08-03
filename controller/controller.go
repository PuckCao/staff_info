package controller

import (
	"github.com/gin-gonic/gin"
	"gorm/model"
	"net/http"
	"strconv"
)

func GetAllUsers(c *gin.Context)  {
	Users , err:= model.FindAll()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK,Users)
	}

}
func GetUser(c *gin.Context) {

	id , _:= strconv.Atoi(c.Param("id"))
	User,err := model.FindOne(id)
	if err != nil {
		c.JSON(http.StatusNotFound,gin.H{"error": err.Error()})
	}else {
		c.JSON(http.StatusOK,User)
	}
}
func AddUser(c *gin.Context)  {
		name := c.PostForm("name")
		age,_ := strconv.Atoi(c.PostForm("age"))
		sex := c.PostForm("sex")
		department := c.PostForm("department")
		level, _ := strconv.Atoi(c.PostForm("level"))
		tel,_ := strconv.Atoi(c.PostForm("tel"))
		var u = model.User{
			Name: name,
			Age:  age,
			Sex:  sex,
			Department: department,
			Level: level,
			Tel: tel,
		}
		res := model.DB.Create(&u)
		if res.Error != nil{
			c.JSON(http.StatusOK,gin.H{
				"msg":"插入失败",
			})
		}else {
			c.JSON(http.StatusOK,gin.H{
				"msg":"插入成功",
			})
		}
}

func DeleteUser(c *gin.Context) {
	id , _:= strconv.Atoi(c.PostForm("id"))
	model.Delete(id)
	c.JSON(http.StatusOK,gin.H{
			"msg":"删除成功",
		})
}

