package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/GO/controller"
	_ "github.com/swaggo/swag/example/GO/docs"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		fmt.Println(c.Request.Method)

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func main() {
	r := gin.Default()

	c := controller.NewController()

	r.Use(CORSMiddleware())


	v1 := r.Group("/api/v1") 
	{
		auth := v1.Group("/auth")
		{
			auth.POST("", c.RegisterUser)
			auth.GET(":email/:password", c.Login)
		}

		users := v1.Group("/allusers")
		{	
			users.GET("", c.ListUsers)
		}

		post := v1.Group("/post")
		{
			post.POST("", c.AddPost)
			post.DELETE(":id", c.DeletePost)
			post.PUT("", c.UpdatePost)
		}
		
		posts := v1.Group("/allposts")
		{
			posts.GET("", c.ListPosts)
		}
	
	}
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}

