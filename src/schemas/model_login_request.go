package schemas

// LoginRequest - ログイン情報
type LoginRequest struct {

	// メールアドレス
	Email string `json:"email"`

	// パスワード
	Password string `json:"password"`

	// csrfトークン
	Csrf string `json:"csrf,omitempty"`
}
