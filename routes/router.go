package router

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"goProject/handler"
	"goProject/models"

	"github.com/gin-gonic/gin"
)

type ParamUser struct {
	Name string `json:"name" binding:"required"`
}

func SetupRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		idParam := c.Query("id")
		id, err := strconv.Atoi(idParam)
		// 사용자 가져오기
		user, err := models.GetUser(db, id)
		if err != nil {
			log.Fatal(err)
		}

		if user.Name == "" {
			c.JSON(http.StatusOK, gin.H{
				"message": "No User",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": user,
			})
		}

	})

	router.GET("/all", func(c *gin.Context) {
		// 사용자 가져오기
		user, err := models.GetUserAll(db)
		if err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"message": user,
		})

	})

	router.GET("/getBooks", func(c *gin.Context) {
		author := c.Query("name")
		books := handler.GetBooks(author)
		c.JSON(http.StatusOK, gin.H{
			"List": books,
		})
	})

	router.POST("/insert", func(c *gin.Context) {

		var paramUser ParamUser
		if err := c.ShouldBindJSON(&paramUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// DB에 데이터 삽입
		_, err := db.Exec("INSERT INTO users (name) VALUES (?)", paramUser.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data into database"})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "result": "success"})
		return

	})

	return router
}
