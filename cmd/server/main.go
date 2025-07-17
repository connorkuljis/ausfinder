package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

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
	dbstr = "file:database/business_names_202503.sqlite3?journal_mode=WAL"
)

type App struct {
	TotalBusinesses int
	DB              *sqlx.DB
}

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
	e.GET("/search", handleSearch(app))
	e.GET("/search/business/:id", handleBusiness(db))

	e.Start(":8080")
}

func handleSearch(app App) echo.HandlerFunc {
	type req struct {
		Query string
		State string

		Limit  int64
		Offset int64
	}

	type resp struct {
		ShownResults int64
		TotalResults int64
		Message      string
		Paginator    string

		Results []model.BusinessSearch
	}

	return func(c echo.Context) error {
		printer := message.NewPrinter(language.English)

		req := req{
			Query:  c.QueryParam("q"),
			State:  c.QueryParam("state"),
			Limit:  30,
			Offset: 0,
		}

		if c.QueryParam("limit") != "" {
			limit, err := strconv.ParseInt(c.QueryParam("limit"), 10, 64)
			if err != nil {
				return err
			}
			req.Limit = limit
		}

		if c.QueryParam("offset") != "" {
			offset, err := strconv.ParseInt(c.QueryParam("offset"), 10, 64)
			if err != nil {
				return err
			}
			req.Offset = offset
		}

		resp := resp{
			Message: printer.Sprintf("Search %d active business names", app.TotalBusinesses),
		}

		if req.Query != "" {
			query := `SELECT COUNT(*) FROM business_search WHERE name MATCH ?`
			params := []interface{}{req.Query}

			if req.State != "" {
				query += ` AND state = ?`
				params = append(params, req.State)
			}

			var count int64
			app.DB.Get(&count, query, params...)
			resp.TotalResults = count
		}

		if req.Query != "" {
			query := `SELECT * FROM business_search WHERE name MATCH ?`
			params := []interface{}{req.Query}

			if req.State != "" {
				query += ` AND state = ?`
				params = append(params, req.State)
			}

			query += `ORDER BY abn DESC`

			query += ` LIMIT ?`
			params = append(params, req.Limit)

			query += ` OFFSET ?`
			params = append(params, req.Offset)

			app.DB.Select(&resp.Results, query, params...)
			resp.ShownResults = int64(len(resp.Results)) + req.Offset
			resp.Message = printer.Sprintf("Showing %d of %d results for '%s'", resp.ShownResults, resp.TotalResults, req.Query)
			resp.Paginator = fmt.Sprintf("%s?q=%s&state=%s&limit=%d&offset=%d", c.Path(), req.Query, req.State, req.Limit, req.Offset+req.Limit)

		}

		if c.Request().Header.Get("X-Alpine-Request") == "true" {
			return c.Render(http.StatusOK, "_index-partial.html", resp)
		}

		return c.Render(http.StatusOK, "_index.html", resp)
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
