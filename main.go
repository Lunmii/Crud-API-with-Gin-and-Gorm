package main

import (
	"JSON_CRUD_API_With_Gin_and_GORM/Initializers"
	"JSON_CRUD_API_With_Gin_and_GORM/controllers"
	"github.com/gin-gonic/gin"
)

func init() {
	Initializers.LoadEnvVariables()
	Initializers.ConnectToDB()
}
func main() {
	r := gin.Default()
	r.POST("posts", controllers.PostsCreate)
	r.Run()
}
