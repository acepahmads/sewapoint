// cmd/main.go
package main

import (
	"internal/config"
	"log"
	"net/http"
	"os"
	"sewapoint/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.NewDB(&cfg)
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}
	defer db.Close()
	r := gin.Default()
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
