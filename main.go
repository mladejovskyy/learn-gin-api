package main

import "github.com/gin-gonic/gin"

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{
	{ID: "1", Title: "1984", Author: "George Orwell"},
	{ID: "2", Title: "The Hobbit", Author: "J.R.R. Tolkien"},
}

func getBooks(c *gin.Context) {
	c.JSON(200, books)
}

func getBookByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through books to find matching ID
	for _, book := range books {
		if book.ID == id {
			c.JSON(200, book)
			return
		}
	}

	c.JSON(404, gin.H{"message": "book not found"})
}

func createBook(c *gin.Context) {
	var newBook Book

	// Bind JSON from request body to newBook struct
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	books = append(books, newBook)

	c.JSON(201, newBook)
}

func updateBook(c *gin.Context) {
	id := c.Param("id")

	// find book
	for i, book := range books {
		if book.ID == id {
			// bind updated data
			if err := c.ShouldBindJSON(&books[i]); err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}
			c.JSON(200, books[i])
			return
		}
	}

	c.JSON(404, gin.H{"error": "book not found"})
}

func deleteBook(c *gin.Context) {
	id := c.Param("id")

	for i, book := range books {
		if book.ID == id {
			// Remove from slice: append everything before + everything after
			books = append(books[:i], books[i+1:]...)
			c.JSON(200, gin.H{"message": "book deleted"})
			return
		}
	}

	c.JSON(404, gin.H{"error": "book not found"})
}

// main func
func main() {
	r := gin.Default()

	r.GET("/books", getBooks)

	r.GET("/books/:id", getBookByID)

	r.POST("/books", createBook)

	r.PUT("/books/:id", updateBook)

	r.DELETE("/books/:id", deleteBook)

	r.Run()
}
