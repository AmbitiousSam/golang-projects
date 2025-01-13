package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func getTaskHandler(c *gin.Context) {
	var tasks []Task
	db.Find(&tasks)
	c.JSON(200, tasks)
}

func createTaskHandler(c *gin.Context) {
	var task Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": "Invalid Input"})
		return
	}
	db.Create(&task)
	c.JSON(201, task)
}

func updateTaskHandler(c *gin.Context) {
	id := c.Param("id")                          // Extract the ID from the URL
	log.Printf("Received ID for update: %s", id) // Debug log for ID

	var task Task
	// Explicitly query by ID
	if err := db.First(&task, "id = ?", id).Error; err != nil {
		log.Printf("Error finding task with ID %s: %v", id, err)
		c.JSON(404, gin.H{"error": "Task not found"})
		return
	}

	// Bind updated fields from the request
	var updatedFields struct {
		Title       string `json:"title,omitempty"`
		Description string `json:"description,omitempty"`
		Completed   bool   `json:"completed,omitempty"`
	}
	if err := c.BindJSON(&updatedFields); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	if updatedFields.Title != "" {
		task.Title = updatedFields.Title
	}
	if updatedFields.Description != "" {
		task.Description = updatedFields.Description
	}
	task.Completed = updatedFields.Completed

	// Save the updated task
	if err := db.Save(&task).Error; err != nil {
		log.Printf("Error saving task with ID %s: %v", id, err)
		c.JSON(500, gin.H{"error": "Failed to update task"})
		return
	}

	c.JSON(200, task)
}

func deleteTaskHandler(c *gin.Context) {
	id := c.Param("id")
	if err := db.Delete(&Task{}, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(204, nil)
}
