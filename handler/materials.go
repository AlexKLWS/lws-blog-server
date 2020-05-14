package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/AlexKLWS/lws-blog-server/materials"
	"github.com/AlexKLWS/lws-blog-server/models"
	"github.com/labstack/echo"
)

func GetMaterials(c echo.Context) error {
	data := struct {
		FromDate string
		Category models.Category
	}{}

	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&data)
	var allMaterialRecords []models.MaterialRecord
	// Maybe refactor this bit later
	if err != nil {
		if err.Error() == "EOF" {
			allMaterialRecords = materials.Get()
		} else {
			log.Printf("Failed processing article request: %s\n", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	} else {
		if data.FromDate != "" {
			t, tError := time.Parse(time.RFC3339, data.FromDate)
			if tError != nil {
				log.Printf("Failed parsing from date: %s\n", err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			if data.Category != 0 {
				allMaterialRecords = materials.GetFromDateForCategory(t, data.Category)
			} else {
				allMaterialRecords = materials.GetFromDate(t)
			}
		} else {
			if data.Category != 0 {
				allMaterialRecords = materials.GetForCategory(data.Category)
			} else {
				allMaterialRecords = materials.Get()
			}
		}
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
