package handlers

import (
	"math/rand"
	"net/http"

	"github.com/GIT_USER_ID/GIT_REPO_ID/src/schemas"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

/*
CSRFトークンの作成
*/
func (handler *Handler) createCsrfToken() string {
	// トークに割り当てられる文字列
	tokenParts := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	csrfToken := make([]byte, 20)
	for i := range csrfToken {
		// 割当文字列から、ランダムに設定
		csrfToken[i] = tokenParts[rand.Intn(len(tokenParts))]
	}

	return string(csrfToken)
}

/*
CSRFトークンの取得
*/
func (handler *Handler) GetCsrfToken(c echo.Context) error {
	// CSRFトークンの生成
	csrfToken := handler.createCsrfToken()

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
