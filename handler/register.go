package handler

import (
	"github.com/AlexKLWS/lws-blog-server/router"
	customMiddleware "github.com/AlexKLWS/lws-blog-server/router/middleware"
)

func RegisterHandlers(serverRouter *router.Router) {
	serverRouter.Auth.POST("/login", Login)

	serverRouter.Articles.PUT("", UpdateOrCreateArticle, customMiddleware.CookieCheck)
	serverRouter.Articles.GET("", GetArticle)

	serverRouter.Pages.PUT("", UpdateOrCreatePage, customMiddleware.CookieCheck)
	serverRouter.Pages.GET("", GetPage)

	serverRouter.Materials.GET("", GetMaterials)

	serverRouter.Files.PUT("/metadata", AddNewFileMetaData, customMiddleware.CookieCheck)
	serverRouter.Files.PUT("", AddNewFiles, customMiddleware.CookieCheck)
}
