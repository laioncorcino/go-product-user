package database

import (
	"github.com/laioncorcino/go-product-user/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"testing"
)

var db *gorm.DB
var userDB *UserDB

func TestMain(m *testing.M) {
	setup()
	defer dropDB()
	code := m.Run()
	clearDatabase()
	os.Exit(code)
}

func setup() {
	var err error
	db, err = gorm.Open(sqlite.Open("file::memory:"))
	if err != nil {
		log.Fatal(err)
	}

	_ = db.AutoMigrate(&entity.User{})
	userDB = NewUserDB(db)
}

func dropDB() {
	s, _ := db.DB()
	s.Close()
}

func clearDatabase() {
	db.Delete(&entity.User{})
}

func TestUserDB_Create(t *testing.T) {
	user, _ := entity.NewUser("Jonh", "jonh@email.com", "12345")
	err := userDB.Create(user)

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
	user, _ := entity.NewUser("Jonh", "jonh@email.com", "12345")
	db.Create(user)

	userSaved, err := userDB.FindByEmail(user.Email)

	assert.Nil(t, err)
	assert.Equal(t, userSaved.UserID, user.UserID)
	assert.Equal(t, userSaved.Name, user.Name)
	assert.Equal(t, userSaved.Email, user.Email)
	assert.NotEmpty(t, userSaved.Password)
}
