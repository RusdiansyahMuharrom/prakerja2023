package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type LoginRequest struct {
	Email string `json:"email"`
	Pass  string `json:"pass"`
}

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var data = []User{
	{"alta", 10},
	{"academy", 11},
}

func main() {
	e := echo.New()
	e.GET("/users", GetUsers)
	e.GET("/users/:id", GetDeatailUsers)
	e.GET("/users/params", GetDeatailUsersParamsQuery)
	e.POST("/login/form", LoginForm)
	e.POST("/login/json", LoginJson)
	e.Logger.Fatal(e.Start(":3002"))
}

func GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, Response{
		Message: "succes",
		Data:    data,
	})
}

func GetDeatailUsers(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, Response{
		Message: "succes",
		Data:    id,
	})
}

func GetDeatailUsersParamsQuery(c echo.Context) error {
	page := c.QueryParam("page")
	limit := c.QueryParam("limit")
	return c.JSON(http.StatusOK, Response{
		Message: "succes",
		Data:    page + " " + limit,
	})
}

func LoginForm(c echo.Context) error {
	email := c.FormValue("email")
	pass := c.FormValue("pass")
	return c.JSON(http.StatusOK, Response{
		Message: "succes",
		Data:    email + " " + pass,
	})
}

func LoginJson(c echo.Context) error {
	var loginrequest LoginRequest
	c.Bind(&loginrequest)
	return c.JSON(http.StatusOK, Response{
		Message: "succes",
		Data:    loginrequest.Email + " " + loginrequest.Pass,
	})
}
