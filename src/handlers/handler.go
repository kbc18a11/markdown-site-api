package handlers

import (
	"github.com/GIT_USER_ID/GIT_REPO_ID/src/modules"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

/*
テストコード作成時に、モック化が必要になるものは、プロパティに関数を定義する。
*/
type Handler struct {
	/*
		セッション取得
	*/
	GetSession func(c echo.Context) (*sessions.Session, error)

	/*
		セッション状態を保存
	*/
	SaveSession func(session *sessions.Session, c echo.Context)

	csrf modules.Csrf
}

/*
初期化処理
*/
func (handler *Handler) Init() {
	handler.csrf.Init()

	// セッション取得処理を初期化
	handler.GetSession = func(c echo.Context) (*sessions.Session, error) {
		return session.Get("session", c)
	}

	// セッション状態の保存処理を初期化
	handler.SaveSession = func(session *sessions.Session, c echo.Context) {
		session.Save(c.Request(), c.Response())
	}
}
