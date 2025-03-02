package routes

import (
	workout "github.com/avilldaniel/pullup-api/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/workouts", workout.GetWorkouts)
	router.POST("/workout", workout.CreateWorkout)
	router.GET("/contacts", workout.GetContacts)
}
