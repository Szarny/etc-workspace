package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type Comment struct {
	Id   int    `json:"id"`
	Name string `json:"name" form:"name"`
	Text string `json:"text" form:"text"`
}

func (c *Comment) String() string {
	return fmt.Sprintf("[ID %v] %v .oO{ %v }", c.Id, c.Name, c.Text)
}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		name := c.QueryParam("name")
		return c.String(http.StatusOK, "Hello, "+name)
	})

	e.GET("/greet/:name", func(c echo.Context) error {
		name := c.Param("name")
		return c.String(http.StatusOK, "Hello, "+name)
	})

	e.POST("/", func(c echo.Context) error {
		name := c.FormValue("name")
		return c.String(http.StatusOK, "Hello, "+name)
	})

	e.GET("/api/comments", func(c echo.Context) error {
		comments := []Comment{
			{Id: 1, Name: "Tsubasa", Text: "Hello"},
			{Id: 2, Name: "Someone", Text: "Hey!"},
		}
		return c.JSON(http.StatusOK, comments)
	})

	e.POST("/api/comments", func(c echo.Context) error {
		var comment Comment

		if err := c.Bind(&comment); err != nil {
			return c.String(http.StatusBadRequest, "Bind:"+err.Error())
		}

		return c.String(http.StatusOK, comment.String())
	})

	e.Logger.Fatal(e.Start(":8080"))
}
