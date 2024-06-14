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

func EnrollUser(c *gin.Context) {
    var enrollment models.Enrollment
    if err := c.ShouldBindJSON(&enrollment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    enrollment.EnrollmentID = uuid.New()
    enrollment.EnrollmentDate = time.Now()
    enrollment.CreatedAt = time.Now()
    enrollment.UpdatedAt = time.Now()

    _, err := database.DB.Exec(`INSERT INTO enrollments (enrollment_id, user_id, course_id, enrollment_date, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`,
        enrollment.EnrollmentID, enrollment.UserID, enrollment.CourseID, enrollment.EnrollmentDate, enrollment.CreatedAt, enrollment.UpdatedAt)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, enrollment)
}

func GetEnrollment(c *gin.Context) {
    id := c.Param("id")
    var enrollment models.Enrollment
    err := database.DB.QueryRow(`SELECT enrollment_id, user_id, course_id, enrollment_date, created_at, updated_at, deleted_at FROM enrollments WHERE enrollment_id = $1`, id).Scan(
        &enrollment.EnrollmentID, &enrollment.UserID, &enrollment.CourseID, &enrollment.EnrollmentDate, &enrollment.CreatedAt, &enrollment.UpdatedAt, &enrollment.DeletedAt)
    if err != nil {
        if err == sql.ErrNoRows {
            c.JSON(http.StatusNotFound, gin.H{"error": "Enrollment not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, enrollment)
}

func GetAllEnrollments(c *gin.Context) {
    rows, err := database.DB.Query(`SELECT enrollment_id, user_id, course_id, enrollment_date, created_at, updated_at, deleted_at FROM enrollments`)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    var enrollments []models.Enrollment
    for rows.Next() {
        var enrollment models.Enrollment
        if err := rows.Scan(&enrollment.EnrollmentID, &enrollment.UserID, &enrollment.CourseID, &enrollment.EnrollmentDate, &enrollment.CreatedAt, &enrollment.UpdatedAt, &enrollment.DeletedAt); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        enrollments = append(enrollments, enrollment)
    }
    c.JSON(http.StatusOK, enrollments)
}

func DeleteEnrollment(c *gin.Context) {
    id := c.Param("id")
    deletedAt := time.Now()

    _, err := database.DB.Exec(`UPDATE enrollments SET deleted_at = $1 WHERE enrollment_id = $2`, deletedAt, id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Enrollment deleted successfully"})
}
