package routers

import (
	"gin-h8/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/books", controllers.CreateBook)       // Add new book
	router.GET("/books", controllers.GetBooks)          // Get all books
	router.GET("/books/:id", controllers.GetBookById)   // Get book by id
	router.PUT("/books/:id", controllers.UpdateBook)    // Update book by id
	router.DELETE("/books/:id", controllers.DeleteBook) // Delete book by id

	return router
}
