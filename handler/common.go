package handler

import (
	"admin/models"
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

var (
	db *gorm.DB
)


func InitDB() (err error){
	dsn := "root:123456@tcp(127.0.0.1)/test?charset=utf8mb4&parseTime=True"
	db, err = gorm.Open("mysql", dsn)
	if  err != nil{
		return
	}
	var order models.HOrder
	var user models.HUser
	db.AutoMigrate(&order)
	db.AutoMigrate(&user)
	db.LogMode(true)
	return
}
func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

func CheckCookie(c *gin.Context) {
	_, err := c.Cookie("username")
	if err != nil{
		c.Redirect(303, "/login")
	}
}

func debug(c *gin.Context, msg interface{}){
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
	return
}

func pageHandle(pagesHandle map[string]interface{}, pages int, currentPage string) (handle map[string]interface{}){
	if pages > 0 {
		pagesHandle["totalPage"] = pages
		pagesHandle["firstPage"] = 1
		pagesHandle["currentPage"], _ = strconv.Atoi(currentPage)
		if currentPage == "1"{
			pagesHandle["lastPage"] = 1
		}else  {
			pagesHandle["lastPage"] = pagesHandle["currentPage"].(int)-1
		}
		if pages > pagesHandle["currentPage"].(int) {
			pagesHandle["nextPage"] = pagesHandle["currentPage"].(int) + 1
		}else {
			pagesHandle["nextPage"] = pagesHandle["currentPage"].(int)
		}
	}
	handle = pagesHandle
	return
}

