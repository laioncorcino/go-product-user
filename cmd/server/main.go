package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/laioncorcino/go-product-user/config"
	"github.com/laioncorcino/go-product-user/internal/entity"
	"github.com/laioncorcino/go-product-user/internal/handler"
	"github.com/laioncorcino/go-product-user/internal/infra/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	conf, err := config.LoadConfig("../")
	if err != nil {
		panic(err)
	}

	fmt.Print("load configs\n")

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	_ = db.AutoMigrate(&entity.User{}, entity.Product{})

	productDB := database.NewProductDB(db)
	userDB := database.NewUserDB(db)

	productHandle := handler.NewProductHandler(productDB)
	userHandle := handler.NewUserHandler(userDB)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", conf.TokenAuth))
	r.Use(middleware.WithValue("expires", conf.JWTExpiresIn))

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(conf.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Post("/", productHandle.CreateProduct)
		r.Get("/", productHandle.GetProducts)
		r.Get("/{productId}", productHandle.GetProductByID)
		r.Put("/{productId}", productHandle.UpdateProduct)
		r.Delete("/{productId}", productHandle.DeleteProduct)
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/", userHandle.CreateUser)
		r.Post("/token", userHandle.Login)
	})

	_ = http.ListenAndServe(":8001", r)
}
