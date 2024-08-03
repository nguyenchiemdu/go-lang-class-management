package main

import (
	"http_request/class-management/config"
	"http_request/class-management/controller"
	"http_request/class-management/middleware"

	"github.com/gin-gonic/gin"
)

func main() {

	appConfig := config.LoadAppConfig()
	ctrl := controller.InitController()

	router := gin.Default()

	router.POST("/login", ctrl.HandleLogin)
	router.POST("/register", ctrl.HandleRegister)

	// Protected routes
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware([]byte(appConfig.JWTSecret)))
	{
		// Students routers
		protected.POST("/students", ctrl.CreateStudent)
		protected.GET("/students/:id", ctrl.GetStudent)
		protected.PUT("/students/:id", ctrl.UpdateStudent)
		protected.DELETE("/students/:id", ctrl.DeleteStudent)

		// Teachers routers
		protected.POST("/teachers", ctrl.CreateTeacher)
		protected.GET("/teachers/:id", ctrl.GetTeacher)
		router.PUT("/teachers/:id", ctrl.UpdateTeacher)
		router.DELETE("/teachers/:id", ctrl.DeleteTeacher)

		// Class routes
		protected.POST("/classes", ctrl.CreateClass)
		protected.PUT("/classes/:id/teacher", ctrl.UpdateClassTeacher)
		protected.PUT("/classes/:id/add-student", ctrl.AddStudentToClass)
		protected.PUT("/classes/:id/remove-student", ctrl.RemoveStudentFromClass)
		protected.GET("/classes/:id", ctrl.GetClass)
		protected.DELETE("/classes/:id", ctrl.DeleteClass)
	}

	router.Run("localhost:" + appConfig.Port)
}
