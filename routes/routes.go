package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	handlers "github.com/karansinghgit/golib/handlers"
)

func Routes(router *gin.Engine) {
	router.GET("/", welcome)

	router.GET("/books", handlers.GetBooks)
	router.POST("/book", handlers.CreateBook)
	router.GET("/book/:bookId", handlers.GetBook)
	router.PUT("/book/:bookId", handlers.EditBook)
	router.DELETE("/book/:bookId", handlers.DeleteBook)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "A Library API for GoLang, with Gin and MongoDB",
	})
	return
}
