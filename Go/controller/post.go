package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/GO/model"
	"github.com/swaggo/swag/example/GO/mysql"
)

func (c *Controller) AddPost(ctx *gin.Context){
	db := mysql.Connect()

	var post model.Post
	err := ctx.ShouldBindJSON(&post)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	
	temppost := model.Post {
		Title: post.Title,	
		Image: post.Image,
		Description: post.Description,
		User_id: post.User_id,
	}
	
	query := "INSERT INTO testGO.posts (title, image, description, user_id) VALUES (?, ?, ?, ?);"
	result, err := db.Exec(query, temppost.Title, temppost.Image, temppost.Description, temppost.User_id)
	
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	
	id, err := result.LastInsertId()

	ctx.JSON(http.StatusOK, id)
}

func (c* Controller) DeletePost(ctx *gin.Context){
	db := mysql.Connect()
	id := ctx.Params.ByName("id")

	query := "Delete from posts where id = ?"
	result, err := db.Exec(query, id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	
	affected, err := result.RowsAffected()

	if affected == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, "Post Not found to Delete")
		return
	}

	ctx.JSON(http.StatusOK, affected)
}

func (c* Controller) UpdatePost(ctx *gin.Context){
	db := mysql.Connect()

	var post model.Post
	err := ctx.ShouldBindJSON(&post)
	if err != nil {
	 ctx.JSON(http.StatusBadRequest, err)
	 return
	}

	temppost := model.Post {
		Id: post.Id,
		Title: post.Title,
		Image: post.Image,
		Description: post.Description,
	}

	query := "UPDATE posts set title = ?, image = ?, description = ? where id = ?"

	result, err := db.Exec(query, &temppost.Title, &temppost.Image, &temppost.Description, &temppost.Id)

	if err != nil  {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	rowsAffected, err := result.RowsAffected()

	if rowsAffected == 0 || err != nil {
		ctx.JSON(http.StatusNotFound, "Found nothing to update!")
		return
	}

	ctx.JSON(http.StatusOK, rowsAffected)
}

func (c* Controller) ListAllPosts(ctx *gin.Context){
	db := mysql.Connect()
	
	rows, err := db.Query(`select * from posts`)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	} 

	posts := make([]*model.Post, 0)

	for rows.Next(){
		post := new(model.Post)
		rows.Scan(&post.Id, &post.Title, &post.Image, &post.Description, &post.User_id)

		posts = append(posts, post)
	}

	ctx.JSON(http.StatusOK, posts)
}