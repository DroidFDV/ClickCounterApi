package main

import (
	"ClickCounterApi/internal/app"
	"ClickCounterApi/internal/handler"
	"ClickCounterApi/internal/repository"
	"ClickCounterApi/internal/storage"
	"ClickCounterApi/internal/usecase"
	"log"

	"github.com/pkg/errors"
)

func main() {
	connStr := "postgres://postgres:postgres@postgres:5432/postgres?sslmode=disable"
	pool, err := storage.GetConnect(connStr)
	if err != nil {
		log.Fatal(errors.Wrap(err, "GetConnect: Failed to connect db"))
	}
	defer pool.Close()

	clickRepository := repository.New(pool)
	clickProvider := usecase.NewClick(clickRepository)
	handler := handler.New(clickProvider)
	router := app.NewRouter(handler)

	router.Run(":8080")
}
