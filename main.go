package main

import (
	"gorm/model"
	"gorm/router"
)




func main() {
    model.InitDB()
    model.DB.AutoMigrate(&model.User{})
	//r.POST("/user", func(c *gin.Context) {
	//	db, _ := gorm.Open(mysql.New(mysql.Config{
	//		DSN: "root:1997@tcp(127.0.0.1:3306)/db1?charset=utf8&parseTime=True&loc=Local"}), &gorm.Config{})
	//	db.AutoMigrate(&model.User{})
	//	name := c.PostForm("name")
	//	age0 := (c.PostForm("age"))
	//	age1, _ := strconv.Atoi(age0)
	//	sex := c.PostForm("sex")
	//	department := c.PostForm("department")
	//	level, _ := strconv.Atoi(c.PostForm("level"))
	//	tel,_ := strconv.Atoi(c.PostForm("tel"))
	//	var u = model.User{
	//		Name: name,
	//		Age:  age1,
	//		Sex:  sex,
	//		Department: department,
	//		Level: level,
	//		Tel: tel,
	//	}
	//	res := db.Create(&u)
	//	if res.Error != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{
	//			"msg": "插入失败",
	//		})
	//	} else {
	//		c.JSON(http.StatusOK, gin.H{
	//			"msg": "插入成功",
	//		})
	//	}
	//
	//})
	//r.GET("/user", func(c *gin.Context) {
	//	db, _ := gorm.Open(mysql.New(mysql.Config{
	//		DSN: "root:1997@tcp(127.0.0.1:3306)/db1?charset=utf8&parseTime=True&loc=Local"}), &gorm.Config{})
	//	db.AutoMigrate(&model.User{})
	//	name := c.Query("name")
	//	var users []model.User
	//	if name == "" {
	//		db.Find(&users)
	//		c.JSON(http.StatusOK, gin.H{
	//			"resutls": users,
	//		})
	//	} else {
	//		db.Where("name <> ?", name).Find(&users)
	//		c.JSON(http.StatusOK, gin.H{
	//			"results": users,
	//		})
	//
	//	}
	//})
	r := router.Router()
	r.Run(":5050")

}
