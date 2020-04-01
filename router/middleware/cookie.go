package middleware

import (
	"github.com/AlexKLWS/lws-blog-server/auth"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func CookieCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("token")
		if err != nil {
			log.Printf("NO COOKIE HERE!")
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		if !auth.TokenExistsInStorage(cookie.Value) {
			log.Printf("User with wrong cookie attempts to access the secret page! \n Their token is: %s\n", cookie.String())
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return next(c)
	}
}
