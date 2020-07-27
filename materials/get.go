package materials

import (
	"fmt"
	"log"

	"github.com/AlexKLWS/lws-blog-server/config"
	"github.com/AlexKLWS/lws-blog-server/models"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

func GetMaterialsPageForCategory(page models.PageIndex, category models.Category) []models.MaterialRecord {
	db, err := gorm.Open(viper.GetString(config.GormDialect), viper.GetString(config.GormConnectionString))
	if err != nil {
		log.Printf("DB open error: %s\n", err)
		panic("Failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&models.IconData{})
	db.AutoMigrate(&models.PageData{})
	db.AutoMigrate(&models.ArticleData{})
	db.AutoMigrate(&models.GuideData{})

	var materials []models.MaterialRecord
	var intermediateData []models.JoinedArticlePage
	if page.Page == 1 {
		if category != models.Misc {
			db.Table(config.PagesTableName).
				Where(fmt.Sprintf("%s.created_at >= ? AND %s.category = ?", config.PagesTableName, config.PagesTableName), page.EndDate, category).
				Select("*").
				Joins(fmt.Sprintf("JOIN %s ON %s.icon_refer = %s.id", config.IconTableName, config.PagesTableName, config.IconTableName)).
				Find(&intermediateData)
		} else {
			db.Table(config.PagesTableName).
				Where(fmt.Sprintf("%s.created_at >= ?", config.PagesTableName), page.EndDate).
				Select("*").
				Joins(fmt.Sprintf("JOIN %s ON %s.icon_refer = %s.id", config.IconTableName, config.PagesTableName, config.IconTableName)).
				Find(&intermediateData)
		}
	} else {
		if category != models.Misc {
			db.Table(config.PagesTableName).
				Where(fmt.Sprintf("%s.created_at <= ? AND %s.created_at >= ? AND %s.category = ?", config.PagesTableName, config.PagesTableName, config.PagesTableName), page.StartDate, page.EndDate, category).
				Select("*").
				Joins(fmt.Sprintf("JOIN %s ON %s.icon_refer = %s.id", config.IconTableName, config.PagesTableName, config.IconTableName)).
				Find(&intermediateData)
		} else {
			db.Table(config.PagesTableName).
				Where(fmt.Sprintf("%s.created_at <= ? AND %s.created_at >= ?", config.PagesTableName, config.PagesTableName), page.StartDate, page.EndDate).
				Select("*").
				Joins(fmt.Sprintf("JOIN %s ON %s.icon_refer = %s.id", config.IconTableName, config.PagesTableName, config.IconTableName)).
				Find(&intermediateData)
		}
	}
	for i := range intermediateData {
		page := models.CreatePageDataFromJoinedRecord(intermediateData[i])
		materials = append(materials, page)
	}

	if page.Page == 1 {
		if category != models.Misc {
			db.Table(config.ArticleTableName).
				Where(fmt.Sprintf("%s.created_at >= ? AND %s.category = ?", config.ArticleTableName, config.ArticleTableName), page.EndDate, category).
				Select("*").
				Joins(fmt.Sprintf("JOIN %s ON %s.icon_refer = %s.id", config.IconTableName, config.ArticleTableName, config.IconTableName)).
				Find(&intermediateData)
		} else {
			db.Table(config.ArticleTableName).
				Where(fmt.Sprintf("%s.created_at >= ?", config.ArticleTableName), page.EndDate).
				Select("*").
				Joins(fmt.Sprintf("JOIN %s ON %s.icon_refer = %s.id", config.IconTableName, config.ArticleTableName, config.IconTableName)).
				Find(&intermediateData)
		}
	} else {
		if category != models.Misc {
			db.Table(config.ArticleTableName).
				Where(fmt.Sprintf("%s.created_at <= ? AND %s.created_at >= ? AND %s.category = ?", config.ArticleTableName, config.ArticleTableName, config.ArticleTableName), page.StartDate, page.EndDate, category).
				Select("*").
				Joins(fmt.Sprintf("JOIN %s ON %s.icon_refer = %s.id", config.IconTableName, config.ArticleTableName, config.IconTableName)).
				Find(&intermediateData)
		} else {
			db.Table(config.ArticleTableName).
				Where(fmt.Sprintf("%s.created_at <= ? AND %s.created_at >= ?", config.ArticleTableName, config.ArticleTableName), page.StartDate, page.EndDate).
				Select("*").
				Joins(fmt.Sprintf("JOIN %s ON %s.icon_refer = %s.id", config.IconTableName, config.ArticleTableName, config.IconTableName)).
				Find(&intermediateData)
		}
	}
	for i := range intermediateData {
		article := models.CreateArticleDataFromJoinedRecord(intermediateData[i])
		materials = append(materials, article)
	}

	if page.Page == 1 {
		db.Table(config.GuidesTabelName).
			Where(fmt.Sprintf("%s.created_at >= ?", config.GuidesTabelName), page.EndDate).
			Select("*").
			Joins(fmt.Sprintf("JOIN %s ON %s.icon_refer = %s.id", config.IconTableName, config.GuidesTabelName, config.IconTableName)).
			Find(&intermediateData)
	} else {
		db.Table(config.GuidesTabelName).
			Where(fmt.Sprintf("%s.created_at <= ? AND %s.created_at >= ?", config.GuidesTabelName, config.GuidesTabelName), page.StartDate, page.EndDate).
			Select("*").
			Joins(fmt.Sprintf("JOIN %s ON %s.icon_refer = %s.id", config.IconTableName, config.GuidesTabelName, config.IconTableName)).
			Find(&intermediateData)
	}
	for i := range intermediateData {
		article := models.CreateArticleDataFromJoinedRecord(intermediateData[i])
		materials = append(materials, article)
	}

	return materials
}
