package materials

import (
	"fmt"
	"log"
	"time"

	"github.com/AlexKLWS/lws-blog-server/config"
	"github.com/AlexKLWS/lws-blog-server/models"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

func GetFromDate(date time.Time) []models.MaterialRecord {
	db, err := gorm.Open(viper.GetString(config.GormDialect), viper.GetString(config.GormConnectionString))
	if err != nil {
		log.Printf("DB open error: %s\n", err)
		panic("Failed to connect database")
	}
	defer db.Close()

	var materials []models.MaterialRecord
	var intermediateData []models.JoinedArticlePage
	db.Table(config.PagesTableName).
		Where(fmt.Sprintf("%s.created_at < ?", config.PagesTableName), date).
		Select("*").
		Joins(fmt.Sprintf("JOIN %s ON %s.icon_refer = %s.id", config.IconTableName, config.PagesTableName, config.IconTableName)).
		Find(&intermediateData)
	for i := range intermediateData {
		page := models.CreatePageDataFromJoinedRecord(intermediateData[i])
		materials = append(materials, page)
	}

	db.Table(config.ArticleTableName).
		Where(fmt.Sprintf("%s.created_at < ?", config.ArticleTableName), date).
		Select("*").
		Joins(fmt.Sprintf("JOIN %s ON %s.icon_refer = %s.id", config.IconTableName, config.ArticleTableName, config.IconTableName)).
		Find(&intermediateData)
	for i := range intermediateData {
		article := models.CreateArticleDataFromJoinedRecord(intermediateData[i])
		materials = append(materials, article)
	}

	return materials
}

func GetFromDateForCategory(date time.Time, category models.Category) []models.MaterialRecord {
	db, err := gorm.Open(viper.GetString(config.GormDialect), viper.GetString(config.GormConnectionString))
	if err != nil {
		log.Printf("DB open error: %s\n", err)
		panic("Failed to connect database")
	}
	defer db.Close()

	var materials []models.MaterialRecord
	var intermediateData []models.JoinedArticlePage
	db.Table(config.PagesTableName).
		Where(fmt.Sprintf("%s.created_at < ? AND %s.category = ?", config.PagesTableName, config.PagesTableName), date, category).
		Select("*").
		Joins(fmt.Sprintf("JOIN %s ON %s.icon_refer = %s.id", config.IconTableName, config.PagesTableName, config.IconTableName)).
		Find(&intermediateData)
	for i := range intermediateData {
		page := models.CreatePageDataFromJoinedRecord(intermediateData[i])
		materials = append(materials, page)
	}

	db.Table(config.ArticleTableName).
		Where(fmt.Sprintf("%s.created_at < ? AND %s.category = ?", config.ArticleTableName, config.ArticleTableName), date, category).
		Select("*").
		Joins(fmt.Sprintf("JOIN %s ON %s.icon_refer = %s.id", config.IconTableName, config.ArticleTableName, config.IconTableName)).
		Find(&intermediateData)
	for i := range intermediateData {
		article := models.CreateArticleDataFromJoinedRecord(intermediateData[i])
		materials = append(materials, article)
	}

	return materials
}

func Get() []models.MaterialRecord {
	db, err := gorm.Open(viper.GetString(config.GormDialect), viper.GetString(config.GormConnectionString))
	if err != nil {
		log.Printf("DB open error: %s\n", err)
		panic("Failed to connect database")
	}
	defer db.Close()

	var materials []models.MaterialRecord
	var intermediateData []models.JoinedArticlePage
	db.Table(config.PagesTableName).
		Select("*").
		Joins(fmt.Sprintf("JOIN %s ON %s.icon_refer = %s.id", config.IconTableName, config.PagesTableName, config.IconTableName)).
		Find(&intermediateData)
	for i := range intermediateData {
		page := models.CreatePageDataFromJoinedRecord(intermediateData[i])
		materials = append(materials, page)
	}

	db.Table(config.ArticleTableName).
		Select("*").
		Joins(fmt.Sprintf("JOIN %s ON %s.icon_refer = %s.id", config.IconTableName, config.ArticleTableName, config.IconTableName)).
		Find(&intermediateData)
	for i := range intermediateData {
		article := models.CreateArticleDataFromJoinedRecord(intermediateData[i])
		materials = append(materials, article)
	}

	return materials
}

func GetForCategory(category models.Category) []models.MaterialRecord {
	db, err := gorm.Open(viper.GetString(config.GormDialect), viper.GetString(config.GormConnectionString))
	if err != nil {
		log.Printf("DB open error: %s\n", err)
		panic("Failed to connect database")
	}
	defer db.Close()

	var materials []models.MaterialRecord
	var intermediateData []models.JoinedArticlePage
	db.Table(config.PagesTableName).
		Where(fmt.Sprintf("%s.category = ?", config.PagesTableName), category).
		Select("*").
		Joins(fmt.Sprintf("JOIN %s ON %s.icon_refer = %s.id", config.IconTableName, config.PagesTableName, config.IconTableName)).
		Find(&intermediateData)
	for i := range intermediateData {
		page := models.CreatePageDataFromJoinedRecord(intermediateData[i])
		materials = append(materials, page)
	}

	db.Table(config.ArticleTableName).
		Where(fmt.Sprintf("%s.category = ?", config.ArticleTableName), category).
		Select("*").
		Joins(fmt.Sprintf("JOIN %s ON %s.icon_refer = %s.id", config.IconTableName, config.ArticleTableName, config.IconTableName)).
		Find(&intermediateData)
	for i := range intermediateData {
		article := models.CreateArticleDataFromJoinedRecord(intermediateData[i])
		materials = append(materials, article)
	}

	return materials
}
