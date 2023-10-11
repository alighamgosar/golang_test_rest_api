package main

import (
	"github.com/alighamgosar/golang_test_rest_api/initializers"
	"github.com/alighamgosar/golang_test_rest_api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConntectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
