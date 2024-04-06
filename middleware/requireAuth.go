package middleware

import (
	"fmt"
	"net/http"
	"os"
	"ticketing_movies_API_GO/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {

	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Login terlebih dahulu"})
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Login terlebih dahulu"})
		}

		var user models.User
		if err := models.DB.First(&user, claims["sub"]).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Login terlebih dahulu"})
		}

		c.Set("user", user)

		c.Next()

	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Login terlebih dahulu"})
	}

}
