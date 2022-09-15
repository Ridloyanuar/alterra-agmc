package routes

import (
	"advance-golang/constant"
	"advance-golang/controllers"
	"advance-golang/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	middlewares.LogMiddlewares(e)
	//users endpoint
	// e.GET("/users", controllers.GetAllUsers)
	// e.GET("/users/:id", controllers.GetSpecificUser)
	// e.POST("users", controllers.AddNewUser)
	// e.PUT("users/:id", controllers.UpdateUser)
	// e.DELETE("users/:id", controllers.DeleteUser)

	//books endpoint
	// e.GET("books", controllers.GetAllBooks)
	// e.GET("/books/:id", controllers.GetBookByID)
	// e.POST("books", controllers.CreateBook)
	// e.PUT("books/:id", controllers.UpdateBook)
	// e.DELETE("books/:id", controllers.DeleteBook)

	//Public API (NO AUTHENTICATION NEEDED)
	e.POST("/login", controllers.LoginUser)

	//books endpoint
	e.GET("books", controllers.GetAllBooks)
	e.GET("/books/:id", controllers.GetBookByID)

	//users endpoint
	e.POST("users", controllers.AddNewUser)

	//Authetication Needed
	authenticationRoute := e.Group("/auth")
	authenticationRoute.Use(middleware.JWT([]byte(constant.SECRET_JWT)))
	authenticationRoute.GET("/users", controllers.GetAllUsers)       //get all users
	authenticationRoute.GET("/users/:id", controllers.Me)            //get profile
	authenticationRoute.PUT("/users/:id", controllers.UpdateUser)    //update user
	authenticationRoute.DELETE("/users/:id", controllers.DeleteUser) //delete user

	authenticationRoute.POST("books", controllers.CreateBook)       //create book
	authenticationRoute.PUT("books/:id", controllers.UpdateBook)    //update book
	authenticationRoute.DELETE("books/:id", controllers.DeleteBook) //delete book

	return e
}
