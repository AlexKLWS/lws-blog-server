package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AlexKLWS/lws-blog-server/articles"
	"github.com/AlexKLWS/lws-blog-server/models"
	"github.com/AlexKLWS/lws-blog-server/pageindex"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
)

func UpdateOrCreateArticle(c echo.Context) error {
	articleData := models.ArticleData{}

	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&articleData)
	if err != nil {
		log.Printf("Failed processing article submit request: %s\n", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	if articleData.ReferenceId == "" {
		u := uuid.Must(uuid.NewV4())
		articleData.ReferenceId = u.String()
	}
	go func() {
		articles.UpdateOrCreate(&articleData)
		if articleData.Category != models.Misc {
			pageindex.Update(models.Misc)
		}
		pageindex.Update(articleData.Category)
	}()

	return c.String(http.StatusOK, "OK")
}

func GetArticle(c echo.Context) error {
	id := c.QueryParam("id")

	article := articles.Get(id)

	return c.JSON(http.StatusOK, article)
}
