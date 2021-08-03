package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID        int
	Name      string
	Age       int
	Sex       string
	Department string
	Level int
	Tel int

}

var DB *gorm.DB

func InitDB()  {
	dsn := "root:1997@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	DB ,_= gorm.Open(mysql.Open(dsn),&gorm.Config{})


}
func FindAll() (UserList []*User,err error) {
	if err = DB.Find(&UserList).Error; err != nil {
		return nil,err
	}
	return

}
func FindOne(id int) (User *User, err error) {
	DB.First(&User,id)
	return
}
func Delete(id int) (User *User) {
	 DB.Where("id=?",id).Delete(&User)
	 return
}


