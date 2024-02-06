package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/laioncorcino/go-product-user/config"
	"github.com/laioncorcino/go-product-user/internal/entity"
	"github.com/laioncorcino/go-product-user/internal/handler"
	"github.com/laioncorcino/go-product-user/internal/infra/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	_, err := config.LoadConfig("../")
	if err != nil {
		panic(err)
	}

	fmt.Print("load configs")

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	_ = db.AutoMigrate(&entity.User{}, entity.Product{})

	productDB := database.NewProductDB(db)
	productHandle := handler.NewProductHandle(productDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/products", productHandle.CreateProduct)
	r.Get("/products", productHandle.GetProducts)
	r.Get("/products/{productId}", productHandle.GetProductByID)
	r.Put("/products/{productId}", productHandle.UpdateProduct)
	r.Delete("/products/{productId}", productHandle.DeleteProduct)

	_ = http.ListenAndServe(":8001", r)
}
