package controllers

import (
	"app/app/database"
	"app/app/models"
	"log"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostsStateDTO struct {
	State string `gorm:"type:varchar(2);not null" binding:"required"`
}

func UpdatePost(c *gin.Context) {
	var updatePost models.Posts
	var post models.Posts

	err := c.ShouldBindJSON(&updatePost)
	if err != nil {
		c.JSON(400, gin.H{
			"err": "invalid ",
		})
		c.Abort()
	}
	post.UpdatePost(&updatePost, int(updatePost.ID))
	c.JSON(200, gin.H{
		"işlem başarılı": "işlem başarılı",
	})
}

func UpdatePostState(c *gin.Context) {
	var updatePost models.Posts
	var post models.Posts
	err := c.ShouldBindJSON(&updatePost.State)
	if err != nil {
		c.JSON(400, gin.H{
			"err": "invalid ",
		})
		c.Abort()
	}
	post.UpdatePost(&updatePost, int(updatePost.ID))
	c.JSON(200, gin.H{
		"işlem başarılı": "işlem başarılı",
	})

}
func CreatePost(c *gin.Context) {
	var post models.Posts
	err := c.ShouldBindJSON(&post)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"err": "ivalid post require",
		})

		c.Abort()
		return
	}
	err = post.CreatePost()
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"err": "Error Creating Post	",
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"Message": "Sucessfully created cpost",
	})

}

func GetAllPost(c *gin.Context) {
	pageStr := c.Query("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, _ := strconv.Atoi(pageStr)
	limit := 5
	offsets := (page - 1) * limit
	var total int64

	var blogpost []models.Posts
	db := database.GlobalDB

	db.Preload("Posts.User").Offset(offsets).Limit(limit).Find(&blogpost)
	db.Model(&models.Posts{}).Count(&total)

	c.JSON(200, gin.H{
		"data": blogpost,
		"meta": gin.H{
			"total":     total,
			"page ":     page,
			"last_page": math.Ceil(float64(total) / float64(limit)),
		},
	})

}

func Profile(c *gin.Context) {

	var user models.User

	email, _ := c.Get("email")

	result := database.GlobalDB.Where("email = ?", email.(string)).First(&user)

	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(404, gin.H{
			"Error": "User Not Found",
		})
		c.Abort()
		return
	}

	if result.Error != nil {
		c.JSON(500, gin.H{
			"Error": "Could Not Get User Profile",
		})
		c.Abort()
		return
	}

	user.Password = ""

	c.JSON(200, user)
}
