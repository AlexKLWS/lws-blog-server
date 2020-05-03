package handler

import (
	"github.com/AlexKLWS/lws-blog-server/materials"
	"github.com/AlexKLWS/lws-blog-server/models"
	"github.com/labstack/echo"
	"net/http"
	"sort"
)

func GetMaterials(c echo.Context) error {
	allMaterialRecords := materials.Get()
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
