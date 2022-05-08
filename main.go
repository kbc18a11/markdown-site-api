package main

import (
	"github.com/GIT_USER_ID/GIT_REPO_ID/config"
	"github.com/GIT_USER_ID/GIT_REPO_ID/controller"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// ミドルウェアの設定
	config.InitMiddlewares(e)

	e.GET("/csrf", controller.GetCsrfToken)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
