package middleware

import (
	// "fmt"

	"fmt"
	"net/http"
	"sewapoint/pkg/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			fmt.Println("TOken tidak ada")
			c.Redirect(http.StatusFound, "/")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token tidak ditemukan"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := jwt.ValidateJWT(tokenString)
		if err != nil {
			c.Redirect(http.StatusFound, "/")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
			return
		}

		// Simpan user ID ke context
		// c.Set("userID", claims.UserID)
		utils.SetUserID(c, claims.UserID)
		c.Next()
	}
}
