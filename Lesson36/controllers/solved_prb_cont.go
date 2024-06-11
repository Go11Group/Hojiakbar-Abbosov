package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go-gin-crud/database"
	"go-gin-crud/models"
)

func GetSolvedProblems(c *gin.Context) {
	solvedProblems, err := models.GetSolvedProblems(database.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, solvedProblems)
}

func GetSolvedProblem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	solvedProblem, err := models.GetSolvedProblem(database.DB, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Solved problem not found"})
		return
	}
	c.JSON(http.StatusOK, solvedProblem)
}

func CreateSolvedProblem(c *gin.Context) {
	var newSolvedProblem models.SolvedProblem
	if err := c.ShouldBindJSON(&newSolvedProblem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := models.CreateSolvedProblem(database.DB, &newSolvedProblem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newSolvedProblem)
}

func UpdateSolvedProblem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedSolvedProblem models.SolvedProblem
	if err := c.ShouldBindJSON(&updatedSolvedProblem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	updatedSolvedProblem.ID = id

	if err := models.UpdateSolvedProblem(database.DB, &updatedSolvedProblem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedSolvedProblem)
}

func DeleteSolvedProblem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := models.DeleteSolvedProblem(database.DB, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Solved problem deleted"})
}
