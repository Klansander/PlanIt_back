package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	router "plan_it/internal/api"
	api_create_item "plan_it/internal/api/create_item"
	api_get_item "plan_it/internal/api/get_item"
	"plan_it/internal/service/create_item"
	"plan_it/internal/storage"

	"go.uber.org/zap"

	"plan_it/internal/service/get_item"
)

const (
	success = 0
	fail    = 1
)

func main() {
	os.Exit(run())
}

func run() int {
	var err error

	// Инициализация логгера
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Println("Error initializing logger:", err)
		return fail
	}
	defer logger.Sync()

	// Контекст с обработкой сигнала
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// Инициализация репозитория
	repo := storage.NewRepository()

	// Инициализация сервисов
	createItemService := create_item.NewCreateItemService(repo)
	getItemsService := get_item.NewGetItemsService(repo)

	// Инициализация хендлеров
	createItemHandler := api_create_item.NewCreateItemHandler(createItemService)
	getItemsHandler := api_get_item.NewGetItemsHandler(getItemsService)

	// Инициализация роутера
	router := router.NewRouter(createItemHandler, getItemsHandler)

	// Запуск HTTP сервера
	go func() {
		if err := router.Run(":8080"); err != nil {
			logger.Fatal("Error starting server", zap.Error(err))
		}
	}()

	// Ожидание сигнала для завершения
	<-ctx.Done()

	logger.Info("Shutting down gracefully")
	return success
}
