package database

import (
	"github.com/laioncorcino/go-product-user/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestUserDB_Create(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	_ = db.AutoMigrate(&entity.User{})
	userDB := NewUserDB(db)

	user, _ := entity.NewUser("Jonh", "jonh@email.com", "12345")
	err = userDB.Create(user)

	assert.Nil(t, err)

	var userSaved entity.User
	err = db.First(&userSaved, "user_id = ?", user.UserID).Error

	assert.Nil(t, err)
	assert.Equal(t, userSaved.UserID, user.UserID)
	assert.Equal(t, userSaved.Name, user.Name)
	assert.Equal(t, userSaved.Email, user.Email)
	assert.NotEmpty(t, userSaved.Password)
}

func TestUserDB_FindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	_ = db.AutoMigrate(&entity.User{})
	userDB := NewUserDB(db)

	user, _ := entity.NewUser("Jonh", "jonh@email.com", "12345")
	err = userDB.Create(user)

	userSaved, err := userDB.FindByEmail(user.Email)

	assert.Nil(t, err)
	assert.Equal(t, userSaved.UserID, user.UserID)
	assert.Equal(t, userSaved.Name, user.Name)
	assert.Equal(t, userSaved.Email, user.Email)
	assert.NotEmpty(t, userSaved.Password)

	db.Unscoped().Delete(&entity.User{})
}
