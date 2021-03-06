package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AlexKLWS/lws-blog-server/models"
	"github.com/AlexKLWS/lws-blog-server/pageindex"
	"github.com/AlexKLWS/lws-blog-server/pages"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
)

func UpdateOrCreatePage(c echo.Context) error {
	pageData := models.PageData{}

	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&pageData)
	if err != nil {
		log.Printf("Failed processing page submit request: %s\n", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	if pageData.ReferenceId == "" {
		u := uuid.Must(uuid.NewV4())
		pageData.ReferenceId = u.String()
	}
	go func() {
		pages.UpdateOrCreate(&pageData)
		if pageData.Category != models.Misc {
			pageindex.Update(models.Misc)
		}
		pageindex.Update(pageData.Category)
	}()

	return c.String(http.StatusOK, "OK")
}

func GetPage(c echo.Context) error {
	id := c.QueryParam("id")

	article := pages.Get(id)

	return c.JSON(http.StatusOK, article)
}
