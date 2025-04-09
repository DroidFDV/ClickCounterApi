package handler

import (
	"ClickCounterApi/internal/models"
	"ClickCounterApi/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handle struct {
	clickProvider usecase.ClickProvider
}

func New(provider usecase.ClickProvider) *Handle {
	return &Handle{
		clickProvider: provider,
	}
}

func (h *Handle) IncrementClick(c *gin.Context) {
	bannerID := c.Param("bannerID")
	err := h.clickProvider.IncrementClick(bannerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Click counted"})
}

func (h *Handle) GetStats(c *gin.Context) {
	bannerID := c.Param("bannerID")
	var req models.StatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	stats, err := h.clickProvider.GetStats(bannerID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stats)
}
