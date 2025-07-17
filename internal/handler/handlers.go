package handler

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/connorkuljis/backtrace/internal/model"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type App struct {
	TotalBusinesses int
	DB              *sqlx.DB
}

type SearchReq struct {
	Query  string
	State  string
	Limit  int64
	Offset int64
}

type SearchResp struct {
	TotalResults int64
	Message      string
	Paginator    string
	Results      []model.BusinessSearch
}

func HandleSearch(app App) echo.HandlerFunc {
	return func(c echo.Context) error {
		printer := message.NewPrinter(language.English)

		req := SearchReq{
			Query:  c.QueryParam("q"),
			State:  c.QueryParam("state"),
			Limit:  30, // default to show 30 entries
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

		resp := SearchResp{
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
			resp.Message = printer.Sprintf("Found %d results for '%s'", resp.TotalResults, req.Query)
			resp.Paginator = fmt.Sprintf("%s?q=%s&state=%s&limit=%d&offset=%d", c.Path(), req.Query, req.State, req.Limit, req.Offset+req.Limit)
		}

		if c.Request().Header.Get("X-Alpine-Request") == "true" {
			return c.Render(http.StatusOK, "_index-partial.html", resp)
		}

		return c.Render(http.StatusOK, "_index.html", resp)
	}
}

func HandleBusiness(app App) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		var business model.Business
		err := app.DB.Get(&business, `SELECT * FROM business_names WHERE abn = ?`, id)
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
