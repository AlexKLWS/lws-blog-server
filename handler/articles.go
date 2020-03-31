package handler

import (
	"encoding/json"
	"github.com/AlexKLWS/lws-blog-server/models"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func NewArticle(c echo.Context) error {
	articleData := models.ArticleData{}

	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&articleData)
	if err != nil {
		log.Printf("Failed processing article submit request: %s\n", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.String(http.StatusOK, "OK")
}