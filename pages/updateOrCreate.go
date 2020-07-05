package pages

import (
	"log"

	"github.com/AlexKLWS/lws-blog-server/config"
	"github.com/AlexKLWS/lws-blog-server/models"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

func UpdateOrCreate(page *models.PageData) {
	db, err := gorm.Open(viper.GetString(config.GormDialect), viper.GetString(config.GormConnectionString))
	if err != nil {
		log.Printf("DB open error: %s\n", err)
		panic("Failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&models.IconData{})
	db.AutoMigrate(&models.PageData{})

	var p models.PageData
	if err := db.Table(config.PagesTableName).Where("reference_id = ?", page.ReferenceId).First(&p).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			db.Table(config.PagesTableName).Create(page)
		}
	} else {
		db.Table(config.PagesTableName).Where("reference_id = ?", page.ReferenceId).Update(page)
	}
}
