package handler

import (
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/AlexKLWS/lws-blog-server/materials"
	"github.com/AlexKLWS/lws-blog-server/models"
	"github.com/labstack/echo"
)

func GetMaterials(c echo.Context) error {
	categoryNumber, conversionError := strconv.Atoi(c.QueryParam("category"))
	if conversionError != nil {
		log.Printf("Failed parsing category: %s\n", conversionError)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	category := models.Category(categoryNumber)
	fromDate := c.QueryParam("fromDate")

	var allMaterialRecords []models.MaterialRecord
	// Maybe refactor this bit later
	if fromDate != "" {
		t, tError := time.Parse(time.RFC3339, fromDate)
		if tError != nil {
			log.Printf("Failed parsing from date: %s\n", tError)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		if category != 0 {
			allMaterialRecords = materials.GetFromDateForCategory(t, category)
		} else {
			allMaterialRecords = materials.GetFromDate(t)
		}
	} else {
		if category != 0 {
			allMaterialRecords = materials.GetForCategory(category)
		} else {
			allMaterialRecords = materials.Get()
		}
	}

	sort.Slice(allMaterialRecords, func(i, j int) bool {
		return allMaterialRecords[i].GetCreatedAt().After(allMaterialRecords[j].GetCreatedAt())
	})
	var selectedMaterialRecords []models.MaterialRecord
	if len(allMaterialRecords) > 16 {
		selectedMaterialRecords = allMaterialRecords[:16]
	} else {
		selectedMaterialRecords = allMaterialRecords
	}
	return c.JSON(http.StatusOK, selectedMaterialRecords)
}
