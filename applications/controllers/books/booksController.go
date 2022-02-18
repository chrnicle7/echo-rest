package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/naomihc/echo-rest/applications/models"
	"github.com/naomihc/echo-rest/database"
)

type BooksJson struct {
	BookId     int
	Title      string
	Year       int
	AuthorName string
}

type BooksResponse struct {
	Status int
	Error  bool
	Books  []BooksJson
}

func GetBooks(c echo.Context) error {
	db := database.GetDBInstance()

	var scannedBooks []BooksJson
	db.Model(&models.Book{}).Select("books.id as book_id, books.title, books.year, books.author_id, authors.id, authors.name as author_name").
		Joins("left join authors on authors.id = books.author_id").
		Scan(&scannedBooks)

	return c.JSON(http.StatusOK, scannedBooks)
}

func GetBook(c echo.Context) error {
	db := database.GetDBInstance()
	reqId, _ := strconv.Atoi(c.Param("id"))

	var scannedBook BooksJson
	db.Model(&models.Book{}).Select("books.id as book_id, books.title, books.year, books.author_id, authors.id, authors.name as author_name").
		Where("books.id = ?", reqId).
		Joins("left join authors on authors.id = books.author_id").
		Scan(&scannedBook)

	return c.JSON(http.StatusOK, scannedBook)
}

func UpdateBook(c echo.Context) error {
	db := database.GetDBInstance()
	reqId, _ := strconv.Atoi(c.Param("id"))

	var bookFound models.Book
	db.First(&bookFound, reqId)

	jsonMap := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonMap)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "Check your form again")
	}

	db.Model(&bookFound).Where("id = ?", reqId).Updates(jsonMap)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "There's something wrong")
	}

	return c.JSON(http.StatusOK, "Book successfully updated")
}

func CreateBook(c echo.Context) error {
	db := database.GetDBInstance()

	jsonMap := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonMap)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "Check your form again")
	}

	db.Model(&models.Book{}).Create(jsonMap)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "There's something wrong")
	}

	return c.JSON(http.StatusOK, "Book successfully created")
}

func DeleteBook(c echo.Context) error {
	db := database.GetDBInstance()
	reqId, _ := strconv.Atoi(c.Param("id"))

	var bookFound models.Book
	db.Where("id = ?", reqId).Delete(&bookFound)

	return c.JSON(http.StatusOK, "Book successfully deleted")
}
