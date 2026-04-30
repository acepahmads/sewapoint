// cmd/main.go
package main

import (
	"log"
	"net/http"
	"os"
	"sewapoint/internal/config"
	"sewapoint/internal/database"
	"sewapoint/internal/middleware"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Database connection
	cfg := config.LoadConfig()
	db, err := database.NewDB(&cfg)
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}
	defer db.Close()

	// Setup Gin router
	r := gin.Default()
	r.Use(middleware.LoggingMiddleware())
	r.Use(middleware.RateLimitMiddleware(100, time.Minute))
	r.LoadHTMLGlob("templates/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "indexMobile.html", nil)
	})
	r.GET("/dashboard", func(c *gin.Context) {
		c.HTML(200, "dashboard.html", nil)
	})

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "20266"
	}
	log.Println("Server running at :" + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
