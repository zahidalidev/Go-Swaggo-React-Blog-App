package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/GO/controller"
	_ "github.com/swaggo/swag/example/GO/docs"
)


func main() {
	r := gin.Default()

	c := controller.NewController()


	v1 := r.Group("/api/v1") 
	{
		auth := v1.Group("/auth")
		{
			auth.POST("", c.RegisterUser)
			auth.GET(":email/:password", c.Login)
		}

		users := v1.Group("/allusers")
		{	
			users.GET("", c.ListAccounts)
		}

		post := v1.Group("/post")
		{
			post.POST("", c.AddPost)
			post.DELETE(":id", c.DeletePost)
			post.PUT("", c.UpdatePost)
		}
		
		posts := v1.Group("/listallposts")
		{
			posts.GET("", c.ListAllPosts)
		}
	
	}
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}

