package repository

import (
	"ClickCounterApi/internal/models"
	"time"
)

type RepoProvider interface {
	IncrementClick(bannerID string, time time.Time) error
	GetStats(bannerID string, timeInterval models.StatRequest) ([]models.MinuteEntry, error)
}
