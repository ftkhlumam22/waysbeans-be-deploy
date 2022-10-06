package repositories

import (
	"waysbeans/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	RegisterUser(user models.User) (models.User, error)
	FindUserByEmail(email string) (*models.User, error)
	Login(email string) (models.User, error)
	Getuser(ID int) (models.User, error)
}

func RepositoryAuth(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) RegisterUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) Login(email string) (models.User, error) {
	var user models.User
	err := r.db.First(&user, "email=?", email).Error

	return user, err
}

func (r *repository) FindUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := r.db.Where("email = ?", email).First(&user).Error

	return user, err
}

func (r *repository) Getuser(ID int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, ID).Error

	return user, err
}
