package model

type User struct {
	Id       uint64 `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
type UserInfo struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type Admin struct {
	Id       uint64 `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

// WebUser frontend
type WebUser struct {
	id    int
	token string
}

type WebLoginUser struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
