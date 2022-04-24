package model

type Post struct {
	Id int 			`json:"id" example:"1"`
	Title string 	`json:"title" example:"post title"`
	Image string 	`json:"image" example:"image url"`
	Description string `json:"description" example:"description of the post"`
	User_id int 	`json:"user_id" example:"1"`
}