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

func (ur *UserDB) Create(user *entity.User) error {
	return ur.DB.Create(user).Error
}

func (ur *UserDB) FindByEmail(email string) (*entity.User, error) {
	var user entity.User

	err := ur.DB.
		Where("email = ?", email).
		First(&user).
		Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
