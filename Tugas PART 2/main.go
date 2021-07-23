package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

func GetUsersController(c echo.Context) error {
	fmt.Println("get GET Req")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

func GetUserController(c echo.Context) error {
	userId := c.Param("id")
	j, _ := strconv.Atoi(userId)

	return c.JSON(http.StatusOK, map[string]string{
		"user": userId,
		"name": users[j-1].Name,
	})
}

func DeleteUserController(c echo.Context) error {
	userId := c.Param("id")
	j, _ := strconv.Atoi(userId)

	for i := j; i < len(users); i++ {
		users[i].Id -= 1
	}

	newUsers := append(users[:j-1], users[j:]...)
	users = newUsers
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "delete success",
	})
}

func UpdateUserController(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	userId := c.Param("id")
	id, _ := strconv.Atoi(userId)
	if u.Name != "" {
		users[id].Name = u.Name
	}
	if u.Email != "" {
		users[id].Email = u.Email
	}
	if u.Password != "" {
		users[id].Password = u.Password
	}

	return c.JSON(http.StatusOK, users[id])
}

func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	fmt.Println("get POST Req")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

func main() {
	e := echo.New()
	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)
	e.GET("/users/:id", GetUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)

	//Start Server, log if fail
	e.Logger.Fatal(e.Start(":1323"))
}
