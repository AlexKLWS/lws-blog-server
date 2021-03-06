package pages

import (
	"log"

	"github.com/AlexKLWS/lws-blog-server/config"
	"github.com/AlexKLWS/lws-blog-server/models"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

func Get(id string) models.PageData {
	db, err := gorm.Open(viper.GetString(config.GormDialect), viper.GetString(config.GormConnectionString))
	if err != nil {
		log.Printf("DB open error: %s\n", err)
		panic("Failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&models.IconData{})
	db.AutoMigrate(&models.PageData{})

	var page models.PageData
	db.Table(config.PagesTableName).First(&page, "reference_id = ?", id)

	var icon models.IconData
	db.Table(config.IconTableName).First(&icon, "id = ?", page.IconRefer)

	page.Icon = icon

	return page
}
