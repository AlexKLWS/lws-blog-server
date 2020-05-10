package handler

import "github.com/AlexKLWS/lws-blog-server/router"

func RegisterHandlers(serverRouter *router.Router) {
	serverRouter.Auth.POST("/login", Login)

	serverRouter.Articles.PUT("", NewArticle)
	serverRouter.Articles.GET("", GetArticle)

	serverRouter.Pages.PUT("", NewPage)

	serverRouter.Materials.GET("", GetMaterials)

	serverRouter.Files.PUT("/metadata", AddNewFileMetaData)
	serverRouter.Files.PUT("", AddNewFiles)
}
