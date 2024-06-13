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

func CreateLesson(c *gin.Context) {
    var lesson models.Lesson
    if err := c.ShouldBindJSON(&lesson); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    lesson.LessonID = uuid.New()
    lesson.CreatedAt = time.Now()
    lesson.UpdatedAt = time.Now()

    _, err := database.DB.Exec(`INSERT INTO lessons (lesson_id, course_id, title, content, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`,
        lesson.LessonID, lesson.CourseID, lesson.Title, lesson.Content, lesson.CreatedAt, lesson.UpdatedAt)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, lesson)
}

func GetLesson(c *gin.Context) {
    id := c.Param("id")
    var lesson models.Lesson
    err := database.DB.QueryRow(`SELECT lesson_id, course_id, title, content, created_at, updated_at, deleted_at FROM lessons WHERE lesson_id = $1`, id).Scan(
        &lesson.LessonID, &lesson.CourseID, &lesson.Title, &lesson.Content, &lesson.CreatedAt, &lesson.UpdatedAt, &lesson.DeletedAt)
    if err != nil {
        if err == sql.ErrNoRows {
            c.JSON(http.StatusNotFound, gin.H{"error": "Lesson not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, lesson)
}

func GetAllLessons(c *gin.Context) {
    rows, err := database.DB.Query(`SELECT lesson_id, course_id, title, content, created_at, updated_at, deleted_at FROM lessons`)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    var lessons []models.Lesson
    for rows.Next() {
        var lesson models.Lesson
        if err := rows.Scan(&lesson.LessonID, &lesson.CourseID, &lesson.Title, &lesson.Content, &lesson.CreatedAt, &lesson.UpdatedAt, &lesson.DeletedAt); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        lessons = append(lessons, lesson)
    }
    c.JSON(http.StatusOK, lessons)
}

func UpdateLesson(c *gin.Context) {
    id := c.Param("id")
    var lesson models.Lesson
    if err := c.ShouldBindJSON(&lesson); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    lesson.UpdatedAt = time.Now()

    _, err := database.DB.Exec(`UPDATE lessons SET title = $1, content = $2, updated_at = $3 WHERE lesson_id = $4`,
        lesson.Title, lesson.Content, lesson.UpdatedAt, id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, lesson)
}

func DeleteLesson(c *gin.Context) {
    id := c.Param("id")
    deletedAt := time.Now()

    _, err := database.DB.Exec(`UPDATE lessons SET deleted_at = $1 WHERE lesson_id = $2`, deletedAt, id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Lesson deleted successfully"})
}
