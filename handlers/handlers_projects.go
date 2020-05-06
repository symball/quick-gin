package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/symball/quick-gin/models"
)

// ProjectsHandler - Get a collection of projects from the DB
func ProjectsHandler(c *gin.Context) {

	projects, err := models.ProjectsFetch()
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"projects": projects,
	})

}
