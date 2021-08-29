package repository

import (
	"golang-rest-crud/entities"

	"gorm.io/gorm"
)

type UserRepo interface {
	Save(user entities.User) (entities.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepo {
	return &userRepo{db}
}

func (r *userRepo) Save(user entities.User) (entities.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
