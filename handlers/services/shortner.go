package services

import (
	"fmt"
	"net/http"
	"url-shortner/db/models"
	"url-shortner/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterShortnerHandler(router *gin.Engine, db *gorm.DB) {
	router.POST("/shorten", func(c *gin.Context) {
		HandleUrlShorten(c, db)
	})

	router.GET("/:shortCode", func(c *gin.Context) {
		HandleShortUrlRedirect(c, db)
	})
}

func HandleUrlShorten(c *gin.Context, db *gorm.DB) {
	var urlInfo models.Urls

	longUrl := c.PostForm("longUrl")

	urlInfo.LongUrl = longUrl
	urlInfo.ShortCode = utils.GenerateShortCode()

	if err := db.Create(&urlInfo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Unable to create shortened url for '%v'. Error: %v", longUrl, err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"short_url": "http://localhost:8080/" + urlInfo.ShortCode,
	})
}

func HandleShortUrlRedirect(c *gin.Context, db *gorm.DB) {
	var urlInfo models.Urls

	shortCode := c.Param("shortCode")

	if len(shortCode) != 5 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Invalid url code given.",
		})
		return
	}

	if err := db.Where("short_code = ?", shortCode).First(&urlInfo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Unable to find url for code '%v'.", shortCode),
		})
		return
	}

	c.Redirect(http.StatusMovedPermanently, urlInfo.LongUrl)
}
