package database

import "github.com/laioncorcino/go-product-user/internal/entity"

type UserQuery interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type ProductQuery interface {
	Create(product *entity.Product) error
	FindAll(page, limit int, sort string) ([]entity.Product, error)
	FindByID(productId string) (*entity.Product, error)
	Update(product *entity.Product) error
	Delete(productId string) error
}
