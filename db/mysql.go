package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
type HUser struct{
	gorm.Model
	Id int `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
	Mobile string `json:"mobile"`
	Last_login_time int `json:"last_login_time"`
	Status int `json:"status"`
	Create_time int `json:"create_time"`
	Update_time int `json:"update_time"`
}




