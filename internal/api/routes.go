package router

import (
	"github.com/gin-gonic/gin"
	"plan_it/internal/api/create_item"
	"plan_it/internal/api/get_item"
)

func NewRouter(createItemHandler *create_item.CreateItemHandler, getItemsHandler *get_item.GetItemsHandler) *gin.Engine {
	router := gin.Default()

	// Создание элемента
	router.POST("/items", createItemHandler.Handle)

	// Получение элементов
	router.GET("/items", getItemsHandler.Handle)

	return router
}
