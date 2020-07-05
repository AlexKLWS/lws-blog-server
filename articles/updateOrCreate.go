package articles

import (
	"log"

	"github.com/AlexKLWS/lws-blog-server/config"
	"github.com/AlexKLWS/lws-blog-server/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

func UpdateOrCreate(article *models.ArticleData) {
	db, err := gorm.Open(viper.GetString(config.GormDialect), viper.GetString(config.GormConnectionString))
	if err != nil {
		log.Printf("DB open error: %s\n", err)
		return
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&models.IconData{})
	db.AutoMigrate(&models.ArticleData{})

	// Create
	db.Table(config.ArticleTableName).Create(article)
}
