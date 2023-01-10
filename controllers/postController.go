package controllers

import (
	"JSON_CRUD_API_With_Gin_and_GORM/Initializers"
	"JSON_CRUD_API_With_Gin_and_GORM/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	post := models.Post{Title: "Hello", Body: "Post body"}

	result := Initializers.DB.Create(&post)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"Post": post,
	})
}
