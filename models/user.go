package models

type User struct {
	Key      string `json:"_key"`
	Username string `json:"username"`
	Password string `json:"password"`
}
