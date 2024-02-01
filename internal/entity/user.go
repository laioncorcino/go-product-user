package entity

import (
	"errors"
	"github.com/laioncorcino/go-product-user/pkg"
	"golang.org/x/crypto/bcrypt"
	_ "golang.org/x/crypto/bcrypt"
)

var ErrEmailIsRequired = errors.New("email is required")

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

	user := &User{
		UserID:   pkg.NewID(),
		Name:     name,
		Email:    email,
		Password: string(hash),
	}

	err = user.Validate()

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) Validate() error {
	if u.Name == "" {
		return ErrNameIsRequired
	}

	if u.Email == "" {
		return ErrEmailIsRequired
	}

	return nil
}

func (u *User) ValidatePass(pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pass))
	return err == nil
}
