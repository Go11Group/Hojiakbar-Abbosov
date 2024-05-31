package user

import (
	"module/models"
	"gorm.io/gorm"
)

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
	return db.Create(&user).Error
}

func ReadUsers(db *gorm.DB) ([]models.User, error) {
	var users []models.User
	err := db.Find(&users).Error
	return users, err
}

func UpdateUser(db *gorm.DB, id uint, email, password string) error {
	var user models.User
	err := db.First(&user, id).Error
	if err != nil {
		return err
	}
	user.Email = email
	user.Password = password
	return db.Save(&user).Error
}

func DeleteUser(db *gorm.DB, id uint) error {
	return db.Delete(&models.User{}, id).Error
}

func GetFilteredUsers(db *gorm.DB, firstName, lastName, email string, ageFrom, ageTo int, field, gender string, isEmployee bool) ([]models.User, error) {
	var users []models.User
	query := db.Model(&models.User{})

	if firstName != "" {
		query = query.Where("first_name = ?", firstName)
	}
	if lastName != "" {
		query = query.Where("last_name = ?", lastName)
	}
	if email != "" {
		query = query.Where("email = ?", email)
	}
	if ageFrom > 0 {
		query = query.Where("age >= ?", ageFrom)
	}
	if ageTo > 0 {
		query = query.Where("age <= ?", ageTo)
	}
	if field != "" {
		query = query.Where("field = ?", field)
	}
	if gender != "" {
		query = query.Where("gender = ?", gender)
	}
	query = query.Where("is_employee = ?", isEmployee)

	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
