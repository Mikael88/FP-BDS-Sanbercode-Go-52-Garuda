package controllers

import (
	"golang-review-phone/models"
	"golang-review-phone/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

const invalidRequestBodyMsg = "Invalid request body"

func InitModels() {
	models.InitModels()
}

func SetupRoutes(router *gin.Engine) {
	router.POST("/reviews", CreateReview)
	router.GET("/reviews", GetReviews)
	router.GET("/reviews/:id", GetReviewByID)
	router.POST("/reviews/:id/comment", PostComment)

	router.POST("/register", RegisterUser)
	router.POST("/login", LoginUser)
}

func CreateReview(c *gin.Context) {
	var review models.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, invalidRequestBodyMsg)
		return
	}

	result, err := models.CreateReview(review)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create review")
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetReviews(c *gin.Context) {
	reviews, err := models.GetReviews()
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to get reviews")
		return
	}

	c.JSON(http.StatusOK, reviews)
}

func GetReviewByID(c *gin.Context) {
	id := c.Param("id")

	review, err := models.GetReviewByID(id)
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Review not found")
		return
	}

	c.JSON(http.StatusOK, review)
}

func PostComment(c *gin.Context) {
	id := c.Param("id")

	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, invalidRequestBodyMsg)
		return
	}

	result, err := models.PostComment(id, comment)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to post comment")
		return
	}

	c.JSON(http.StatusOK, result)
}

func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, invalidRequestBodyMsg)
		return
	}

	// Hash the password before storing it
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to register user")
		return
	}

	user.Password = hashedPassword

	result, err := models.RegisterUser(user)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to register user")
		return
	}

	c.JSON(http.StatusCreated, result)
}

func LoginUser(c *gin.Context) {
	var credentials models.LoginCredentials
	if err := c.ShouldBindJSON(&credentials); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, invalidRequestBodyMsg)
		return
	}

	user, err := models.GetUserByUsername(credentials.Username)
	if err != nil {
		utils.RespondWithError(c, http.StatusUnauthorized, "Invalid username or password")
		return
	}

	// Compare the hashed password
	err = utils.ComparePassword(user.Password, credentials.Password)
	if err != nil {
		utils.RespondWithError(c, http.StatusUnauthorized, "Invalid username or password")
		return
	}

	// Generate and return a JWT token
	token, err := utils.GenerateToken(user.ID.Hex(), user.Role)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
