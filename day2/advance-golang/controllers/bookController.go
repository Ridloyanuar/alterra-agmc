package controllers

import (
	"advance-golang/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var (
	books = map[int]*models.Books{}
	seq   = 1
)

func CreateBook(c echo.Context) error {
	b := &models.Books{
		ID: seq,
	}
	if err := c.Bind(b); err != nil {
		return err
	}
	books[b.ID] = b
	seq++
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status": "success created books",
		"books":  b,
	})
}

func GetAllBooks(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success get all books",
		"books":  books,
	})
}

func GetBookByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success get book by id",
		"books":  books[id],
	})
}

func UpdateBook(c echo.Context) error {
	b := new(models.Books)
	if err := c.Bind(b); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	books[id].Title = b.Title
	return c.JSON(http.StatusOK, books[id])
}

func DeleteBook(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(books, id)
	return c.NoContent(http.StatusNoContent)
}
