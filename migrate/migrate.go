package main

import (
	"JSON_CRUD_API_With_Gin_and_GORM/Initializers"
	"JSON_CRUD_API_With_Gin_and_GORM/models"
)

func init() {
	Initializers.LoadEnvVariables()
	Initializers.ConnectToDB()
}

func main() {
	Initializers.DB.AutoMigrate(&models.Post{})
}
