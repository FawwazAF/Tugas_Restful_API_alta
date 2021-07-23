package database

import (
	"example/pawpi/config"
	"example/pawpi/middlewares"
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

func GetDetailAdmin(userId int) (interface{}, error) {
	var user models.User
	if err := config.DB.Find(&user, userId).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func LoginAdmin(email, password string) (interface{}, error) {
	var user models.User
	var err error
	if err = config.DB.Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		return nil, err
	}
	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}
	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, err
}
