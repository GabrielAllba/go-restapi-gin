package reviewcontroller

import (
	"go-restapi-gin/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var reviews []models.Review
	models.DB.Find(&reviews)
	c.JSON(http.StatusOK, gin.H{"reviews": reviews})
}

func Show(c *gin.Context) {
	var review models.Review
	id := c.Param("id")
	if err := models.DB.First(&review, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Review not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"review": review})
}

func Create(c *gin.Context) {
	var review models.Review

	// Bind form data fields
	review.Content = c.PostForm("content")
	
	// Convert product_id to int
	productID, err := strconv.Atoi(c.PostForm("product_id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Product id tidak boleh kosong"})
		return
	}
	review.ProductId = productID

	// Check if required fields are missing
	if review.Content == "" || review.ProductId == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Missing required fields"})
		return
	}

	// Create the review in the database
	if err := models.DB.Create(&review).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to create review"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"review": review})
}

func Update(c *gin.Context) {
	var review models.Review
	id := c.Param("id")
	if err := models.DB.First(&review, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Review not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}
	}

	// Bind form data fields
	review.Content = c.PostForm("content")
	
	// Convert product_id to int
	productID, err := strconv.Atoi(c.PostForm("product_id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid product_id"})
		return
	}
	review.ProductId = productID

	// Check if required fields are missing
	if review.Content == "" || review.ProductId == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Missing required fields"})
		return
	}

	// Update the review in the database
	if err := models.DB.Save(&review).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to update review"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Review updated successfully"})
}


func Delete(c *gin.Context) {
	var review models.Review
	id := c.Param("id")
	if err := models.DB.First(&review, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Review not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}
	}
	if err := models.DB.Delete(&review, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete review"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Review deleted successfully"})
}
