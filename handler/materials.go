package handler

import (
	"log"
	"net/http"
	"sort"
	"strconv"

	"github.com/AlexKLWS/lws-blog-server/config"
	"github.com/AlexKLWS/lws-blog-server/materials"
	"github.com/AlexKLWS/lws-blog-server/models"
	"github.com/AlexKLWS/lws-blog-server/pageindex"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

type MaterilsResponse struct {
	Materials []models.MaterialRecord `json:"materials" xml:"materials"`
	PageCount int                     `json:"pageCount" xml:"pageCount"`
}

func GetMaterials(c echo.Context) error {
	categoryNumber, conversionError := strconv.Atoi(c.QueryParam("category"))
	if conversionError != nil {
		log.Printf("Failed parsing category: %s\n", conversionError)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	category := models.Category(categoryNumber)
	page := c.QueryParam("page")
	pageNumber := 1
	if page != "" {
		pageNumber, conversionError = strconv.Atoi(c.QueryParam("page"))
		if conversionError != nil {
			log.Printf("Failed parsing category: %s\n", conversionError)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	}
	pageIndex := pageindex.Get(pageNumber, category)
	pageCount := pageindex.GetPagesCount(category)

	var allMaterialRecords []models.MaterialRecord
	allMaterialRecords = materials.GetMaterialsPageForCategory(pageIndex, category)

	sort.Slice(allMaterialRecords, func(i, j int) bool {
		return allMaterialRecords[i].GetCreatedAt().After(allMaterialRecords[j].GetCreatedAt())
	})
	var selectedMaterialRecords []models.MaterialRecord
	if len(allMaterialRecords) > viper.GetInt(config.PageSize) {
		selectedMaterialRecords = allMaterialRecords[:viper.GetInt(config.PageSize)]
	} else {
		selectedMaterialRecords = allMaterialRecords
	}

	response := MaterilsResponse{
		Materials: selectedMaterialRecords,
		PageCount: pageCount,
	}

	return c.JSON(http.StatusOK, response)
}
