package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/GIT_USER_ID/GIT_REPO_ID/src/handlers"
	"github.com/GIT_USER_ID/GIT_REPO_ID/src/schemas"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var csrfToken = "token"

/*
APIの処理を初期化
*/
func HandlerMockInit() *handlers.Handler {
	handler := &handlers.Handler{}
	handler.Init()

	// セッション取得処理をモック化
	handler.GetSession = func(c echo.Context) (*sessions.Session, error) {
		session := &sessions.Session{
			Values: map[interface{}]interface{}{},
		}

		return session, nil
	}

	// セッション状態の保存処理をモック化
	handler.SaveSession = func(session *sessions.Session, c echo.Context) {}

	handler.Csrf.CreateCsrfToken = func() string {
		return csrfToken
	}

	return handler
}

/*
レスポンステスト
*/
func TestResponse(t *testing.T) {
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	responseRecorder := httptest.NewRecorder()

	c := e.NewContext(request, responseRecorder)
	c.SetPath("/api/v1/csrf")

	// 検証
	if assert.NoError(t, HandlerMockInit().GetCsrfToken(c)) {
		assert.Equal(t, http.StatusOK, responseRecorder.Code)

		// レスポンスボディの構造体化
		var responseBody schemas.GetCsrfTokenResponse
		json.Unmarshal([]byte(responseRecorder.Body.Bytes()), &responseBody)

		assert.Equal(t, responseBody, schemas.GetCsrfTokenResponse{
			Token: csrfToken,
		})
	}
}
