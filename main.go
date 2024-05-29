package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lokeshllkumar/pswd-extension/handlers/endpoints"
)

func main() {
	r := gin.Default()

	r.GET("/pswd", endpoints.GetPasswordHandler)
	r.POST("/pswd", endpoints.AddPasswordHandler)

	if err := r.Run(":8080"); err != nil {
		fmt.Printf("Failed to run server: %v\n", err)
	}
}
