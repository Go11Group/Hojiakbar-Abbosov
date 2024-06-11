package routes

import (
	"github.com/gin-gonic/gin"
	"go-gin-crud/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/problems", controllers.GetProblem)
	r.GET("/problems/:id", controllers.GetProblem)
	r.POST("/problems", controllers.CreateProblem)
	r.PUT("/problems/:id", controllers.UpdateProblem)
	r.DELETE("/problems/:id", controllers.DeleteProblem)

	r.GET("/users", controllers.GetUsers)
	r.GET("/users/:id", controllers.GetUser)
	r.POST("/users", controllers.CreateUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	r.GET("/solved", controllers.GetSolvedProblems)
	r.GET("/solved/:id", controllers.GetSolvedProblem)
	r.POST("/solved", controllers.CreateSolvedProblem)
	r.PUT("/solved/:id", controllers.UpdateSolvedProblem)
	r.DELETE("/solved/:id", controllers.DeleteSolvedProblem)

	return r
}
