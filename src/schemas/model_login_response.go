package schemas

// LoginResponse - ユーザー情報
type LoginResponse struct {

	// ID
	Id string `json:"id"`

	// ユーザーネーム
	Name string `json:"name"`

	// メールアドレス
	Email string `json:"email"`
}
