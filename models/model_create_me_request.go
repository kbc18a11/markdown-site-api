package models

// CreateMeRequest - 初期登録時のユーザー情報
type CreateMeRequest struct {

	// ユーザーネーム
	Name string `json:"name"`

	// メールアドレス
	Email string `json:"email"`

	// パスワード
	Password string `json:"password"`

	// csrfトークン
	Csrf string `json:"csrf,omitempty"`
}
