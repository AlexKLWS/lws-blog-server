package router

import (
	"net/http"

	customMiddleware "github.com/AlexKLWS/lws-blog-server/router/middleware"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Router struct {
	Server *echo.Echo
	Auth *echo.Group
	Articles *echo.Group
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
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	// Serving the website
	e.Static("/", "../client/build")

	// "Secret" part of the website is loosely protected by cookie,
	// It could only be accessed via history/router
	e.Group("/secret", customMiddleware.CookieCheck)

	api	:= e.Group("/api")

	return &Router{
		Server:   e,
		Auth:     api.Group("/auth"),
		Articles: api.Group("/articles"),
	}
}
