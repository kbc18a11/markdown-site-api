package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/GIT_USER_ID/GIT_REPO_ID/src/controller"
	"github.com/labstack/echo/v4"
)

/*
レスポンステスト
*/
func TestResponse(t *testing.T) {
	server := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	responseRecorder := httptest.NewRecorder()
	c := server.NewContext(request, responseRecorder)
	c.SetPath("/api/v1/csrf")

	if err := controller.GetCsrfToken(c); err != nil {
		t.Fatal(err)
	}
}
