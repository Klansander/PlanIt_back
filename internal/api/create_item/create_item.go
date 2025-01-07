package create_item

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateItemHandler struct {
	service createItem
}

func NewCreateItemHandler(service createItem) *CreateItemHandler {
	return &CreateItemHandler{service: service}
}

func (h *CreateItemHandler) Handle(c *gin.Context) {
	var item int
	createdItem, err := h.service.CreateItem(item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create item"})
		return
	}

	c.JSON(http.StatusCreated, createdItem)
}
