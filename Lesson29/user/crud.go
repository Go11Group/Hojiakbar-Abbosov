package user

import (
	"gorm.io/gorm"
	"module/models"
)

// CreateUser inserts a new user into the database
func CreateUser(db *gorm.DB, firstName, lastName, email, password string, age int, field, gender string, isEmployee bool) error {
	user := models.User{
		FirstName:  firstName,
		LastName:   lastName,
		Email:      email,
		Password:   password,
		Age:        age,
		Field:      field,
		Gender:     gender,
		IsEmployee: isEmployee,
	}
	result := db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// ReadUsers retrieves all users from the database
func ReadUsers(db *gorm.DB) ([]models.User, error) {
	var users []models.User
	result := db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// UpdateUser updates the email and password of a user based on their ID
func UpdateUser(db *gorm.DB, id uint, email, newPassword string) error {
	var user models.User
	result := db.First(&user, id)
	if result.Error != nil {
		return result.Error
	}
	user.Email = email
	user.Password = newPassword
	db.Save(&user)
	return nil
}

// DeleteUser deletes a user based on their ID
func DeleteUser(db *gorm.DB, id uint) error {
	var user models.User
	result := db.First(&user, id)
	if result.Error != nil {
		return result.Error
	}
	db.Delete(&user)
	return nil
}
