package user

// User ユーザーモデル
type User struct {
	ID   int `json:"id" sql:"AUTO_INCREMENT"`
	Name int `json:"name"`
}
