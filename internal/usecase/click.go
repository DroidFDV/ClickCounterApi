package usecase

import (
	"ClickCounterApi/internal/models"
	"ClickCounterApi/internal/repository"
	"time"

	"github.com/go-faster/errors"
)

type ClickUsecase struct {
	repository repository.RepoProvider
}

func NewClick(repository repository.RepoProvider) *ClickUsecase {
	return &ClickUsecase{
		repository: repository,
	}
}

// Функция для увеличения счетчика кликов
func (cu *ClickUsecase) IncrementClick(bannerID string) error {
	now := time.Now().Truncate(time.Minute)
	err := cu.repository.IncrementClick(bannerID, now)
	if err != nil {
		return errors.Wrap(err, "ClickUsecase IncrementClick")
	}
	return nil
}

// Функция для получения статистики
func (cu *ClickUsecase) GetStats(bannerID string, timeInterval models.StatRequest) (models.StatResponse, error) {
	stats, err := cu.repository.GetStats(bannerID, timeInterval)
	if err != nil {
		return models.StatResponse{Stats: nil}, errors.Wrap(err, "ClickUsecase GetStats")
	}

	response := models.StatResponse{
		Stats: make([]models.StatEntry, len(stats)),
	}

	for i, stat := range stats {
		response.Stats[i] = models.StatEntry{
			Ts: stat.Timestamp.Format(time.RFC3339),
			V:  stat.Count,
		}
	}
	return response, nil
}
