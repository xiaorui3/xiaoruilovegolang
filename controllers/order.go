package controllers

import "github.com/gin-gonic/gin"

type OrderController struct {
}

func (o OrderController) SelectList(c *gin.Context) {
	ReturnSuccess(c, true, 200, "Success", nil, 0)
}
