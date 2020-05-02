package handler

import "github.com/AlexKLWS/lws-blog-server/router"

func RegisterHandlers(serverRouter *router.Router) {
	serverRouter.Auth.POST("/login", Login)

	serverRouter.Articles.PUT("", NewArticle)
	serverRouter.Articles.GET("", GetArticles)

	serverRouter.Pages.PUT("", NewPage)
	serverRouter.Pages.GET("", GetPages)

	serverRouter.Materials.GET("", GetMaterials)

	serverRouter.Files.PUT("/metadata", AddNewFileMetaData)
	serverRouter.Files.PUT("", AddNewFiles)
}
