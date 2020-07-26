package guides

import (
	"log"

	"github.com/AlexKLWS/lws-blog-server/config"
	"github.com/AlexKLWS/lws-blog-server/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

func UpdateOrCreate(guide *models.GuideData) {
	db, err := gorm.Open(viper.GetString(config.GormDialect), viper.GetString(config.GormConnectionString))
	if err != nil {
		log.Printf("DB open error: %s\n", err)
		return
	}
	defer db.Close()

	db.AutoMigrate(&models.IconData{})
	db.AutoMigrate(&models.GuideData{})

	var g models.GuideData
	if err := db.Table(config.GuidesTabelName).Where("reference_id = ?", guide.ReferenceId).First(&g).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			db.Table(config.GuidesTabelName).Create(guide)
		}
	} else {
		db.Table(config.IconTableName).Where("id = ?", g.IconRefer).Updates(guide.Icon)
		var guideWithoutIcon = guide
		erasedIconArticle.Icon = models.IconData{}
		db.Table(config.GuidesTabelName).Where("reference_id = ?", guide.ReferenceId).Updates(guideWithoutIcon)
	}
}
