package handler

import (
	"encoding/json"
	"github.com/AlexKLWS/lws-blog-server/materials"
	"github.com/AlexKLWS/lws-blog-server/models"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"sort"
	"time"
)

func GetMaterials(c echo.Context) error {
	data := struct{ FromDate string }{}

	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&data)
	var allMaterialRecords []models.MaterialRecord
	if err != nil {
		if err.Error() == "EOF" {
			allMaterialRecords = materials.Get()
		} else {
			log.Printf("Failed processing article request: %s\n", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	} else {
		t, tError := time.Parse(time.RFC3339, data.FromDate)
		if tError != nil {
			log.Printf("Failed parsing from date: %s\n", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		allMaterialRecords = materials.GetFromDate(t)
	}

	sort.Slice(allMaterialRecords, func(i, j int) bool {
		return allMaterialRecords[i].GetCreatedAt().After(allMaterialRecords[j].GetCreatedAt())
	})
	var selectedMaterialRecords []models.MaterialRecord
	if len(allMaterialRecords) > 20 {
		selectedMaterialRecords = allMaterialRecords[:20]
	} else {
		selectedMaterialRecords = allMaterialRecords
	}
	return c.JSON(http.StatusOK, selectedMaterialRecords)
}
