package services

import (
	"fmt"
	"net/http"
	"url-shortner/db/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterDiscoveryRoutes(router *gin.Engine, db *gorm.DB) {
	router.GET("/codes/discovery", func(c *gin.Context) {
		HandleGetAllCodes(c, db)
	})
}

func HandleGetAllCodes(c *gin.Context, db *gorm.DB) {
	var urlCodes []models.Urls

	if err := db.Find(&urlCodes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Unable to retrieve all the codes. Error: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, urlCodes)
}
