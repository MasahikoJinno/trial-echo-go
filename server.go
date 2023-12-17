package main

import (
    "net/http"
    "strconv"

    "github.com/labstack/echo/v4"
)

type (
    User struct {
        ID   int    `json:"id"`
        Name string `json:"name"`
    }
    Error struct {
        Message string `json:"message"`
    }
)

var (
    users = []User{
        {ID: 1, Name: "Jon"},
        {ID: 2, Name: "Doe"},
    }
)

func main() {
    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })
    e.GET("/users", getUsers)
    e.GET("/users/:id", getUser)
    e.Logger.Fatal(e.Start(":8080"))
}

func getUser(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, Error{Message: "Invalid ID"})
    }

    for _, user := range users {
        if user.ID == id {
            return c.JSON(http.StatusOK, user)
        }
    }

    return c.JSON(http.StatusNotFound, Error{Message: "User not found"})
}

func getUsers(c echo.Context) error {
    return c.JSON(http.StatusOK, users)
}