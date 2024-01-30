package entity

import (
	"github.com/laioncorcino/go-product-user/pkg"
	"golang.org/x/crypto/bcrypt"
	_ "golang.org/x/crypto/bcrypt"
)

type User struct {
	UserID   string
	Name     string
	Email    string
	Password string
}

func NewUser(name, email, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		UserID:   pkg.NewID(),
		Name:     name,
		Email:    email,
		Password: string(hash),
	}, nil
}

func (u User) ValidatePass(pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pass))
	return err == nil
}
