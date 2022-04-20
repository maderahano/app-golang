package route

import (
	"app-golang/constants"
	"app-golang/controller"
	m "app-golang/middleware"

	"github.com/labstack/echo"
	mid "github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	eAuthJWT := e.Group("")
	eAuthJWT.Use(mid.JWT([]byte(constants.SECRET_JWT)))

	// Login
	e.POST("/login", controller.LoginUserController)

	// User Route
	eAuthJWT.GET("/users", controller.GetUsersController)
	eAuthJWT.GET("/users/:id", controller.GetUserController)
	e.POST("/users", controller.CreateUserController)
	eAuthJWT.PUT("/users/:id", controller.UpdateUserController)
	eAuthJWT.DELETE("/users/:id", controller.DeleteUserController)

	// Book Route
	e.GET("/books", controller.GetBooksController)
	e.GET("/books/:id", controller.GetBookController)
	eAuthJWT.POST("/books", controller.CreateBookController)
	eAuthJWT.PUT("/books/:id", controller.UpdateBookController)
	eAuthJWT.DELETE("/books/:id", controller.DeleteBookController)

	m.LogMiddleware(e)
	m.LogMiddlewares(eAuthJWT)

	return e
}
