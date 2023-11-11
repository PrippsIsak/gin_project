package routes

import (
	"gin-twitter/app/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {

	// Example route
	router.GET("/persons/:username", handlers.GetPerson)
	router.POST("persons/create", handlers.CreatePerson)
	router.DELETE("/persons/delete/:username", handlers.DeletePerson)
	router.PUT("/persons/update/:username", handlers.UpdatePerson)
	// Add more routes as needed
}
