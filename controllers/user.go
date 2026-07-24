package controllers

import (
	"awesomeProject/model"
	"awesomeProject/pkg/logger"
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
