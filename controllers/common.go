package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type JSONStruct struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
	Count   int64       `json:"count"`
}

func ReturnSuccess(c *gin.Context, Success bool, Code int, message interface{}, data interface{}, count int64) {
	json := &JSONStruct{Success: Success, Code: Code, Message: message, Data: data, Count: count}
	fmt.Println(json)
	fmt.Println(*json)
	fmt.Println(&json)
	fmt.Println(&*json)
	fmt.Println(&json.Data)
	fmt.Println(*&json)
	c.JSON(200, *json)
}

func ReturnError(c *gin.Context, Success bool, Code int, message interface{}, data interface{}, count int64) {
	json := &JSONStruct{Success: Success, Code: Code, Message: message, Data: data, Count: count}
	c.JSON(500, *json)
}
