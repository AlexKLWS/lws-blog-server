package handler

import (
	"github.com/AlexKLWS/lws-blog-server/router"
	customMiddleware "github.com/AlexKLWS/lws-blog-server/router/middleware"
)

func RegisterHandlers(serverRouter *router.Router) {
	serverRouter.Auth.POST("/login", Login)

	serverRouter.Articles.PUT("", NewArticle, customMiddleware.CookieCheck)
	serverRouter.Articles.GET("", GetArticle)

	serverRouter.Pages.PUT("", NewPage, customMiddleware.CookieCheck)

	serverRouter.Materials.GET("", GetMaterials)

	serverRouter.Files.PUT("/metadata", AddNewFileMetaData, customMiddleware.CookieCheck)
	serverRouter.Files.PUT("", AddNewFiles, customMiddleware.CookieCheck)
}
