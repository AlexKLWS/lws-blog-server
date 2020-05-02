package materials

import (
	"github.com/AlexKLWS/lws-blog-server/config"
	"github.com/AlexKLWS/lws-blog-server/models"
	"github.com/davecgh/go-spew/spew"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"log"
	"time"
)

func GetFromDate(date time.Time)  {
	db, err := gorm.Open(viper.GetString(config.GormDialect), viper.GetString(config.GormConnectionString))
	if err != nil {
		log.Printf("DB open error: %s\n", err)
		panic("Failed to connect database")
	}
	defer db.Close()

	var articles []models.ArticleData

	db.Where("created_at > ?", date).Order("created_at").Limit(20).Find(&articles)
	spew.Dump(articles)
}

func Get() []models.MaterialRecord {
	db, err := gorm.Open(viper.GetString(config.GormDialect), viper.GetString(config.GormConnectionString))
	if err != nil {
		log.Printf("DB open error: %s\n", err)
		panic("Failed to connect database")
	}
	defer db.Close()

	var materials []models.MaterialRecord
	var pages []models.PageData
	db.Order("created_at DESC").Limit(20).Find(&pages)
	for i := 0; i < len(pages); i++ {
		materials = append(materials, pages[i])
	}
	var articles []models.ArticleData
	db.Order("created_at DESC").Limit(20).Find(&articles)
	for i := 0; i < len(articles); i++ {
		materials = append(materials, articles[i])
	}

	return materials
}