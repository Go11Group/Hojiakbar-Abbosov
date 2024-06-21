package routes

import (
    "github.com/gin-gonic/gin"
    "language-learning-app/controllers"
)

func SetupRoutes(router *gin.Engine) {
    userRoutes := router.Group("/users")
    {
        userRoutes.POST("/", controllers.CreateUser)
        userRoutes.GET("/:id", controllers.GetUser)
        userRoutes.GET("/", controllers.GetAllUsers)
        userRoutes.PUT("/:id", controllers.UpdateUser)
        userRoutes.DELETE("/:id", controllers.DeleteUser)
        userRoutes.GET("/search", controllers.SearchUsers)
        userRoutes.GET("/:id/courses", controllers.GetCoursesByUser)

    }

    courseRoutes := router.Group("/courses")
    {
        courseRoutes.GET("/popular", controllers.GetPopularCourses)
        courseRoutes.POST("/", controllers.CreateCourse)
        courseRoutes.GET("/:id", controllers.GetCourse)
        courseRoutes.GET("/", controllers.GetAllCourses)
        courseRoutes.PUT("/:id", controllers.UpdateCourse)
        courseRoutes.DELETE("/:id", controllers.DeleteCourse)
        courseRoutes.GET("/:id/lessons", controllers.GetLessonsByCourse)
        courseRoutes.GET("/:id/enrollments", controllers.GetEnrolledUsersByCourse)
    }

    lessonRoutes := router.Group("/lessons")
    {
        lessonRoutes.POST("/", controllers.CreateLesson)
        lessonRoutes.GET("/:id", controllers.GetLesson)
        lessonRoutes.GET("/", controllers.GetAllLessons)
        lessonRoutes.PUT("/:id", controllers.UpdateLesson)
        lessonRoutes.DELETE("/:id", controllers.DeleteLesson)
    }

    enrollmentRoutes := router.Group("/enrollments")
    {
        enrollmentRoutes.POST("/", controllers.EnrollUser)
        enrollmentRoutes.GET("/:id", controllers.GetEnrollment)
        enrollmentRoutes.GET("/", controllers.GetAllEnrollments)
        enrollmentRoutes.DELETE("/:id", controllers.DeleteEnrollment)
    }
}
