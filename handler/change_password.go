package handler

import (
	"admin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)
type ChangePassword struct {
	Password string `form:"password" json:"password"`
	SurePassword string `form:"sure_password" json:"sure_password"`
}

func Change(c *gin.Context){
	c.HTML(http.StatusOK, "change_password/index.html", gin.H{
		"title": "更改密码",
	})
}

func Handle(c *gin.Context){
	var  changePassword ChangePassword
	if err := c.ShouldBind(&changePassword); err == nil {
		if changePassword.Password != changePassword.SurePassword || changePassword.Password == ""{
			c.JSON(http.StatusOK, gin.H{
				"status": 1,
				"msg": "请保证输入的密码跟确认密码一致并且不为空!",
			})
			return
		}else  {
			id, err := c.Cookie("id")
			if err != nil{
				panic(err)
			}
			var  user models.HUser
			db.Table("h_user").First(&user, id)
			user.Password = MD5(changePassword.Password)
			db.Table("h_user").Save(&user)
			c.JSON(http.StatusOK, gin.H{
				"status": 0,
				"msg": "修改成功！",
			})
		}
	}
}
