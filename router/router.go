package router

import (
	"github.com/AlexKLWS/lws-blog-server/config"
	"github.com/spf13/viper"
	"net/http"

	customMiddleware "github.com/AlexKLWS/lws-blog-server/router/middleware"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Router struct {
	Server   *echo.Echo
	Auth     *echo.Group
	Articles *echo.Group
	Pages    *echo.Group
	Files    *echo.Group
}

// New echo router
func New() *Router {
	e := echo.New()

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "../client/build",
		HTML5:  true,
		Browse: false,
	}))
	e.Use(middleware.Recover())

	if viper.GetString(config.Env) == config.Debug {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowCredentials: true,
			AllowOrigins: []string{"http://localhost:3000"},
			AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		}))
	}

	// Serving the website
	e.Static("/", "../client/build")

	// "Secret" part of the website is loosely protected by cookie,
	// It could only be accessed via history/router
	e.Group("/secret", customMiddleware.CookieCheck)

	a := e.Group("/api", customMiddleware.CookieCheck)

	return &Router{
		Server:   e,
		Auth:     e.Group("/auth"),
		Articles: a.Group("/articles"),
		Pages:    a.Group("/pages"),
		Files:    a.Group("/files"),
	}
}
