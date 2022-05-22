package main

import (
	"github.com/GIT_USER_ID/GIT_REPO_ID/src/config"
	"github.com/GIT_USER_ID/GIT_REPO_ID/src/controller"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// ミドルウェアの設定
	config.InitMiddlewares(e)

	e.GET("/api/v1/csrf", controller.GetCsrfToken)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
