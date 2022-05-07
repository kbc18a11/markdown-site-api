package handlers
import (
    "github.com/GIT_USER_ID/GIT_REPO_ID/models"
    "github.com/labstack/echo/v4"
    "net/http"
)

// GetCsrfToken - csrfトークン取得
func (c *Container) GetCsrfToken(ctx echo.Context) error {
    return ctx.JSON(http.StatusOK, models.HelloWorld {
        Message: "Hello World",
    })
}


// Login - ログイン
func (c *Container) Login(ctx echo.Context) error {
    return ctx.JSON(http.StatusOK, models.HelloWorld {
        Message: "Hello World",
    })
}
