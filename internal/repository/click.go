package repository

import (
	"ClickCounterApi/internal/models"
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

type ClickRepository struct {
	pool *pgxpool.Pool
}

func New(conn *pgxpool.Pool) *ClickRepository {
	repo := &ClickRepository{
		pool: conn,
	}
	repo.initTable()
	return repo
}

// TODO: migrations
func (cr *ClickRepository) initTable() error {
	// SQL-запрос для создания таблицы clicks
	query := `
			CREATE TABLE IF NOT EXISTS clicks (
				timestamp TIMESTAMP NOT NULL,
				banner_id INT NOT NULL,
				count INT NOT NULL DEFAULT 1,
				PRIMARY KEY (timestamp, banner_id)
			);
		`
	// Выполнение запроса
	_, err := cr.pool.Exec(context.Background(), query)
	return errors.Wrap(err, "ClickRepository initTable")
}

func (cr *ClickRepository) IncrementClick(bannerID string, time time.Time) error {
	query := `
        INSERT INTO clicks (timestamp, banner_id, count)
        VALUES ($1, $2, 1)
        ON CONFLICT (timestamp, banner_id) DO UPDATE
        SET count = clicks.count + 1;
    `
	_, err := cr.pool.Exec(context.Background(), query, time, bannerID)

	return errors.Wrap(err, "IncrementClick pool.Exec")
}

func (cr *ClickRepository) GetStats(bannerID string, timeInterval models.StatRequest) ([]models.MinuteEntry, error) {
	query := `
        SELECT 
			DATE_TRUNC('minute', timestamp) AS minute, 
			SUM(count) AS total_clicks
		FROM clicks
		WHERE banner_id = $1 AND timestamp >= $2 AND timestamp <= $3
		GROUP BY minute
		ORDER BY minute ASC;
    `
	rows, err := cr.pool.Query(context.Background(), query, bannerID, timeInterval.TsFrom, timeInterval.TsTo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	stats := make([]models.MinuteEntry, 0)
	for rows.Next() {
		var clickStat models.MinuteEntry
		if err := rows.Scan(&clickStat.Timestamp, &clickStat.Count); err != nil {
			return nil, errors.Wrap(err, "GetStats rows.Scan")
		}
		stats = append(stats, clickStat)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "GetStats Rows.Err")
	}

	return stats, nil
}
