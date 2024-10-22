package router

import (
	"net/http"

	"goProject/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	router.GET("/getBooks", func(c *gin.Context) {
		author := c.Query("name")
		books := handler.GetBooks(author)
		c.JSON(http.StatusOK, gin.H{
			"List": books,
		})
	})

	return router
}
