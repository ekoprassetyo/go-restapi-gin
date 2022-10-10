package usercontroller

import (
	"go-restapi-gin/models"
	"go-restapi-gin/pkg/mysql"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var user []models.User

	mysql.DB.Find(&user)

	c.JSON(http.StatusOK, gin.H{"user": user})

}
func Show(c *gin.Context) {
	var user models.User

	id := c.Param("id")

	if err := mysql.DB.First(&user, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return

		}
	}

	c.JSON(http.StatusOK, gin.H{"user": user})

}
func Create(c *gin.Context) {

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	mysql.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"user": user})

}
func Update(c *gin.Context) {

}
func Delete(c *gin.Context) {

}
