// cmd/main.go
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
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
