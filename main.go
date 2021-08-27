package main

import (
	"fmt"
	"go-restful-api/app"
	"go-restful-api/controller"
	"go-restful-api/helper"
	"go-restful-api/middleware"
	"go-restful-api/repository"
	"go-restful-api/service"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	helper.PanicIfError(err)

	host := os.Getenv("HOST")
	dbport := os.Getenv("PORT")
	serverPort := os.Getenv("SERVER_PORT")
	dbUser := os.Getenv("USER")
	dbPassword := os.Getenv("PASSWORD")
	dbName := os.Getenv("DBNAME")

	validate := validator.New()
	db := app.NewDB(dbport, host, dbUser, dbPassword, dbName)
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    host + ":" + serverPort,
		Handler: middleware.NewAuthMiddleware(router),
	}
	fmt.Printf("Listening to port %s", serverPort)
	err = server.ListenAndServe()
	helper.PanicIfError(err)
}
