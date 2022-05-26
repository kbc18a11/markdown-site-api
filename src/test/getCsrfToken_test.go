package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/GIT_USER_ID/GIT_REPO_ID/src/controller"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

/*
セッションライブラリのモック
*/
type MockSession struct {
	mock.Mock
}

/*
セッションインスタンスの取得関数のモック化
*/
func (mock *MockSession) Get(name string, c echo.Context) (*sessions.Session, error) {
	return &sessions.Session{
			ID:      "",
			Values:  nil,
			Options: nil,
			IsNew:   true,
		},
		nil
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

	// sessionライブラリをモック化しないとヌルポが起きる
	if assert.NoError(t, controller.GetCsrfToken(c)) {
	}
}
