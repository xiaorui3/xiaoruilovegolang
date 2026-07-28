package router

import (
	"awesomeProject/controllers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	user := r.Group("/user")
	{
		user.GET("/user/SelectList", (&controllers.UserController{}).SelectList)
		user.GET("/user/SelectUserIdToName/:id", (&controllers.UserController{}).SelectUserIdToName)
		user.GET("/user/SelectLists", (&controllers.UserController{}).SelectLists)
		user.GET("/user/info", controllers.UserController{}.SelectUserSuccess)
		user.GET("/hello", func(c *gin.Context) {
			//c.String(200, "Hello World")
			controllers.ReturnSuccess(c, true, 200, "Success", nil, 0)
			// fmt.Println("Get")
		})
		user.POST("/user/list", func(context *gin.Context) {
			context.String(200, "User list")

			fmt.Println("Post")
			fmt.Println(context)
		})
		user.GET("/user/get/:id", (&controllers.UserController{}).SelectUserString)
		user.DELETE("/user/delete/{id}", func(context *gin.Context) {
			context.String(200, "User delete"+context.Param("id"))
			fmt.Println("Delete")
		})

		user.PUT("/user/update/{id}", func(context *gin.Context) {
			context.String(200, "User update"+context.Param("id"))
			fmt.Println("Update")
		})
	}

	order := r.Group("/order")
	{
		order.GET("/list/info", (&controllers.OrderController{}).SelectList)
		order.POST("/list/form", (&controllers.OrderController{}).SelectListFrom)
		order.POST("/list/panic", (&controllers.OrderController{}).TestPanic)
	}

	return r
}
