package handler

import (
	"admin/models"
	"strconv"

	// "fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
type Login struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}
func LoginWay(c *gin.Context){
		if _, err := c.Cookie("username"); err == nil{
			c.Redirect(303, "/index")
		}
		c.HTML(http.StatusOK, "login/login.html", gin.H{
			"title": "Admin",
		})
	return
}

func LoginHandle(c *gin.Context){
	var user models.HUser
	var  login Login
	if err := c.ShouldBind(&login); err == nil {
		db.Table("h_user").Where("password=? AND username=?", MD5(login.Password), login.Username).Find(&user)
		if user.Id == 0 {
			c.JSON(http.StatusAccepted, gin.H{
				"status": 1,
				"msg": "账号或者密码错误!",
			})
		}else  {
			c.SetCookie("id", strconv.Itoa(user.Id), 3600, "/", "localhost", false, true)
			c.SetCookie("username", user.Username, 3600, "/", "localhost", false, true)
			c.JSON(http.StatusAccepted, gin.H{
				"status": 0,
				"msg": "登录成功！",
			})
			return
		}
	}
}

func LogOut(c *gin.Context){
	username, _ := c.Cookie("username")
	id, _ := c.Cookie("id")
	c.SetCookie("id", id, -1, "/", "localhost", false, true)
	c.SetCookie("username", username, -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg": "退出成功!",
	})
}

