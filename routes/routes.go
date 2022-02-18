package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/naomihc/echo-rest/routes/booksRoute"
)

func Init() *echo.Echo {
	e := echo.New()

	booksGroup := e.Group("/books")
	booksRoute.BooksSubRoute(booksGroup)

	return e
}
