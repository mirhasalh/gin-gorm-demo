package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mirhasalh/gin-gorm-demo/controllers"
	"github.com/mirhasalh/gin-gorm-demo/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.GET("/posts", controllers.PostsCreate)

	r.Run()
}
