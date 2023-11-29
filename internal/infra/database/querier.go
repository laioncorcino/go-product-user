package database

import "github.com/laioncorcino/go-product-user/internal/entity"

type UserQuery interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
