package database

import (
	"fmt"
	"github.com/laioncorcino/go-product-user/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
	"testing"
	"time"
)

func TestProductDB_Create(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	_ = db.AutoMigrate(&entity.Product{})
	productDB := NewProductDB(db)

	product, err := entity.NewProduct("Monitor", 200.00)

	err = productDB.Create(product)

	assert.Nil(t, err)

	var productSaved entity.Product
	err = db.First(&productSaved, "product_id = ?", product.ProductID).Error

	assert.Nil(t, err)
	assert.Equal(t, productSaved.ProductID, product.ProductID)
	assert.Equal(t, productSaved.Name, product.Name)
	assert.Equal(t, productSaved.Price, product.Price)
	assert.Equal(t, productSaved.CreatedAt.In(time.Local), product.CreatedAt.In(time.Local))
}

func TestProductDB_FindAll(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	_ = db.AutoMigrate(&entity.Product{})
	productDB := NewProductDB(db)

	for i := 1; i < 35; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.NoError(t, err)
		db.Create(product)
	}

	products, err := productDB.FindAll(1, 10, "asc")

	assert.Nil(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productDB.FindAll(1, 10, "desc")

	assert.Nil(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 34", products[0].Name)
	assert.Equal(t, "Product 25", products[9].Name)

	products, err = productDB.FindAll(2, 20, "")

	assert.Nil(t, err)
	assert.Len(t, products, 14)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 34", products[13].Name)
}

func TestProductDB_FindByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	_ = db.AutoMigrate(&entity.Product{})
	productDB := NewProductDB(db)

	product, err := entity.NewProduct("Monitor", 200.00)
	db.Create(product)

	productSaved, err := productDB.FindByID(product.ProductID)

	assert.NoError(t, err)
	assert.NotNil(t, productSaved)
	assert.Equal(t, "Monitor", productSaved.Name)
}

func TestProductDB_Update(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	_ = db.AutoMigrate(&entity.Product{})
	productDB := NewProductDB(db)

	product, err := entity.NewProduct("Monitor", 200.00)
	db.Create(product)

	product.Price = 236.15
	err = productDB.Update(product)

	assert.Nil(t, err)

	productUpdated, err := productDB.FindByID(product.ProductID)

	assert.Nil(t, err)
	assert.Equal(t, "Monitor", productUpdated.Name)
	assert.Equal(t, 236.15, productUpdated.Price)
}

func TestProductDB_Delete(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	_ = db.AutoMigrate(&entity.Product{})
	productDB := NewProductDB(db)

	product, err := entity.NewProduct("Monitor", 200.00)
	db.Create(product)

	err = productDB.Delete(product.ProductID)

	assert.NoError(t, err)

	p, err := productDB.FindByID(product.ProductID)

	assert.Error(t, err)
	assert.Equal(t, "", p.Name)
	assert.Equal(t, "", p.ProductID)
	assert.Equal(t, 0.0, p.Price)
}
