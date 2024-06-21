package models

import (
    "time"
    "github.com/google/uuid"
)

type User struct {
    UserID    uuid.UUID `json:"user_id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    Birthday  time.Time `json:"birthday"`
    Password  string    `json:"password"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    DeletedAt *time.Time `json:"deleted_at"`
}

type Course struct {
    CourseID   uuid.UUID `json:"course_id"`
    Title      string    `json:"title"`
    Description string   `json:"description"`
    CreatedAt  time.Time `json:"created_at"`
    UpdatedAt  time.Time `json:"updated_at"`
    DeletedAt  *time.Time `json:"deleted_at"`
}

type Lesson struct {
    LessonID   uuid.UUID `json:"lesson_id"`
    CourseID   uuid.UUID `json:"course_id"`
    Title      string    `json:"title"`
    Content    string    `json:"content"`
    CreatedAt  time.Time `json:"created_at"`
    UpdatedAt  time.Time `json:"updated_at"`
    DeletedAt  *time.Time `json:"deleted_at"`
}

type Enrollment struct {
    EnrollmentID uuid.UUID `json:"enrollment_id"`
    UserID       uuid.UUID `json:"user_id"`
    CourseID     uuid.UUID `json:"course_id"`
    EnrollmentDate time.Time `json:"enrollment_date"`
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
    DeletedAt    *time.Time `json:"deleted_at"`
}
