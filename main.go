package main

import (
	"awesomeProject/router"
)

func main() {

	r := router.Router()

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
