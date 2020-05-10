package handler

import (
	"encoding/json"
	"github.com/AlexKLWS/lws-blog-server/articles"
	"github.com/AlexKLWS/lws-blog-server/models"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
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

	u := uuid.Must(uuid.NewV4())
	articleData.ReferenceId = u.String()
	go articles.Create(&articleData)

	return c.String(http.StatusOK, "OK")
}

func GetArticle(c echo.Context) error {
	go articles.Get("")

	return c.String(http.StatusOK, "OK")
}