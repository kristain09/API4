package routes

import (
	"github.com/kristain09/API4/app/feature/book"
	"github.com/kristain09/API4/app/feature/user"
	"github.com/kristain09/API4/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Route(e *echo.Echo, uc user.Handler, bc book.Handler) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.POST("/login", uc.Login())
	e.POST("/users", uc.Register())
	// tidak boleh ada, sisa /users
	e.GET("/users", uc.Profile(), middleware.JWT([]byte(config.Secret)))
	e.PUT("/users", uc.UpdateProfile(), middleware.JWT([]byte(config.Secret)))

	e.GET("/books", bc.GetAllBook())
	e.POST("/books", bc.AddBook(), middleware.JWT([]byte(config.Secret)))
	e.PUT("/books", bc.UpdateBook(), middleware.JWT([]byte(config.Secret)))
	e.GET("/books", bc.GetAllBook())
}
