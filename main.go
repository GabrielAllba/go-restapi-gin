package main

import (
	"go-restapi-gin/controllers/productcontroller"
	"go-restapi-gin/controllers/reviewcontroller"
	"go-restapi-gin/controllers/usercontroller"
	"go-restapi-gin/initializers"
	"go-restapi-gin/middleware"
	"go-restapi-gin/models"

	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVariable()
	models.ConnectDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/api/user/signup", usercontroller.Signup)
	r.POST("/api/user/login", usercontroller.Login)
	r.POST("/api/user/logout", usercontroller.Logout)


	r.GET("/api/products", productcontroller.Index);
	r.GET("/api/product/:id", productcontroller.Show);
	r.POST("/api/product", productcontroller.Create);
	r.PUT("/api/product/:id", productcontroller.Update);
	r.DELETE("/api/product/:id", productcontroller.Delete);

	r.GET("/api/reviews", reviewcontroller.Index)
	r.GET("/api/review/:id", reviewcontroller.Show)
	r.POST("/api/review", reviewcontroller.Create)
	r.PUT("/api/review/:id", reviewcontroller.Update)
	r.DELETE("/api/review/:id", reviewcontroller.Delete)

	r.GET("/api/validate", middleware.RequireAuth, usercontroller.Validate)

	r.Run()

}