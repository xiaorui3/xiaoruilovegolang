package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
}

func (o OrderController) SelectList(c *gin.Context) {
	ReturnSuccess(c, true, 200, "Success", nil, 0)
}

func (o OrderController) SelectListFrom(c *gin.Context) {
	name := c.PostForm("name")
	json := make(map[string]interface{})
	err := c.BindJSON(&json)
	data := make([]interface{}, 0)
	data = append(data, name)
	if err == nil {
		// nameJson := json["name"].(string)
		data = append(data, json)

	}
	ReturnSuccess(c, true, 200, "Success", data, 0)
}

// 实现异常 panic defer recover
func (o OrderController) TestPanic(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			ReturnError(c, false, 5001, fmt.Sprintf("抓到了一个异常！: %v", err), nil, 0)
			fmt.Println("抓到了一个异常！:", err)
		}
	}()
	//panic("test panic") // 注释掉这行代码，程序就不会报错了

	num1 := 1
	num2 := 0
	num3 := num1 / num2

	ReturnSuccess(c, true, 200, "Success", num3, 0)
}
