package main

import (
	"admin/handler"
	"github.com/gin-gonic/gin"
	_"github.com/jinzhu/gorm/dialects/mysql"
)

func main(){
	r := gin.Default()
	r.Static("/static/admin", "./static/admin")
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/login", func(c *gin.Context) {
		handler.LoginWay(c)
	})
	r.POST("/login_handle", func(c *gin.Context) {
		handler.LoginHandle(c)
	})
	r.GET("/index", func(c *gin.Context) {
		handler.Index(c)
	})
	r.GET("/welcome", func(c *gin.Context) {
		handler.Welcome(c)
	})
	r.GET("/logOut", func(c *gin.Context) {
		handler.LogOut(c)
	})
	changePassword := r.Group("/change_password")
	{
		changePassword.GET("/index", func(c *gin.Context) {
			handler.Change(c)
		})
		changePassword.PUT("/handle", func(c *gin.Context) {
			handler.Handle(c)
		})
	}
	order := r.Group("/order")
	{
		order.GET("/index", func(c *gin.Context) {
			handler.OrderIndex(c)
		})
		order.GET("/search", func(c *gin.Context) {
			handler.Search(c)
		})
		order.GET("/add", func(c *gin.Context) {
			handler.Add(c)
		})
		order.POST("/addHandle", func(c *gin.Context) {
			handler.AddHandle(c)
		})
		order.POST("/delHandle", func(c *gin.Context) {
			handler.DelHandle(c)
		})
	}
	r.Run()
}

func init() {
	handler.InitDB()
}