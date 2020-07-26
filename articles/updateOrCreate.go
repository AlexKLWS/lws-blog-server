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

	db.AutoMigrate(&models.IconData{})
	db.AutoMigrate(&models.ArticleData{})

	var a models.ArticleData
	if err := db.Table(config.ArticleTableName).Where("reference_id = ?", article.ReferenceId).First(&a).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			db.Table(config.ArticleTableName).Create(article)
		}
	} else {
		db.Table(config.IconTableName).Where("id = ?", a.IconRefer).Updates(article.Icon)
		var articleWithoutIcon = article
		articleWithoutIcon.Icon = models.IconData{}
		db.Table(config.ArticleTableName).Where("reference_id = ?", article.ReferenceId).Updates(articleWithoutIcon)
	}
}
