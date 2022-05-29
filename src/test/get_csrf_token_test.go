package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/GIT_USER_ID/GIT_REPO_ID/src/handlers"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func HandlerMockInit() *handlers.Handler {
	handler := &handlers.Handler{}

	// セッション取得処理をモック化
	handler.GetSession = func(c echo.Context) (*sessions.Session, error) {
		session := &sessions.Session{
			Values: map[interface{}]interface{}{},
		}

		return session, nil
	}

	// セッション状態の保存処理をモック化
	handler.SaveSession = func(session *sessions.Session, c echo.Context) {}

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

	handlerMock := HandlerMockInit()

	// sessionライブラリをモック化しないとヌルポが起きる
	if assert.NoError(t, handlerMock.GetCsrfToken(c)) {
	}
}
