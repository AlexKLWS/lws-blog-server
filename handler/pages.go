package handler

import (
	"encoding/json"
	"github.com/AlexKLWS/lws-blog-server/models"
	"github.com/AlexKLWS/lws-blog-server/pages"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func NewPage(c echo.Context) error {
	pageData := models.PageData{}

	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&pageData)
	if err != nil {
		log.Printf("Failed processing page submit request: %s\n", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	go pages.Create(&pageData)

	return c.String(http.StatusOK, "OK")
}

func GetPages(c echo.Context) error {
	p := pages.Get()
	return c.JSON(http.StatusOK, p)
}