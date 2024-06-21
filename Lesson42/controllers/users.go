package controllers

import (
    "database/sql"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "language-learning-app/database"
    "language-learning-app/models"
    "strings"
    "strconv"
)

func CreateUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    user.UserID = uuid.New()
    user.CreatedAt = time.Now()
    user.UpdatedAt = time.Now()

    _, err := database.DB.Exec(`INSERT INTO users (user_id, name, email, birthday, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`,
        user.UserID, user.Name, user.Email, user.Birthday, user.Password, user.CreatedAt, user.UpdatedAt)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, user)
}

func GetUser(c *gin.Context) {
    id := c.Param("id")
    var user models.User
    err := database.DB.QueryRow(`SELECT user_id, name, email, birthday, password, created_at, updated_at, deleted_at FROM users WHERE user_id = $1`, id).Scan(
        &user.UserID, &user.Name, &user.Email, &user.Birthday, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
    if err != nil {
        if err == sql.ErrNoRows {
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, user)
}

func GetAllUsers(c *gin.Context) {
    rows, err := database.DB.Query(`SELECT user_id, name, email, birthday, password, created_at, updated_at, deleted_at FROM users`)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    var users []models.User
    for rows.Next() {
        var user models.User
        if err := rows.Scan(&user.UserID, &user.Name, &user.Email, &user.Birthday, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        users = append(users, user)
    }
    c.JSON(http.StatusOK, gin.H{"results": users})
}

func GetCoursesByUser(c *gin.Context) {
    userID := c.Param("user_id")

    rows, err := database.DB.Query(`SELECT c.course_id, c.title, c.description FROM enrollments e JOIN courses c ON e.course_id = c.course_id WHERE e.user_id = $1 AND c.deleted_at IS NULL`, userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    var courses []models.Course
    for rows.Next() {
        var course models.Course
        if err := rows.Scan(&course.CourseID, &course.Title, &course.Description); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        courses = append(courses, course)
    }
    if err := rows.Err(); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"user_id": userID, "courses": courses})
}

func UpdateUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    userID := c.Param("id")
    user.UpdatedAt = time.Now()

    _, err := database.DB.Exec(`UPDATE users SET name = $1, email = $2, birthday = $3, password = $4, updated_at = $5 WHERE user_id = $6`,
        user.Name, user.Email, user.Birthday, user.Password, user.UpdatedAt, userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func DeleteUser(c *gin.Context) {
    userID := c.Param("id")
    deletedAt := time.Now()

    _, err := database.DB.Exec(`UPDATE users SET deleted_at = $1 WHERE user_id = $2`, deletedAt, userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func SearchUsers(c *gin.Context) {
    name := c.Query("name")
    email := c.Query("email")

    var queryBuilder strings.Builder
    queryBuilder.WriteString("SELECT user_id, name, email, birthday, password, created_at, updated_at, deleted_at FROM users WHERE 1=1")

    var args []interface{}
    argCounter := 1

    if name != "" {
        queryBuilder.WriteString(" AND name ILIKE $" + strconv.Itoa(argCounter))
        args = append(args, "%"+name+"%")
        argCounter++
    }

    if email != "" {
        queryBuilder.WriteString(" AND email ILIKE $" + strconv.Itoa(argCounter))
        args = append(args, "%"+email+"%")
        argCounter++
    }

    query := queryBuilder.String()
    rows, err := database.DB.Query(query, args...)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    var users []models.User
    for rows.Next() {
        var user models.User
        if err := rows.Scan(&user.UserID, &user.Name, &user.Email, &user.Birthday, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        users = append(users, user)
    }
    c.JSON(http.StatusOK, gin.H{"results": users})
}
