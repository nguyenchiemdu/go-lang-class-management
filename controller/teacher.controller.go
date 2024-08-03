package controller

import (
	"http_request/class-management/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateTeacher creates a new teacher
func (ctrl *Controller) CreateTeacher(c *gin.Context) {
	var teacher models.Teacher

	if err := c.BindJSON(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := ctrl.DatabaseService.CreateTeacher(&teacher)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create teacher"})
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetTeacher retrieves a teacher by ID
func (ctrl *Controller) GetTeacher(c *gin.Context) {
	id := c.Param("id")

	teacher, err := ctrl.DatabaseService.GetTeacherByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve teacher"})
		return
	} else if teacher == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Teacher not found"})
		return
	}

	c.JSON(http.StatusOK, teacher)
}

// UpdateTeacher updates an existing teacher
func (ctrl *Controller) UpdateTeacher(c *gin.Context) {
	id := c.Param("id")
	var teacher models.Teacher

	if err := c.BindJSON(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.DatabaseService.UpdateTeacher(id, &teacher); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update teacher"})
		return
	}

	c.JSON(http.StatusOK, teacher)
}

// DeleteTeacher deletes a teacher by ID
func (ctrl *Controller) DeleteTeacher(c *gin.Context) {
	id := c.Param("id")

	if err := ctrl.DatabaseService.DeleteTeacher(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete teacher"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Teacher deleted successfully"})
}
