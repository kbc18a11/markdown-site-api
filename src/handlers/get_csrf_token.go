package handlers

import (
	"net/http"

	"github.com/GIT_USER_ID/GIT_REPO_ID/src/schemas"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

/*
CSRFトークンの取得
*/
func (handler *Handler) GetCsrfToken(c echo.Context) error {
	// CSRFトークンの生成
	csrfToken := handler.Csrf.CreateCsrfToken()

	// CSRFトークンをセッションに保存するために、セッションの設定
	session, _ := handler.GetSession(c)
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   14400, // 4時間
		HttpOnly: true,
	}
	session.Values["csrf"] = csrfToken

	// CSRFトークンをセッションに保存
	handler.SaveSession(session, c)

	return c.JSON(http.StatusOK, schemas.GetCsrfTokenResponse{
		Token: csrfToken,
	})
}
