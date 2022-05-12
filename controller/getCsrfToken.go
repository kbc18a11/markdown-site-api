package controller

import (
	"math/rand"
	"net/http"

	"github.com/GIT_USER_ID/GIT_REPO_ID/schemas"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

/*
CSRFトークンの作成
*/
func createCsrfToken() string {
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
func GetCsrfToken(c echo.Context) error {
	// CSRFトークンの生成
	csrfToken := createCsrfToken()

	// CSRFトークンをセッションに保存
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   14400, // 4時間
		HttpOnly: true,
	}
	sess.Values["csrf"] = csrfToken
	sess.Save(c.Request(), c.Response())

	return c.JSON(http.StatusOK, schemas.GetCsrfTokenResponse{
		Token: csrfToken,
	})
}
