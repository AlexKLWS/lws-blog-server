package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)


type Login struct {
	Password string `json:"password"`
}

type Session struct {
	Token string `json:"token" xml:"token"`
}

type IconData struct {
	Data   string `json:"data" xml:"data"`
	Height string `json:"height" xml:"height"`
	Width  string `json:"width" xml:"width"`
}

type ArticleData struct {
	Name        string   `json:"name" xml:"name"`
	Subtitle    string   `json:"subtitle" xml:"subtitle"`
	ArticleText string   `json:"articleText" xml:"articleText"`
	Icon        IconData `json:"icon" xml:"icon"`
}

func cookieCheckMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("token")
		if err != nil {
			log.Printf("NO COOKIE HERE!")
			return err
		}
		if cookie.Value != "lol" {
			log.Printf("Actual cookie is: %s\n", cookie.String())
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return next(c)
	}
}

func main() {
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
	e.Static("/", "../client/build")

	// Secret group
	e.Group("/secret", cookieCheckMiddleware)

	api := e.Group("/api")

	api.POST("/login", func(c echo.Context) error {
		loginData := Login{}

		defer c.Request().Body.Close()

		err := json.NewDecoder(c.Request().Body).Decode(&loginData)
		if err != nil {
			log.Printf("Failed processing login request: %s\n", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		// newSessionToken := uuid.Must(uuid.NewV4())

		// s := &Session{
		// 	Token: newSessionToken.String(),
		// }

		var s *Session = nil

		if loginData.Password == "abcd" {
			s = &Session{
				Token: "lol",
			}
		}

		return c.JSON(http.StatusOK, s)
	})

	api.POST("/new-article", func(c echo.Context) error {
		articleData := ArticleData{}

		defer c.Request().Body.Close()

		err := json.NewDecoder(c.Request().Body).Decode(&articleData)
		if err != nil {
			log.Printf("Failed processing login request: %s\n", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return c.String(http.StatusOK, "OK")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
