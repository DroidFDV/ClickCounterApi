package app

import (
	"ClickCounterApi/internal/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(handler *handler.Handle) *gin.Engine {
	router := gin.Default()

	router.GET("/counter/:bannerID", handler.IncrementClick)

	router.POST("/stats/:bannerID", handler.GetStats)

	return router
}
