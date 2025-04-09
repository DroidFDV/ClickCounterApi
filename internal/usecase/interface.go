package usecase

import "ClickCounterApi/internal/models"

type ClickProvider interface {
	IncrementClick(bannerID string) error
	GetStats(bannerID string, timeInterval models.StatRequest) (models.StatResponse, error)
}
