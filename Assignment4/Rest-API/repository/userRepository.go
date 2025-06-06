package repository

import (
	"fmt"
	"rest-api/interfaces"
	"rest-api/models"
	"rest-api/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(database *gorm.DB) interfaces.UserRepository {
	return &userRepo{db: database}
}

func (u *userRepo) GetAllUsers() ([]models.User, error) {
	var users []models.User

	if err := u.db.Find(&users).Error; err != nil {
		return nil, fmt.Errorf("error fetching all users %w", err)
	}

	return users, nil
}

func (r *userRepo) CreateUser(m models.User) (int64, error) {
	result := r.db.Create(&m)
	if result.Error != nil {
		logrus.Errorf("Unable to create user in db: %v", result.Error)
		return 0, fmt.Errorf("could not create user: %w", result.Error)
	}
	return result.RowsAffected, nil
}

func (r *userRepo) Login(passedUser models.User) (*models.User, error) {
	var user models.User
	err := r.db.Where("username = ?", passedUser.Username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("error fetching user %s: %w", passedUser.Username, err)
	}

	passwordMatches, cmpErr := utils.ComparePasswords(user.Password, passedUser.Password)
	if cmpErr != nil {
		logrus.Errorf("Error comparing passwords for user %s: %v", passedUser.Username, cmpErr)
		return nil, fmt.Errorf("internal server error during login")
	}
	if !passwordMatches {
		logrus.Warnf("Password does not match for user %s", passedUser.Username)
		return nil, fmt.Errorf("invalid credentials")
	}

	logrus.Debugf("Password matches for user: %s", user.Username)
	return &user, nil
}
