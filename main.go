package main

import (
	"JSON_CRUD_API_With_Gin_and_GORM/Initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	Initializers.LoadEnvVariables()
}
func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
