package controllers

import (
	"crud_gin_gorm/initializers"
	"crud_gin_gorm/models"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {

	var body struct {
		Body  string `json:"body"`
		Title string `json:"title"`
	}

	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostRead(c *gin.Context) {
	var posts []models.Post
	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"post": posts,
	})

}
