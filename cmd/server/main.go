package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"text/template"

	"github.com/connorkuljis/backtrace/internal/model"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"

	"github.com/jmoiron/sqlx"
)

const (
	dbstr        = "file:data/db.sqlite3"
	randomABN    = "88573118334"
	dbContextKey = "_db"
)

var funcMap = template.FuncMap{}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// returns a new template set.
func templates() *template.Template {
	return template.Must(template.New("").Funcs(funcMap).Option("missingkey=error").ParseGlob("templates/*.html"))
}

func dbMiddleware(db *sqlx.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(dbContextKey, db)
			return next(c)
		}
	}
}

func main() {
	e := echo.New()

	// html template renderer
	t := &Template{
		templates: templates(),
	}
	e.Renderer = t

	// db connection middleware
	db, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	e.Use(dbMiddleware(db))
	e.Use(middleware.Logger())
	// e.Static("/", "public")

	// Using Get
	var count int
	err = db.Get(&count, "SELECT COUNT(*) FROM business_names")
	if err != nil {
		log.Fatal(err)
	}

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{
			"Count": count,
		})
	})

	e.GET("/search", func(c echo.Context) error {
		db := c.Get(dbContextKey).(*sqlx.DB)

		queryStr := c.QueryParam("query")

		var results []model.BusinessSearch
		if queryStr != "" {
			err = db.Select(&results, `SELECT * FROM businesses WHERE business_name MATCH ? ORDER BY abn DESC`, queryStr)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
		}

		return c.Render(http.StatusOK, "search-results", map[string]interface{}{
			"QueryStr": queryStr,
			"Count":    len(results),
			"Results":  results,
		})
	})

	e.Start(":8080")
}

func connect() (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", dbstr)
	if err != nil {
		return nil, err
	}

	fmt.Println("connected to ", dbstr)

	return db, nil
}
