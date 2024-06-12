package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Define a struct to represent a Task
type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

var tasks []Task

func main() {
	router := gin.Default()

	// Routes
	router.GET("/tasks", getTasks)
	router.GET("/tasks/:id", getTaskByID)
	router.POST("/tasks", createTask)
	router.PUT("/tasks/:id", updateTask)
	router.DELETE("/tasks/:id", deleteTask)

	// Start server
	if err := router.Run(":8080"); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

// Handler to get all tasks
func getTasks(c *gin.Context) {
	c.JSON(http.StatusOK, tasks)
}

// Handler to get a task by ID
func getTaskByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	for _, task := range tasks {
		if task.ID == id {
			c.JSON(http.StatusOK, task)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

// Handler to create a new task
func createTask(c *gin.Context) {
	var newTask Task
	if err := c.BindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)

	c.JSON(http.StatusCreated, newTask)
}

// Handler to update a task by ID
func updateTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var updatedTask Task
	if err := c.BindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i] = updatedTask
			c.JSON(http.StatusOK, updatedTask)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

// Handler to delete a task by ID
func deleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}
