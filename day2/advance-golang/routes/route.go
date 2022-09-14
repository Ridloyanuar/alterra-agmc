package routes

import (
	"advance-golang/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	//users endpoint
	e.GET("/users", controllers.GetAllUsers)
	e.GET("/users/:id", controllers.GetSpecificUser)
	e.POST("users", controllers.AddNewUser)
	e.PUT("users/:id", controllers.UpdateUser)
	e.DELETE("users/:id", controllers.DeleteUser)

	//books endpoint
	e.GET("books", controllers.GetAllBooks)
	e.GET("/books/:id", controllers.GetBookByID)
	e.POST("books", controllers.CreateBook)
	e.PUT("books/:id", controllers.UpdateBook)
	e.DELETE("books/:id", controllers.DeleteBook)

	return e
}
