package handler

import "github.com/AlexKLWS/lws-blog-server/router"

func RegisterHandlers(serverRouter *router.Router) {
	serverRouter.Auth.POST("/login", Login)
	serverRouter.Articles.PUT("", NewArticle)
	serverRouter.Files.PUT("/metadata", NewFileMetaData)
	serverRouter.Files.PUT("", NewFiles)
}
