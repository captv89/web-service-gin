package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /books || Get all books
func findBooks(c *gin.Context) {
	var books []Book
	DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

// GET /books/:id  || Find a book
func findBook(c *gin.Context) {  // Get model if exist
  var book Book

  if err := DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  c.JSON(http.StatusOK, gin.H{"data": book})
}