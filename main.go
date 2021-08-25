package main

import (
	"fmt"
	"go-restful-api/app"
	"go-restful-api/controller"
	"go-restful-api/helper"
	"go-restful-api/repository"
	"go-restful-api/service"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	helper.PanicIfError(err)

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	serverPort := os.Getenv("SERVER_PORT")
	dbUser := os.Getenv("USER")
	dbPassword := os.Getenv("PASSWORD")
	dbName := os.Getenv("DBNAME")

	validate := validator.New()
	db := app.NewDB(port, host, dbUser, dbPassword, dbName)
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	server := http.Server{
		Addr:    host + ":" + serverPort,
		Handler: router,
	}
	fmt.Printf("Listening to port %s", port)
	err = server.ListenAndServe()
	helper.PanicIfError(err)
}
