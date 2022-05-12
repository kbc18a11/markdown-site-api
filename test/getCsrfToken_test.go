package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
)

func レスポンスのテスト(t *testing.T) {
	server := echo.New()

	request := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	request.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := server.NewContext(request, rec)
	c.SetPath("/api/healthcheck")
}
