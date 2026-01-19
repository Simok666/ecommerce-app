package controllers

import (
	"net/http"
	"strconv"

	"github.com/Simok666/ecommerce-app.git/internal/database"
	"github.com/Simok666/ecommerce-app.git/internal/models"
	"github.com/Simok666/ecommerce-app.git/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateProduct(c *gin.Context) {
	name := c.PostForm("name")
	price, _ := strconv.ParseFloat(c.PostForm("price"), 64)
	stock, _ := strconv.Atoi(c.PostForm("stock"))
	description := c.PostForm("description")

	form, _ := c.MultipartForm()
	files := form.File["images"]

	if len(files) == 0 {
		c.JSON(400, gin.H{"error": "image required"})
		return
	}

	product := models.Product{
		Name:        name,
		Price:       price,
		Stock:       stock,
		Description: description,
	}

	tx := database.DB.Begin()
	tx.Create(&product)

	for i, file := range files {
		img, thumb, err := services.SaveProductImage(file)
		if err != nil {
			tx.Rollback()
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		tx.Create(&models.ProductImage{
			ProductID:    product.ID,
			ImageURL:     img,
			ThumbnailURL: thumb,
			IsPrimary:    i == 0,
		})
	}

	tx.Commit()
	c.JSON(201, gin.H{"message": "Product created successfully", "data": product})
}

func GetProduct(c *gin.Context) {
	var products []models.Product

	database.DB.Where("is_active = ?", true).Find(&products)

	c.JSON(http.StatusOK, products)
}

func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if err := database.DB.First(&product, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func UpdateProduct(c *gin.Context) {
	var product models.Product

	id := c.Param("id")
	price, _ := strconv.ParseFloat(c.PostForm("price"), 64)
	stock, _ := strconv.Atoi(c.PostForm("stock"))
	description := c.PostForm("description")

	database.DB.Model(&models.Product{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"name":        c.PostForm("name"),
			"price":       price,
			"stock":       stock,
			"description": description,
		})

	form, _ := c.MultipartForm()
	files := form.File["images"]

	for _, file := range files {
		img, thumb, _ := services.SaveProductImage(file)
		database.DB.Create(&models.ProductImage{
			ProductID:    uuid.MustParse(id),
			ImageURL:     img,
			ThumbnailURL: thumb,
		})
	}

	c.JSON(200, gin.H{"message": "updated", "data": product})
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	if err := database.DB.Delete(&models.Product{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product Deleted"})
}

func DeleteProductImage(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&models.ProductImage{}, "id = ?", id)
	c.JSON(200, gin.H{"message": "image deleted"})
}
