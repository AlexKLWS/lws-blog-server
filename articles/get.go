package articles

import (
	"log"

	"github.com/AlexKLWS/lws-blog-server/config"
	"github.com/AlexKLWS/lws-blog-server/models"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

func Get(id string) models.ArticleData {
	db, err := gorm.Open(viper.GetString(config.GormDialect), viper.GetString(config.GormConnectionString))
	if err != nil {
		log.Printf("DB open error: %s\n", err)
		panic("Failed to connect database")
	}
	defer db.Close()

	var article models.ArticleData
	db.Table(config.ArticleTableName).First(&article, "reference_id = ?", id)

	var icon models.IconData
	db.Table(config.IconTableName).First(&icon, "id = ?", article.IconRefer)

	article.Icon = icon

	return article
}
