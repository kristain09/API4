package main

import (
	bHandler "github.com/kristain09/API4/app/feature/book/handler"
	bRepo "github.com/kristain09/API4/app/feature/book/repository"
	bLogic "github.com/kristain09/API4/app/feature/book/usecase"
	uHandler "github.com/kristain09/API4/app/feature/user/handler"
	uRepo "github.com/kristain09/API4/app/feature/user/repository"
	uLogic "github.com/kristain09/API4/app/feature/user/usecase"
	"github.com/kristain09/API4/config"
	"github.com/kristain09/API4/routes"
	"github.com/kristain09/API4/utils/database"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	cfg := config.IniConfig()
	conn := database.InitDB(*cfg)
	database.Migrate(conn)

	uMdl := uRepo.New(conn)
	uSrv := uLogic.New(uMdl)
	uCtl := uHandler.New(uSrv)

	bookMdl := bRepo.New(conn)
	bookSrv := bLogic.New(bookMdl)
	bookCtl := bHandler.New(bookSrv)

	// ROUTING
	routes.Route(e, uCtl, bookCtl)

	if err := e.Start(":8000"); err != nil {
		e.Logger.Fatal("cannot start server", err.Error())
	}
}
