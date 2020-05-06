package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/symball/quick-gin/handlers"
)

func setupRouter() *gin.Engine {
	log.Print("Setting up router")
	r := gin.Default()

	corsConfig := cors.DefaultConfig()

	corsOrigins := os.Getenv("APP_CORS_ORIGINS")
	log.Println("CORS Origins: " + corsOrigins)

	// CORS
	if corsOrigins == "*" {
		corsConfig.AllowAllOrigins = true
	} else {
		corsConfig.AllowOrigins = []string{corsOrigins}
	}
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	r.Use(cors.New(corsConfig))

	// Route Definition
	r.GET("/", handlers.IndexHandler)
	r.OPTIONS("/*anything", preflight)

	// Project Actions
	r.GET("/projects", handlers.ProjectsHandler)

	return r
}

func main() {
	log.Print("Quick Gin started")
	_ = godotenv.Load()

	router := setupRouter()

	// If development environment, use dev https cert
	env := os.Getenv("GIN_MODE")
	if "debug" == env {
		if err := router.RunTLS(":"+os.Getenv("APP_PORT"), "./config/certfull.pem", "./config/certkey.pem"); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := router.Run(":" + os.Getenv("APP_PORT")); err != nil {
			log.Fatal(err)
		}
	}
}

func preflight(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(http.StatusOK, struct{}{})
}
