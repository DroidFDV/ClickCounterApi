package models

import "time"

type StatRequest struct {
	TsFrom time.Time `json:"from"`
	TsTo   time.Time `json:"to"`
}

type MinuteEntry struct {
	Timestamp time.Time
	Count     int
}

type StatEntry struct {
	Ts string `json:"ts"` // Временная метка в формате ISO 8601
	V  int    `json:"v"`  // Количество кликов
}

type StatResponse struct {
	Stats []StatEntry `json:"stats"`
}
