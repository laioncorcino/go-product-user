package main

import "github.com/laioncorcino/go-product-user/config"

func main() {
	conf, _ := config.LoadConfig(".")
	println(conf.DBHost)
	println(conf.WebServerPort)
	println(conf.DBDriver)
	println(conf.DBPort)
	println(conf.DBUser)
	println(conf.DBPassword)
	println(conf.DBName)
	println(conf.JWTSecret)
	println(conf.JWTExpiresIn)
	println(conf.TokenAuth)
}
