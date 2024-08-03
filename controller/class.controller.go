package controller

import (
	"http_request/class-management/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateClass creates a new class with no students and teacher
func (ctrl *Controller) CreateClass(c *gin.Context) {
	var class models.Class

	if err := c.BindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.DatabaseService.CreateClass(&class); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create class"})
		return
	}

	c.JSON(http.StatusOK, class)
}

// UpdateClassTeacher updates the teacher of a class
func (ctrl *Controller) UpdateClassTeacher(c *gin.Context) {
	classID := c.Param("id")
	var request struct {
		TeacherID string `json:"teacher_id" binding:"required"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.DatabaseService.UpdateClassTeacher(classID, request.TeacherID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update class teacher"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Teacher updated successfully"})
}

// AddStudentToClass adds a student to a class
func (ctrl *Controller) AddStudentToClass(c *gin.Context) {
	classID := c.Param("id")
	var request struct {
		StudentID string `json:"student_id" binding:"required"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.DatabaseService.AddStudentToClass(classID, request.StudentID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add student to class"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Student added successfully"})
}

// RemoveStudentFromClass removes a student from a class
func (ctrl *Controller) RemoveStudentFromClass(c *gin.Context) {
	classID := c.Param("id")
	var request struct {
		StudentID string `json:"student_id" binding:"required"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.DatabaseService.RemoveStudentFromClass(classID, request.StudentID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove student from class"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Student removed successfully"})
}

// GetClass retrieves a class by its ID
func (ctrl *Controller) GetClass(c *gin.Context) {
	classID := c.Param("id")

	class, err := ctrl.DatabaseService.GetClassByID(classID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get class"})
		return
	}

	if class == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Class not found"})
		return
	}

	c.JSON(http.StatusOK, class)
}

// DeleteClass deletes a class by its ID
func (ctrl *Controller) DeleteClass(c *gin.Context) {
	classID := c.Param("id")

	if err := ctrl.DatabaseService.DeleteClassByID(classID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete class"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Class deleted successfully"})
}
