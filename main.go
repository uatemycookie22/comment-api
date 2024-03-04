package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"portfolio-2022/comments-api/models"
)

type comment struct {
	Message string `json:"message"`
}

func main() {
	err := models.ConnectDatabase()
	checkErr(err)

	// Start up HTTP server
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.POST("comment", addComment)
	}

	r.Run()
}

func addComment(c *gin.Context) {
	var json models.Comment

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := models.CreateComment(json)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	} else {
		c.JSON(http.StatusCreated, gin.H{"New Comment ID": id})
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
