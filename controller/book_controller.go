package controller

import (
	"app-golang/config"
	"app-golang/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

var books []model.Book

func GetBooksController(c echo.Context) error {
	err := config.DB.Find(&books).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all books",
		"books":    books,
	})
}

func GetBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := config.DB.Find(&books, &model.Book{ID: id}).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get book with id " + c.Param("id"),
		"books":    &books,
	})
}

func CreateBookController(c echo.Context) error {
	book := model.Book{}
	c.Bind(&book)

	err := config.DB.Save(&book).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create new book",
		"user":     book,
	})
}

func UpdateBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	book := model.Book{}
	book.ID = id
	c.Bind(&book)

	err := config.DB.Save(&book).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update book with id " + c.Param("id"),
		"book":     book,
	})
}

func DeleteBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	book := model.Book{}

	err := config.DB.Delete(&book, id).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete book with id " + c.Param("id"),
	})
}
