package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Controller to remove an entry from storage
func (h *Handler) DeleteData(c *gin.Context) {
	key := c.Param("key")

	err := h.librapi.DeleteDataFromStorage(key)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// Controller to read an entry in storage
func (h *Handler) GetData(c *gin.Context) {
	key := c.Param("key")

	data, err := h.librapi.GetDataFromStorage(key)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

// Controller for liveness and readiness probe
func (h *Handler) Healthcheck(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusOK)
}

type createEntry struct {
	Key   string `json:"key" binding:"required"`
	Value string `json:"value" binding:"required"`
}

// Controller to save a new entry in storage
func (h *Handler) PostData(c *gin.Context) {

	var query createEntry

	if err := c.BindJSON(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.librapi.SaveDataInStorage(query.Key, query.Value)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

type updateEntry struct {
	Value string `json:"value" binding:"required"`
}

// Cotnroller to edit an entry in storage
func (h *Handler) PutData(c *gin.Context) {
	key := c.Param("key")

	var query updateEntry

	if err := c.BindJSON(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.librapi.SaveDataInStorage(key, query.Value)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
