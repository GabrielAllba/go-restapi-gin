package productcontroller

import (
	"go-restapi-gin/models"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context){
	var products []models.Product

	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"products": products})

}
func Show(c *gin.Context){
	var product models.Product

	id := c.Param("id")

	if err := models.DB.First(&product, id).Error; err != nil{
		switch err{
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message" : "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message" : err})
			return
			
		}
	}

	c.JSON(http.StatusOK, gin.H{"product": product})

}

func Create(c *gin.Context) {
	var product models.Product

	// Bind other form data fields
	product.NamaProduct = c.PostForm("nama_product")
	product.Deskripsi = c.PostForm("deskripsi")

	// Check if required fields are missing
	if product.NamaProduct == "" || product.Deskripsi == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Missing required fields"})
		return
	}

	// Handle image upload
	file, err := c.FormFile("image")
	if err != nil {
		
		if err == http.ErrMissingFile {
			// No image uploaded, proceed without it
			models.DB.Create(&product)
			c.JSON(http.StatusOK, gin.H{"product": product})
			return
		}

		// c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Image is required"})
		// return
	}

	// Save image to the server
	filename := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, "uploads/"+filename); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to save image"})
		return
	}

	product.Image = "uploads/" + filename

	// Create the product in the database
	if err := models.DB.Create(&product).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to create product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}
func Update(c *gin.Context) {
	var product models.Product

	id := c.Param("id")

	if err := models.DB.First(&product, id).Error; err != nil{
		switch err{
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message" : "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message" : err})
			return
			
		}
	}

	// Bind other form data fields
	product.NamaProduct = c.PostForm("nama_product")
	product.Deskripsi = c.PostForm("deskripsi")

	// Check if required fields are missing
	if product.NamaProduct == "" || product.Deskripsi == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Missing required fields"})
		return
	}

	// Update the product in the database
	if err := models.DB.Model(&product).Where("id = ?", id).Updates(&product).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Failed to update product"})
		return
	}

	// Handle image upload if provided
	file, err := c.FormFile("image")
	if err == nil {
		// Save image to the server
		filename := filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file, "uploads/"+filename); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to save image"})
			return
		}

		product.Image = "uploads/" + filename

		// Save the updated product with the new image path
		if err := models.DB.Save(&product).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Failed to update product with new image"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbaharui"})
}


func Delete(c *gin.Context){
	var product models.Product

	id := c.Param("id")

	if models.DB.Delete(&product, id).RowsAffected == 0{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message":"Tidak dapat menghapus produk"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message" : "Data berhasil dihapus"})


}