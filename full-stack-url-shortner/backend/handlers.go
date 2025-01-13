package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func shortenURLHandler(c *gin.Context) {
	var request struct {
		LongURL string `json:"long_url"`
	}
	if err := c.BindJSON(&request); err != nil || request.LongURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Generate a unique short URL
	shortID := uuid.New().String()[:8]
	shortURL := "http://localhost:8080/" + shortID

	// Save the mapping
	mapping := URLMapping{ShortURL: shortID, LongURL: request.LongURL}
	db.Create(&mapping)

	c.JSON(http.StatusOK, gin.H{"short_url": shortURL})
}

func redirectHandler(c *gin.Context) {
	shortID := c.Param("shortID")

	var mapping URLMapping
	result := db.First(&mapping, "short_url = ?", shortID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
		return
	}

	c.Redirect(http.StatusFound, mapping.LongURL)
}
