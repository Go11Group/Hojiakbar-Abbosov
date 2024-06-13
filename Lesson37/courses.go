package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "time"
    "language-learning-app/database"
    "language-learning-app/models"
	"database/sql"

)

func CreateCourse(c *gin.Context) {
    var course models.Course
    if err := c.ShouldBindJSON(&course); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    course.CourseID = uuid.New()
    course.CreatedAt = time.Now()
    course.UpdatedAt = time.Now()

    _, err := database.DB.Exec(`INSERT INTO courses (course_id, title, description, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`,
        course.CourseID, course.Title, course.Description, course.CreatedAt, course.UpdatedAt)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, course)
}

func GetCourse(c *gin.Context) {
    id := c.Param("id")
    var course models.Course
    err := database.DB.QueryRow(`SELECT course_id, title, description, created_at, updated_at, deleted_at FROM courses WHERE course_id = $1`, id).Scan(
        &course.CourseID, &course.Title, &course.Description, &course.CreatedAt, &course.UpdatedAt, &course.DeletedAt)
    if err != nil {
        if err == sql.ErrNoRows {
            c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, course)
}

func GetAllCourses(c *gin.Context) {
    rows, err := database.DB.Query(`SELECT course_id, title, description, created_at, updated_at, deleted_at FROM courses`)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    var courses []models.Course
    for rows.Next() {
        var course models.Course
        if err := rows.Scan(&course.CourseID, &course.Title, &course.Description, &course.CreatedAt, &course.UpdatedAt, &course.DeletedAt); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        courses = append(courses, course)
    }
    c.JSON(http.StatusOK, courses)
}

func UpdateCourse(c *gin.Context) {
    id := c.Param("id")
    var course models.Course
    if err := c.ShouldBindJSON(&course); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    course.UpdatedAt = time.Now()

    _, err := database.DB.Exec(`UPDATE courses SET title = $1, description = $2, updated_at = $3 WHERE course_id = $4`,
        course.Title, course.Description, course.UpdatedAt, id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, course)
}

func DeleteCourse(c *gin.Context) {
    id := c.Param("id")
    deletedAt := time.Now()

    _, err := database.DB.Exec(`UPDATE courses SET deleted_at = $1 WHERE course_id = $2`, deletedAt, id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Course deleted successfully"})
}

func GetLessonsByCourse(c *gin.Context) {
    courseID := c.Param("course_id")

    rows, err := database.DB.Query(`SELECT lesson_id, title, content FROM lessons WHERE course_id = $1 AND deleted_at IS NULL`, courseID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    var lessons []models.Lesson
    for rows.Next() {
        var lesson models.Lesson
        if err := rows.Scan(&lesson.LessonID, &lesson.Title, &lesson.Content); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        lessons = append(lessons, lesson)
    }
    c.JSON(http.StatusOK, gin.H{"course_id": courseID, "lessons": lessons})
}

func GetEnrolledUsersByCourse(c *gin.Context) {
    courseID := c.Param("course_id")

    rows, err := database.DB.Query(`SELECT u.user_id, u.name, u.email FROM enrollments e JOIN users u ON e.user_id = u.user_id WHERE e.course_id = $1 AND u.deleted_at IS NULL`, courseID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    var users []models.User
    for rows.Next() {
        var user models.User
        if err := rows.Scan(&user.UserID, &user.Name, &user.Email); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        users = append(users, user)
    }
    c.JSON(http.StatusOK, gin.H{"course_id": courseID, "enrolled_users": users})
}

func GetPopularCourses(c *gin.Context) {
    startDate := c.Query("start_date")
    endDate := c.Query("end_date")

    rows, err := database.DB.Query(`SELECT c.course_id, c.title, COUNT(e.enrollment_id) as enrollments_count FROM enrollments e JOIN courses c ON e.course_id = c.course_id WHERE e.enrollment_date BETWEEN $1 AND $2 AND c.deleted_at IS NULL GROUP BY c.course_id, c.title ORDER BY enrollments_count DESC`, startDate, endDate)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    var popularCourses []struct {
        CourseID        uuid.UUID `json:"course_id"`
        CourseTitle     string    `json:"course_title"`
        EnrollmentsCount int      `json:"enrollments_count"`
    }
    for rows.Next() {
        var course struct {
            CourseID        uuid.UUID `json:"course_id"`
            CourseTitle     string    `json:"course_title"`
            EnrollmentsCount int      `json:"enrollments_count"`
        }
        if err := rows.Scan(&course.CourseID, &course.CourseTitle, &course.EnrollmentsCount); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        popularCourses = append(popularCourses, course)
    }
    c.JSON(http.StatusOK, gin.H{"time_period": gin.H{"start_date": startDate, "end_date": endDate}, "popular_courses": popularCourses})
}
