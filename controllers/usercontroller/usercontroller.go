package usercontroller

import (
	"encoding/json"
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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	mysql.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"user": user})

}
func Update(c *gin.Context) {
	var user models.User

	id := c.Param("id")
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if mysql.DB.Model(&user).Where("id = ?", id).Updates(&user).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "You cant update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfuly Update User"})

}
func Delete(c *gin.Context) {
	var user models.User

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if mysql.DB.Delete(&user, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "You cant remove this data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Succesfully Delete This Data"})
}
