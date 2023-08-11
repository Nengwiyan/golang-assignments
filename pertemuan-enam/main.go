package main

import (
	"errors"
	"fmt"
	"pertemuan-enam/database"
	"pertemuan-enam/models"

	"gorm.io/gorm"
)

func main() {
	database.StartDB()

	// createProduct("Smartphone")
	// getProductById(1)
	// updateProductById(1, "Iphone")
	// createVariant(1, "Iphone 11", 10)
	getProductWithVariant()
	// deleteVariantById(1)
}

func createProduct(name string) {
	db := database.GetDB()

	product := models.Product{
		Name: name,
	}

	err := db.Create(&product).Error

	if err != nil {
		fmt.Println("Error creating user data: ", err)
		return
	}

	fmt.Println("New product data: ", product)
}

func getProductById(id uint) {
	db := database.GetDB()

	product := models.Product{}

	err := db.First(&product, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("Product data not found")
			return
		}
		fmt.Print("Error finding product", product)
	}
	fmt.Printf("Product data : %+v \n", product)
}

func updateProductById(id int, name string) {
	db := database.GetDB()

	product := models.Product{}

	err := db.Model(&product).Where("id = ?", id).Updates(models.Product{Name: name}).Error

	if err != nil {
		fmt.Println("Error updating product data:", err)
		return
	}
	fmt.Printf("Update product's name: %+v \n", product.Name)
}

func createVariant(productId uint, variantName string, quantity int) {
	db := database.GetDB()

	Variant := models.Variants{
		ProductID:   productId,
		VariantName: variantName,
		Quantity:    quantity,
	}

	err := db.Create(&Variant).Error

	if err != nil {
		fmt.Println("Error creating variant data", err)
		return
	}
	fmt.Println("New Variant Data", Variant)
}

func getProductWithVariant() {
	db := database.GetDB()

	product := models.Product{}
	err := db.Preload("Variants").Find(&product).Error

	if err != nil {
		fmt.Println("Error getting data with books:", err.Error())
		return
	}

	fmt.Println("Product Datas With Variants")
	fmt.Printf("%+v", product)
}

func deleteVariantById(id uint) {
	db := database.GetDB()

	variant := models.Variants{}

	err := db.Where("id = ?", id).Delete(&variant).Error
	if err != nil {
		fmt.Println("Error deleting variant:", err.Error())
		return
	}
	fmt.Printf("Book with id %d has been successfully deleted", id)
}
