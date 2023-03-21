package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	BookID     string `json:"book_id"`
	BookTitle  string `json:"book_title"`
	BookAuthor string `json:"book_author"`
	BookDesc   string `json:"book_desc"`
}

var BooksData = []Book{}

func CreateBook(ctx *gin.Context) {
	var newBook Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newBook.BookID = fmt.Sprintf("b%d", len(BooksData)+1)
	BooksData = append(BooksData, newBook)

	ctx.JSON(http.StatusCreated, gin.H{
		"car": newBook,
	})
}

func GetBooks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"books": BooksData,
	})
}

func GetBookById(ctx *gin.Context) {
	bookID := ctx.Param("id")

	for _, book := range BooksData {
		if book.BookID == bookID {
			ctx.JSON(http.StatusOK, gin.H{
				"book": book,
			})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"message": "book not found",
	})
}

func DeleteBook(ctx *gin.Context) {
	bookID := ctx.Param("id")

	for i, book := range BooksData {
		if book.BookID == bookID {
			BooksData = append(BooksData[:i], BooksData[i+1:]...)
			ctx.JSON(http.StatusOK, gin.H{
				"message": "book deleted",
			})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"message": "book not found",
	})
}

func UpdateBook(ctx *gin.Context) {
	bookID := ctx.Param("id")

	for i, book := range BooksData {
		if book.BookID == bookID {
			var newBook Book

			if err := ctx.ShouldBindJSON(&newBook); err != nil {
				ctx.AbortWithError(http.StatusBadRequest, err)
				return
			}

			newBook.BookID = bookID
			BooksData[i] = newBook

			ctx.JSON(http.StatusOK, gin.H{
				"book": newBook,
			})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"message": "book not found",
	})
}
