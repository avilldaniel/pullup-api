package workout

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/avilldaniel/pullup-api/config"
	"github.com/gin-gonic/gin"
)

// WorkoutGroup - Custom type for workout group
type WorkoutGroup int

// Declare related constants for each workout group
const (
	Chest WorkoutGroup = iota
	Back
	Legs
	Biceps
	Triceps
	Shoulders
	Glutes
	Core
	Cardio
	Other
)

// Give the workout group a String function
func (w WorkoutGroup) String() string {
	// return [...]string{
	// 	"Chest",
	// 	"Back",
	// 	"Legs",
	// 	"Biceps",
	// 	"Triceps",
	// 	"Shoulders",
	// 	"Glutes",
	// 	"Core",
	// 	"Cardio",
	// 	"Other",
	// }[w]
	workoutNames := [10]string{
		"Chest",
		"Back",
		"Legs",
		"Biceps",
		"Triceps",
		"Shoulders",
		"Glutes",
		"Core",
		"Cardio",
		"Other",
	}
	return workoutNames[w]
}

type workout struct {
	ID     string         `json:"id"`
	Name   string         `json:"name`
	Weight float32        `json:"weight`
	Reps   float32        `json:"reps"`
	Groups []WorkoutGroup `json:"groups"`
}

// Seed data
var workouts = []workout{
	{ID: "1", Name: "Bench Press", Weight: 125, Reps: 8, Groups: []WorkoutGroup{Chest, Shoulders}},
	{ID: "2", Name: "Pullups", Weight: 0, Reps: 6, Groups: []WorkoutGroup{Back, Biceps}},
	{ID: "3", Name: "Deadlift", Weight: 125, Reps: 12, Groups: []WorkoutGroup{Legs, Core}},
}

// Get Workouts - list of all workouts as JSON
func GetWorkouts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, workouts)
}

// Create Workout
func CreateWorkout(c *gin.Context) {
	var newWorkout workout

	// Bind JSON and handle errors properly
	if err := c.BindJSON(&newWorkout); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid workout data: " + err.Error(),
		})
	}

	// Validate workout data
	if newWorkout.Name == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Workout name is required:",
		})
	}

	// Add new workout to workouts slice
	workouts = append(workouts, newWorkout)

	// Return created workout with 201 status
	c.IndentedJSON(http.StatusCreated, newWorkout)
}

// TEST - Get 'contacts' table from Supabase PostgreSQL
func GetContacts(c *gin.Context) {
	rows, err := config.DB.Query(context.Background(), "SELECT id, email FROM contacts")
	fmt.Println(err)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get contacts",
		})
		return
	}

	// Print query results
	for rows.Next() {
		var id int
		var email string
		err := rows.Scan(&id, &email) // Adjust based on actual table columns
		if err != nil {
			log.Fatalf("Failed to scan row: %v\n", err)
		}
		fmt.Printf("ID: %d, Email: %s\n", id, email)
	}

	c.IndentedJSON(http.StatusOK, err)
}
