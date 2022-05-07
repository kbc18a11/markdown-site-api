package models

// CreateMeResponse - ユーザー情報
type CreateMeResponse struct {

	// ID
	Id string `json:"id"`

	// ユーザーネーム
	Name string `json:"name"`

	// メールアドレス
	Email string `json:"email"`
}
