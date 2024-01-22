package models

import (
	"go-api/src/config"

	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name  string `gorm:"type:varchar(255)"`
	Price int
	Stock int
}

func SelectAll() *gorm.DB{
	items := []Product{}
	return config.DB.Find(&items)
}

func Create(newProduct *Product) *gorm.DB  {
	return config.DB.Create(&newProduct)
}

func Select(id string) *gorm.DB {
	var item Product
	return config.DB.First(&item, "id = ?", id)
}

func Updates(id string, updateProduct *Product) *gorm.DB {
	var item Product
	return config.DB.Model(&item).Where("id = ?", id).Updates(&updateProduct)
}

func Deletes(id string) *gorm.DB {
	var item Product
	return config.DB.Delete(&item, "id = ?", id)
}