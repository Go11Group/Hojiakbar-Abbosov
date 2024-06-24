package handler

import (
	"net/http"

	"github.com/Go11Group/Hojiakbar-Abbosov/Lesson43/metro_service/models"
	"github.com/gin-gonic/gin"
)

func (h *handler) CreateStation(c *gin.Context) {
	var req models.CreateStation
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.Station.Create(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Station created successfully"})
}

func (h *handler) GetStationById(c *gin.Context) {
	id := c.Param("id")
	station, err := h.Station.GetById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, station)
}

func (h *handler) UpdateStation(c *gin.Context) {
	var req models.UpdateStation
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.Station.Update(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Station updated successfully"})
}

func (h *handler) DeleteStation(c *gin.Context) {
	id := c.Param("id")
	err := h.Station.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Station deleted successfully"})
}
