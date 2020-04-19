package handler

import (
	"encoding/json"
	"github.com/AlexKLWS/lws-blog-server/auth"
	"github.com/AlexKLWS/lws-blog-server/config"
	"github.com/AlexKLWS/lws-blog-server/models"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func Login(c echo.Context) error {
	loginData := models.Login{}

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&loginData)
	if err != nil {
		log.Printf("Failed processing login request: %s\n", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	var s *models.Session = nil

	allowedPasswords := viper.GetStringSlice(config.AllowedPasswords)

	for _, p := range allowedPasswords{
		if p == loginData.Password {
			s = &models.Session{
				Token: auth.NewToken(),
			}
		}
	}

	return c.JSON(http.StatusOK, s)
}
