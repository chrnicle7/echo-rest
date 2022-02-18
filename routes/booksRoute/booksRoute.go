package booksRoute

import (
	"github.com/labstack/echo/v4"
	controllers "github.com/naomihc/echo-rest/applications/controllers/books"
)

func BooksSubRoute(group *echo.Group) {
	group.GET("", controllers.GetBooks)
	group.POST("", controllers.CreateBook)
	group.GET("/:id", controllers.GetBook)
	group.PUT("/:id", controllers.UpdateBook)
	group.DELETE("/:id", controllers.DeleteBook)
}
