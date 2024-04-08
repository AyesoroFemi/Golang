package main

import (
	"desktop/go-projectt/controllers"
	"desktop/go-projectt/initializers"
	"desktop/go-projectt/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}


func main() {
	fmt.Println("Hello world")
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.Run()
}

//  go mod init desktop/go-project     