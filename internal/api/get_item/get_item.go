package get_item

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetItemsHandler struct {
	service getitems
}

func NewGetItemsHandler(service getitems) *GetItemsHandler {
	return &GetItemsHandler{service: service}
}

func (h *GetItemsHandler) Handle(c *gin.Context) {
	items, err := h.service.GetItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch items"})
		return
	}
	c.JSON(http.StatusOK, items)
}
