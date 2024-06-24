package handler

import (
	"net/http"

	"github.com/Go11Group/Hojiakbar-Abbosov/Lesson43/metro_service/models"
	"github.com/gin-gonic/gin"
)

func (h *TerminalHandler) CreateTerminal(c *gin.Context) {
	var req models.CreateTerminal
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Repo.Create(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Terminal created successfully"})
}

func (h *TerminalHandler) GetTerminalById(c *gin.Context) {
	id := c.Param("id")
	terminal, err := h.Repo.GetById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, terminal)
}

func (h *TerminalHandler) UpdateTerminal(c *gin.Context) {
	var req models.Terminal
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Repo.Update(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Terminal updated successfully"})
}

func (h *TerminalHandler) DeleteTerminal(c *gin.Context) {
	id := c.Param("id")
	if err := h.Repo.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Terminal deleted successfully"})
}
