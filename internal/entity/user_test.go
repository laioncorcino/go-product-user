package entity

import (
	"github.com/laioncorcino/go-product-user/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Jonh Doe", "jonh@email.com", "12345")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.UserID)
	assert.True(t, pkg.IsUUID(user.UserID))
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Jonh Doe", user.Name)
	assert.Equal(t, "jonh@email.com", user.Email)
}

func TestUser_ValidatePass(t *testing.T) {
	user, err := NewUser("Jonh Doe", "jonh@email.com", "12345")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePass("12345"))
	assert.False(t, user.ValidatePass("54321"))
	assert.NotEqual(t, "12345", user.Password)
}
