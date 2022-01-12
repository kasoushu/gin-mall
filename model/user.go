package model

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type Admin struct {
	Id       int    `json:"id"`
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
