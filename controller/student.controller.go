package controller

import (
	"http_request/class-management/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateStudent creates a new student
func (ctrl *Controller) CreateStudent(c *gin.Context) {
	var student models.Student

	if err := c.BindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := ctrl.DatabaseService.CreateStudent(&student)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create student"})
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetStudent retrieves a student by ID
func (ctrl *Controller) GetStudent(c *gin.Context) {
	id := c.Param("id")

	student, err := ctrl.DatabaseService.GetStudentByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve student"})
		return
	} else if student == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	c.JSON(http.StatusOK, student)
}

// UpdateStudent updates an existing student
func (ctrl *Controller) UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	var student models.Student

	if err := c.BindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.DatabaseService.UpdateStudent(id, &student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update student"})
		return
	}

	c.JSON(http.StatusOK, student)
}

// DeleteStudent deletes a student by ID
func (ctrl *Controller) DeleteStudent(c *gin.Context) {
	id := c.Param("id")

	if err := ctrl.DatabaseService.DeleteStudent(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete student"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Student deleted successfully"})
}
