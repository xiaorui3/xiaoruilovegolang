package controllers

import (
	"awesomeProject/model"
	"awesomeProject/pkg/logger"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserController struct {
}

func (u *UserController) SelectUserString(c *gin.Context) {
	id := c.Param("id")
	idstr, _ := strconv.Atoi(id)
	user, _ := model.User{}.GetUserTest(idstr)
	ReturnSuccess(c, true, 200, "Success", user, 0)
	//return user
}

func (u UserController) SelectUserSuccess(c *gin.Context) {
	t := time.Now()
	p := t.Format("2006-01-02 15:04:05")
	ReturnSuccess(c, true, 200, p+" Success!", nil, 0)
}

func (u *UserController) SelectUserError(c *gin.Context) {
	ReturnError(c, false, 500, "Error", nil, 0)
}

func (u *UserController) SelectList(c *gin.Context) {
	logger.Write("日志信息", "user")
	ReturnSuccess(c, true, 200, "Success", nil, 0)
}

func (u *UserController) SelectLists(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			ReturnError(c, false, 5001, fmt.Sprintf("抓到了一个异常！: %v", err), nil, 0)
			fmt.Println("抓到了一个异常！:", err)
		}
	}()

	paramId := c.Param("id")
	postFormId := c.PostForm("id")
	jsonData := make(map[string]interface{})
	_ = c.BindJSON(&jsonData)

	var id int = -1

	if paramId != "" {
		parsed, err := strconv.Atoi(paramId)
		if err != nil {
			ReturnError(c, false, 400, "无效的ID参数", nil, 0)
			return
		}
		id = parsed
	} else if postFormId != "" {
		parsed, err := strconv.Atoi(postFormId)
		if err != nil {
			ReturnError(c, false, 400, "无效的ID参数", nil, 0)
			return
		}
		id = parsed
	} else if jsonData["id"] != nil {
		switch v := jsonData["id"].(type) {
		case float64:
			id = int(v)
		case string:
			parsed, err := strconv.Atoi(v)
			if err != nil {
				ReturnError(c, false, 400, "无效的ID参数", nil, 0)
				return
			}
			id = parsed
		}
	}

	table, err := model.User{}.GetUserTestToDataBaseTable(id)
	if err != nil {
		ReturnError(c, false, 500, "查询失败", nil, 0)
		return
	}
	ReturnSuccess(c, true, 200, "Success", table, 0)
}
