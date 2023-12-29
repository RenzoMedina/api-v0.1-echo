package main

import (
	"echoV1/model"
	"echoV1/pkg/post"
	"echoV1/storage"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	storage.ConnectionDB()
	//storage.DB().AutoMigrate(&model.Post{})

	e := echo.New()
	e.GET("/v1", func(c echo.Context) error {
		return c.JSON(http.StatusOK, post.GetAll(storage.DB()))
	})

	e.POST("/v1", func(c echo.Context) error {

		code := c.FormValue("code")
		name := c.FormValue("name")
		description := c.FormValue("description")
		stock := c.FormValue("stock")

		posts := &model.Post{Code: code, Name: name, Description: description, Stock: stock}

		result := post.CreatePost(storage.DB(), posts)
		return c.JSON(http.StatusOK, result)
	})

	e.GET("/v1/:id", func(c echo.Context) error {

		id := c.Param("id")
		return c.JSON(http.StatusOK, post.GetById(storage.DB(), id))
	})

	e.PUT("/v1/:id", func(c echo.Context) error {

		id := c.Param("id")
		name := c.FormValue("name")
		description := c.FormValue("description")

		posts := &model.Post{Name: name, Description: description}
		return c.JSON(http.StatusOK, post.UpdateById(storage.DB(), id, *posts))
	})
	e.DELETE("/v1/:id", func(c echo.Context) error {

		id := c.Param("id")
		post.DeleteById(storage.DB(), id)
		return c.JSON(http.StatusOK, "Post has been delete!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
