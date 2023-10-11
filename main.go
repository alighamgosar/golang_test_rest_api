package main

//
import (
	"github.com/alighamgosar/golang_test_rest_api/controllers"
	"github.com/alighamgosar/golang_test_rest_api/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConntectToDB()
}
func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)

	r.GET("/posts", controllers.PostsIndex)

	r.GET("/posts/:id", controllers.PostsShow)

	r.PUT("/posts/:id", controllers.PostsUpdate)

	r.DELETE("posts/:id", controllers.PostsDelete)
	r.Run() // listen and serve on 0.0.0.0:8080
}
