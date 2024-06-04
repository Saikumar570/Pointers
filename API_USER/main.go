package main

import (
	"api/controllers"
	"api/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.Create)
	r.GET("/posts", controllers.ReadAll)
	r.GET("/posts/:id", controllers.ReadOne)
	r.PUT("/posts/:id", controllers.Update)
	r.DELETE("/posts/:id", controllers.Deletes)

	r.Run() // listen and serve
}
