package main

import (
	"desktop/go-project/initializers"
	"desktop/go-project/models"
)


func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}

// go run migrate/migrate.go   