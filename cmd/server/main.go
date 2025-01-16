package main

import (
	"log"
	"net/http"

	"github.com/connorkuljis/backtrace/internal/model"
	"github.com/connorkuljis/backtrace/internal/renderer"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"

	"github.com/jmoiron/sqlx"
)

const (
	dbstr = "file:db/db.sqlite3"
)

func main() {
	db, err := sqlx.Open("sqlite3", dbstr)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Renderer = renderer.NewTemplateRenderer()
	e.Use(middleware.Logger())
	e.Static("/", "public")

	e.GET("/", handleIndex(db))
	e.GET("/search", handleSearch(db))
	e.GET("/company/:id", handleCompany(db))

	e.Start(":8080")
}

func handleIndex(db *sqlx.DB) echo.HandlerFunc {
	var count int
	err := db.Get(&count, "SELECT COUNT(*) FROM business_search")
	if err != nil {
		log.Fatal(err)
	}

	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "_index.html", map[string]interface{}{
			"Count": count,
		})
	}
}

func handleSearch(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO: input sanitisation
		queryStr := c.QueryParam("query")
		stateStr := c.QueryParam("state")

		var results []model.BusinessSearch
		if queryStr != "" {
			err := db.Select(&results, `SELECT * FROM business_search WHERE name MATCH ? AND state = ? ORDER BY abn DESC`, queryStr, stateStr)
			if err != nil {
				// TODO: don't throw error, but return error message to user.
				// 	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
		}

		return c.Render(http.StatusOK, "search-results", map[string]interface{}{
			"QueryStr": queryStr,
			"Count":    len(results),
			"Results":  results,
		})
	}
}
func handleCompany(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		var business model.Business
		err := db.Get(&business, `SELECT * FROM business_names WHERE BN_ABN = ?`, id)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.Render(http.StatusOK, "_company-details.html", map[string]interface{}{
			"Business": business,
		})
	}
}
