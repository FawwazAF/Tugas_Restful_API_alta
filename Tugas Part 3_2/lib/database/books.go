package database

import (
	"example/pawpi/config"
	"example/pawpi/models"
)

func GetBooks() (interface{}, error) {
	var books []models.Book
	if err := config.DB.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func CreateBooks(books models.Book) (interface{}, error) {

	if err := config.DB.Save(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}
func UpdateBooks(books models.Book, id int) (interface{}, error) {

	if err := config.DB.Find(&books, "id=?", id).Error; err != nil {
		return nil, err
	}

	if err := config.DB.Save(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func DeleteBooks(id int) (interface{}, error) {
	var books []models.Book
	if err := config.DB.Find(&books, "id=?", id).Error; err != nil {
		return nil, err
	}
	if err := config.DB.Delete(&books, "id=?", id).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func GetBook(id int) (interface{}, error) {
	var books models.Book

	if err := config.DB.Find(&books, "id=?", id).Error; err != nil {
		return nil, err
	}
	return books, nil
}
