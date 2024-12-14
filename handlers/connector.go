package handlers

import (
	"url-shortner/handlers/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	services.RegisterShortnerHandler(router, db)
	services.RegisterDiscoveryRoutes(router, db)
}
