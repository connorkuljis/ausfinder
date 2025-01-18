package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/connorkuljis/backtrace/internal/model"
	"github.com/connorkuljis/backtrace/internal/renderer"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/text/language"
	"golang.org/x/text/message"

	"github.com/jmoiron/sqlx"
)

const (
	dbstr = "file:db/db.sqlite3?journal_mode=WAL"
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

	e.GET("/", func(c echo.Context) error { return c.Redirect(http.StatusSeeOther, "/search") })
	e.GET("/search", handleSearch(db))
	e.GET("/search/business/:id", handleBusiness(db))

	e.Start(":8080")
}

func handleSearch(db *sqlx.DB) echo.HandlerFunc {
	printer := message.NewPrinter(language.English)

	var total int
	err := db.Get(&total, "SELECT COUNT(*) FROM business_search")
	if err != nil {
		log.Fatal(err)
	}

	return func(c echo.Context) error {
		var results []model.BusinessSearch
		var msg = printer.Sprintf("Search %d active business names", total)

		queryStr := c.QueryParam("q")
		stateStr := c.QueryParam("state")

		if queryStr != "" {
			query := `SELECT * FROM business_search WHERE name MATCH ?`
			params := []interface{}{queryStr}

			if stateStr != "" {
				query += ` AND state = ?`
				params = append(params, stateStr)
			}

			query += `ORDER BY abn DESC`

			err := db.Select(&results, query, params...)
			if err != nil {
				msg = fmt.Sprintf("An issue occurred searching for '%s'", queryStr)
			} else {
				msg = printer.Sprintf("Found (%d) results for '%s'", len(results), queryStr)
			}
		}

		data := map[string]any{
			"BusinessSearchResults": results,
			"Message":               msg,
		}

		if c.Request().Header.Get("X-Alpine-Request") == "true" {
			return c.Render(http.StatusOK, "search-results", data)
		}

		return c.Render(http.StatusOK, "_index.html", data)
	}
}

func handleBusiness(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		var business model.Business
		err := db.Get(&business, `SELECT * FROM business_names WHERE abn = ?`, id)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		// Escape the business name for use in the URL
		escapedName := url.QueryEscape(business.Name)

		// Create the LinkedIn URL with the escaped parameter
		linkedinURL := fmt.Sprintf("https://www.linkedin.com/search/results/companies/?keywords=%s", escapedName)

		return c.Render(http.StatusOK, "_business.html", map[string]interface{}{
			"Business":    business,
			"LinkedinURL": linkedinURL,
		})
	}
}
