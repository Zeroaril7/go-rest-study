package main

import (
	"net/http"

	"github.com/Zeroaril7/go-rest-study/controller"
	"github.com/Zeroaril7/go-rest-study/helper"
	"github.com/Zeroaril7/go-rest-study/middleware"
	"github.com/Zeroaril7/go-rest-study/pkg"
	"github.com/Zeroaril7/go-rest-study/pkg/exception"
	"github.com/Zeroaril7/go-rest-study/repository"
	"github.com/Zeroaril7/go-rest-study/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {

	db := pkg.NewDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.GetAll)
	router.GET("/api/categories/:categoryId", categoryController.GetById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
