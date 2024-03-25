package router

import (
	"app/app/controllers"
	"app/app/middlewares"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		public := api.Group("/public")
		{

			public.POST("/login", controllers.Login) // controllers.Login

			public.POST("/signup", controllers.Signup)
		}

		protected := api.Group("/protected").Use(middlewares.Authz())
		{
			protected.GET("/post/", controllers.GetAllPost)
			protected.PUT("/updatePostState")                 // controller.post state
			protected.GET("/posts")                           // controllers.posts
			protected.POST("/update", controllers.UpdatePost) // , controllers.UpdatePost
			protected.GET("/profile", controllers.Profile)    //, controllers.Profile
			protected.POST("/create", controllers.CreatePost) // , controllers.CreatePost
			protected.POST("/delete", controllers.DeletePost) // , controllers.Delete
		}
	}
	return r
}
