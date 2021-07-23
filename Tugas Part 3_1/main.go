package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

var (
	DB *gorm.DB
)

func init() {
	InitDB()
	InitialMigration()
}

type Config struct {
	DB_username string
	DB_password string
	DB_port     string
	DB_host     string
	DB_name     string
}

func InitDB() {
	config := Config{
		DB_username: "root",
		DB_password: "Minus12345",
		DB_port:     "3306",
		DB_host:     "localhost",
		DB_name:     "api_coba",
	}

	connectionString :=
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			config.DB_username,
			config.DB_password,
			config.DB_host,
			config.DB_port,
			config.DB_name,
		)

	var err error
	DB, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
}

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func InitialMigration() {
	DB.AutoMigrate(&User{})
}

// get all users
func GetUsersController(c echo.Context) error {
	users := []User{}
	if err := DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get users data",
		"user":    users,
	})
}

func GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	users := User{}

	if err := DB.Find(&users, "id=?", id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user data",
		"user":    users,
	})
}

func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	users := User{}
	if err := DB.Find(&users, "id=?", id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error)
	}
	if err := DB.Delete(&users, "id=?", id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user data",
		"user":    users,
	})

}

func UpdateUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	users := User{}
	if err := DB.Find(&users, "id=?", id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error)
	}
	c.Bind(&users)
	if err := DB.Save(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    users,
	})
}

func CreateUserController(c echo.Context) error {
	// binding data
	users := User{}
	c.Bind(&users)

	if err := DB.Save(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     users,
	})
}

func main() {
	// create a new echo instance
	e := echo.New()

	//Route / to handler function
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)

	//START
	e.Logger.Fatal(e.Start(":8000"))
}
