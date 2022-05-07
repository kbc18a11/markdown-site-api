package main

import (
	"github.com/GIT_USER_ID/GIT_REPO_ID/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

    //todo: handle the error!
	c, _ := handlers.NewContainer()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())


	// GetCsrfToken - csrfトークン取得
	e.GET("/api/v1/csrf", c.GetCsrfToken)

	// Login - ログイン
	e.POST("/api/v1/login", c.Login)

	// CreateMe - ユーザー情報登録
	e.POST("/api/v1/register", c.CreateMe)


	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}