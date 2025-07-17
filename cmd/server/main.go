package main

import (
	"github.com/connorkuljis/backtrace/internal/handler"
	"github.com/connorkuljis/backtrace/internal/renderer"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

const (
	dbstr = "file:database/business_names.sqlite3?journal_mode=WAL"
)

type App = handler.App

func main() {
	db, err := sqlx.Open("sqlite3", dbstr)
	if err != nil {
		log.Fatal(err)
	}

	app := App{
		DB: db,
	}

	err = app.DB.Get(&app.TotalBusinesses, "SELECT COUNT(*) FROM business_search")
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Renderer = renderer.NewTemplateRenderer()
	e.Use(middleware.Logger())
	e.Static("/", "public")

	e.GET("/", func(c echo.Context) error { return c.Redirect(http.StatusSeeOther, "/search") })
	e.GET("/search", handler.HandleSearch(app))
	e.GET("/search/business/:id", handler.HandleBusiness(app))

	e.Start(":8080")
}
