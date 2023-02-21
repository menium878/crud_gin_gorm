package main

import (
	"crud_gin_gorm/initializers"
	"crud_gin_gorm/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
