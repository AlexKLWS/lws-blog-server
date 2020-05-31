package pageindex

import (
	"fmt"
	"log"

	"github.com/AlexKLWS/lws-blog-server/config"
	"github.com/AlexKLWS/lws-blog-server/models"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

func Get(pageNumber int, category models.Category) models.PageIndex {
	db, err := gorm.Open(viper.GetString(config.GormDialect), viper.GetString(config.GormConnectionString))
	if err != nil {
		log.Printf("DB open error: %s\n", err)
		panic("Failed to connect database")
	}
	defer db.Close()

	var pageIndex models.PageIndex

	db.Table(config.PageIndexTableName).
		Where(fmt.Sprintf("%s.category = ? AND %s.page = ?", config.PageIndexTableName, config.PageIndexTableName), category, pageNumber).
		Find(&pageIndex)

	return pageIndex
}

func GetPagesCount(category models.Category) int {
	db, err := gorm.Open(viper.GetString(config.GormDialect), viper.GetString(config.GormConnectionString))
	if err != nil {
		log.Printf("DB open error: %s\n", err)
		panic("Failed to connect database")
	}
	defer db.Close()

	var count int

	db.Table(config.PageIndexTableName).
		Where(fmt.Sprintf("%s.category = ?", config.PageIndexTableName), category).
		Count(&count)

	return count
}
