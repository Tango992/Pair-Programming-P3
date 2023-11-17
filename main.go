package main

import (
	"pair-programming/config"
	"pair-programming/helpers"

	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Validator = &helpers.CustomValidator{NewValidator: validator.New()}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	
	db := config.ConnectDB().Database("pair-programmingDB")
	transactionCollection := db.Collection("transactions")

	
	e.Logger.Fatal(e.Start(":8080"))
}