package handler

import "github.com/AlexKLWS/lws-blog-server/router"

func RegisterHandlers(serverRouter *router.Router) {
	serverRouter.Auth.POST("/login", Login)
	serverRouter.Articles.PUT("", NewArticle)
	serverRouter.Pages.PUT("", NewPage)
	serverRouter.Files.PUT("/metadata", AddNewFileMetaData)
	serverRouter.Files.PUT("", AddNewFiles)
}
