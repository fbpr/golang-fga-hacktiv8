package router

import (
	"final-project/controllers"
	"final-project/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	usersRouter := r.Group("/users")
	{
		usersRouter.POST("/register", controllers.CreateUser)
		usersRouter.POST("/login", controllers.LoginUser)
		usersRouter.Use(middlewares.Authentication())
		{
			usersRouter.PUT("/:userId", middlewares.UserAuthorization(), controllers.UpdateUserByID)
			usersRouter.DELETE("/:userId", middlewares.UserAuthorization(), controllers.DeleteUserByID)
		}
	}

	photosRouter := r.Group("/photos")
	{
		photosRouter.Use(middlewares.Authentication())
		{
			photosRouter.POST("/", controllers.AddPhoto)
			photosRouter.GET("/", controllers.GetPhotos)
			photosRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhotoByID)
			photosRouter.DELETE("/:photoId", middlewares.PhotoAuthorization(), controllers.DeletePhotoByID)
		}
	}

	commentsRouter := r.Group("/comments")
	{
		commentsRouter.Use(middlewares.Authentication())
		{
			commentsRouter.POST("/", controllers.AddComment)
			commentsRouter.GET("/", controllers.GetComments)
			commentsRouter.PUT("/:commentId", middlewares.CommentAuthorization(), controllers.UpdateCommentByID)
			commentsRouter.DELETE("/:commentId", middlewares.CommentAuthorization(), controllers.DeleteCommentByID)
		}
	}

	socialMediaRouter := r.Group("/socialmedias")
	{
		socialMediaRouter.Use(middlewares.Authentication())
		{
			socialMediaRouter.POST("/", controllers.AddSocialMedia)
			socialMediaRouter.GET("/", controllers.GetSocialMedias)
			socialMediaRouter.PUT("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.UpdateSocialMediaByID)
			socialMediaRouter.DELETE("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.DeleteSocialMediaByID)
		}
	}

	return r
}
