package guides

import (
	"log"

	"github.com/AlexKLWS/lws-blog-server/config"
	"github.com/AlexKLWS/lws-blog-server/models"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

func Get(id string) models.GuideData {
	db, err := gorm.Open(viper.GetString(config.GormDialect), viper.GetString(config.GormConnectionString))
	if err != nil {
		log.Printf("DB open error: %s\n", err)
		panic("Failed to connect database")
	}
	defer db.Close()

	var guide models.GuideData
	db.Table(config.GuidesTabelName).First(&guide, "reference_id = ?", id)

	var icon models.IconData
	db.Table(config.IconTableName).First(&icon, "id = ?", guide.IconRefer)

	guide.Icon = icon

	return guide
}
