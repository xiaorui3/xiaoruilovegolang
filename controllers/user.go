package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserController struct {
}

func (u *UserController) TableName() string {
	return "users"
}

func (u *UserController) SelectUserString() string {
	return "user list"
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
	ReturnSuccess(c, true, 200, "Success", nil, 0)
}
