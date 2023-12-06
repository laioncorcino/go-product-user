package database

import (
	"github.com/laioncorcino/go-product-user/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"testing"
)

var db *gorm.DB
var userDB *UserDB
var productDB *ProductDB

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

	_ = db.AutoMigrate(&entity.Product{})
	productDB = NewProductDB(db)
}

func dropDB() {
	s, _ := db.DB()
	s.Close()
}

func clearDatabase() {
	userDB.DB.Unscoped().Delete(&entity.User{})
	productDB.DB.Unscoped().Delete(&entity.Product{})
}
