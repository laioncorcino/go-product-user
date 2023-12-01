package database

import (
	"github.com/laioncorcino/go-product-user/internal/entity"
	"gorm.io/gorm"
)

type UserDB struct {
	DB *gorm.DB
}

func NewUserDB(db *gorm.DB) *UserDB {
	return &UserDB{DB: db}
}

func (u *UserDB) Create(user *entity.User) error {
	return u.DB.Create(user).Error
}

func (u *UserDB) FindByEmail(email string) (*entity.User, error) {
	var user entity.User

	err := u.DB.
		Where("email = ?", email).
		First(&user).
		Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
