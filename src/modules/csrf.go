package modules

import (
	"math/rand"
)

/*
テストコード作成時に、モック化が必要になるものは、プロパティに関数を定義する。
*/
type Csrf struct {
	/*
		CSRFトークンの作成
	*/
	CreateCsrfToken func() string
}

/*
初期化処理
*/
func (csrf *Csrf) Init() {
	// CSRFトークンの作成処理の初期化
	csrf.CreateCsrfToken = func() string {
		// トークに割り当てられる文字列
		tokenParts := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

		csrfToken := make([]byte, 20)
		for i := range csrfToken {
			// 割当文字列から、ランダムに設定
			csrfToken[i] = tokenParts[rand.Intn(len(tokenParts))]
		}

		return string(csrfToken)
	}
}
