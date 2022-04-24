package model

type User struct {
	Id   int       `json:"id" example:"1"`
	Name string    `json:"name" example:"account name"`
	Email string   `json:"email" example:"abc@gmail.com"`
	Password string `json:"password" example:"account password"`
}