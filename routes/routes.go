package routes

import (
	"github.com/gin-gonic/gin"
	"golang-review-phone/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/reviews", controllers.CreateReview)
	router.GET("/reviews", controllers.GetReviews)
	router.GET("/reviews/:id", controllers.GetReviewByID)
	router.POST("/reviews/:id/comment", controllers.PostComment)

	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.LoginUser)

	return router
}
