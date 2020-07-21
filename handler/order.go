package handler

import (
	"admin/models"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
	"time"
)
var (
	totalType = [...]string{"门诊","体检","疫苗","团体"}
)
type Delete struct {
	Id int `form:"id" json:"id"`
}
func OrderIndex(c *gin.Context){
	CheckCookie(c)
	var order[] models.HOrder
	var count int
	db.Table("h_order").Count(&count)
	pages := int(math.Ceil(float64(count)/5.000))
	params := make(map[string]int, 1)
	params["type"] = -1
	pagesHandle := make(map[string]interface{}, 10)
	pagesHandle["lastPage"] = 1
	pagesHandle["totalPage"] = pages
	pagesHandle["currentPage"] = 1
	pagesHandle["firstPage"] = 1
	if pages > 1 {
		pagesHandle["nextPage"] = 2
	}else  {
		pagesHandle["nextPage"] = 1
	}
	db.Table("h_order").Limit(5).Find(&order)
	order = typeMark(order)
	c.HTML(http.StatusOK, "order/index.html", gin.H{
		"data": order,
		"title": "数据列表",
		"count": count,
		"totalType": totalType,
		"params": params,
		"pagesHandle": pagesHandle,
	})
}

func Search(c *gin.Context){
	CheckCookie(c)
	var count int
	params := make(map[string]interface{}, 10)
	start := c.Query("start")
	end := c.Query("end")
	params["startTime"] = start
	params["endTime"] = end
	currentPage := c.DefaultQuery("page", "1")
	offsetIndex, _ := strconv.Atoi(currentPage)
	offset := (offsetIndex - 1) * 5
	if currentPage == "" {
		if start >= end {
			c.HTML(http.StatusBadRequest, "error/throw.html", gin.H{
				"status": 1,
				"message": "开始时间不得大于结束时间!",
			})
			return
		}
	}
	var startTime int64
	var endTime int64
	timeLayout := "2006-01-02"  //转化所需模板
	loc, _ := time.LoadLocation("Local")    //获取时区
	if start != "" {
		tmp, _ := time.ParseInLocation(timeLayout, start, loc)
		startTime = tmp.Unix()    //转化为时间戳 类型是int64
	}else  {
		startTime = 0
	}
	if end != "" {
		tmp, _ := time.ParseInLocation(timeLayout, end, loc)
		endTime = tmp.Unix()    //转化为时间戳 类型是int64
	}else  {
		endTime = time.Now().Unix()
	}
	var  order[] models.HOrder
	Db := db
	params["num"] = ""
	num := c.DefaultQuery("num", "")
	if  num != "" && num != "0"{
		numInt, _ := strconv.Atoi(num)
		Db = Db.Where("num = ?", numInt)
		params["num"] = numInt
	}
	params["type"] = -1
	numType := c.DefaultQuery("type","-1")
	typeInt, _ := strconv.Atoi(numType)
	if typeInt != -1{
		Db = Db.Where("type = ?", typeInt)
		params["type"] = typeInt
	}

	Db = Db.Where("complete_time BETWEEN ? AND ?", startTime, endTime)
	Db.Find(&order).Count(&count)
	Db.Offset(offset).Limit(5).Find(&order)
	order = typeMark(order)
	pages := int(math.Ceil(float64(count)/5.000))
	getPagesHandle := make(map[string]interface{},10)
	pagesHandle := pageHandle(getPagesHandle, pages, currentPage)
	//res := Paginator(, "",  pages)
	//params["paginator"] = res
	c.HTML(http.StatusOK, "order/index.html", gin.H{
		"data": order,
		"title": "数据列表",
		"count": count,
		"totalType": totalType,
		"params": params,
		"pagesHandle": pagesHandle,
	})
}

func Add(c *gin.Context){
	c.HTML(http.StatusOK, "order/add.html", gin.H{
		"title": "添加",
	})
}

func AddHandle(c *gin.Context){
	num := c.PostForm("num")
	idCard := c.PostForm("idCard")
	orderNo := c.PostForm("orderNo")
	rcptNo := c.PostForm("rcptNo")
	transNo := c.PostForm("transNo")
	bill := c.PostForm("bill")
	getType := c.PostForm("type")
	order := models.HOrder{Num:num,IdCard:idCard,OrderNo:orderNo,RcptNo:rcptNo,TransNo:transNo,CreateTime:time.Now().Unix(),ApplyTime:time.Now().Unix(),RefundTime:time.Now().Unix(),
		CompleteTime:time.Now().Unix(),Type:getType,Openid:"123213",Status:"0", Bill:bill, IfCheck:"0"}
	if err := db.Create(&order).Error; err != nil{
		c.JSON(http.StatusOK, gin.H{
			"status": 1,
			"msg": err,
		})
	}else {
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg": "创建成功!",
		})
	}
}

func DelHandle(c *gin.Context){
	var delete Delete
	if err := c.Bind(&delete); err == nil {
		var order models.HOrder
		if err := db.Where("id=?", delete.Id).Delete(&order).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": 1,
				"msg":    err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": 0,
				"msg":    "删除成功!",
			})
		}
	}
}

func typeMark(orderChange[] models.HOrder) (order[] models.HOrder){
	for key, v := range orderChange{
		switch v.Type {
		case "0":
			orderChange[key].Type = "门诊"
			break
		case "1":
			orderChange[key].Type = "体检"
			break
		case "2":
			orderChange[key].Type = "疫苗"
			break
		case "3":
			orderChange[key].Type = "团体"
			break
		}
	}
	order = orderChange
	return
}