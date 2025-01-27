package main

import (
	"github.com/mirhasalh/gin-gorm-demo/initializers"
	"github.com/mirhasalh/gin-gorm-demo/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}