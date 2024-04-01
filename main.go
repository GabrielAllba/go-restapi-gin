package main

import (
	"go-restapi-gin/controllers/productcontroller"
	"go-restapi-gin/controllers/reviewcontroller"
	"go-restapi-gin/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

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

	r.Run()

}