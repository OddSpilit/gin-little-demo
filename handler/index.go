package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Index(c *gin.Context){
	CheckCookie(c)
	c.HTML(http.StatusOK, "index/index.html", gin.H{
		"title": "Admin",
	})
}

func Welcome(c *gin.Context){
	CheckCookie(c)
	username, _ := c.Cookie("username")
	c.HTML(http.StatusOK, "index/welcome.html", gin.H{
		"title": "Admin",
		"username": username,
		"time": time.Now().Format("2006-01-02 15:04:05.000 Mon Jan"),
	})
}
