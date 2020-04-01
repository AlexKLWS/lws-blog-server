package main

import (
	"github.com/AlexKLWS/lws-blog-server/auth"
	"github.com/AlexKLWS/lws-blog-server/config"
	"github.com/AlexKLWS/lws-blog-server/handler"
	"github.com/AlexKLWS/lws-blog-server/router"
)

func main() {
	config.SetupViper()
	auth.InitializeTokenStorage()

	r := router.New()
	handler.RegisterHandlers(r)

	r.Server.Logger.Fatal(r.Server.Start(":1323"))
}
