package pageindex

import (
	"log"
	"sort"

	"github.com/AlexKLWS/lws-blog-server/config"
	"github.com/AlexKLWS/lws-blog-server/models"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

func Update(category models.Category) {
	db, err := gorm.Open(viper.GetString(config.GormDialect), viper.GetString(config.GormConnectionString))
	if err != nil {
		log.Printf("DB open error: %s\n", err)
		return
	}

	// Migrate the schema
	db.AutoMigrate(&models.PageIndex{})

	var materials []models.MaterialRecord
	var intermediateData []models.Model
	if category != models.Misc {
		db.Table(config.PagesTableName).
			Where("category = ?", category).
			Select("id, created_at").
			Find(&intermediateData)
	} else {
		db.Table(config.PagesTableName).
			Select("id, created_at").
			Find(&intermediateData)
	}
	for i := range intermediateData {
		materials = append(materials, intermediateData[i])
	}

	db.Table(config.ArticleTableName).
		Select("id, created_at").
		Find(&intermediateData)
	for i := range intermediateData {
		materials = append(materials, intermediateData[i])
	}

	sort.Slice(materials, func(i, j int) bool {
		return materials[i].GetCreatedAt().After(materials[j].GetCreatedAt())
	})

	pageSize := viper.GetInt(config.PageSize)
	if len(materials) > pageSize {
		p := 1
		db.Where(models.PageIndex{Page: p}).Assign(models.PageIndex{Category: category, EndDate: materials[pageSize-1].GetCreatedAt()}).FirstOrCreate(&models.PageIndex{})
		for i := pageSize; i < len(materials); i = i + pageSize {
			p++
			var endDateItemIndex int
			if i+pageSize-1 < len(materials) {
				endDateItemIndex = i + pageSize - 1
			} else {
				endDateItemIndex = len(materials) - 1
			}
			db.Where(models.PageIndex{Page: p}).Assign(models.PageIndex{Category: category, StartDate: materials[i].GetCreatedAt(), EndDate: materials[endDateItemIndex].GetCreatedAt()}).FirstOrCreate(&models.PageIndex{})
		}
	} else {
		i := len(materials) - 1
		db.Where(models.PageIndex{Page: 1}).Assign(models.PageIndex{Category: category, EndDate: materials[i].GetCreatedAt()}).FirstOrCreate(&models.PageIndex{})
	}
}
