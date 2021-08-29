package repository

import (
	"golang-rest-crud/entities"

	"gorm.io/gorm"
)

type UserRepo interface {
	Save(user entities.User) (entities.User, error)
	FindByEmail(email string) (entities.User, error)
	FindByID(ID int) (entities.User, error)
	Update(user entities.User) (entities.User, error)
	FindAll() ([]entities.User, error)
	Delete(ID int) (entities.User, error)
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

func (r *userRepo) FindByEmail(email string) (entities.User, error) {
	var user entities.User
	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepo) FindByID(ID int) (entities.User, error) {
	var user entities.User

	err := r.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepo) Update(user entities.User) (entities.User, error) {
	err := r.db.Model(&user).Where("email = ?", user.Email).Updates(entities.User{Address: user.Address, Pwd: user.Pwd, UpdateAt: user.UpdateAt}).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepo) FindAll() ([]entities.User, error) {
	var users []entities.User

	err := r.db.Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}

func (r *userRepo) Delete(ID int) (entities.User, error) {
	var user entities.User
	err := r.db.Delete(&user, ID).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
