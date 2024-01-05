package main

import (
	"golang-review-phone/controllers"
	"golang-review-phone/database"
	"golang-review-phone/middleware"
	"golang-review-phone/routes"
)

func main() {
	database.InitDB()
	controllers.InitModels()

	router := routes.SetupRouter()
	router.Use(middleware.AuthMiddleware())
	controllers.SetupRoutes(router)

	router.Run(":8080")
}
