// package middleware

// import (
// 	"desktop/go-projectt/initializers"
// 	"fmt"
// 	"net/http"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/golang-jwt/jwt/v5"
// )

// func RequireAuth(c *gin.Context) {

// 	tokenString, err := c.Cookie("Authorization")

// 	if err != nil {
// 		c.AbortWithStatus(http.StatusUnauthorized)
// 	}

// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

// 		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 			if float64(time.Now().Unix()) > claims["exp"].(float64) {
// 				c.AbortWithStatus(http.StatusUnauthorized)
// 			}
// 			var user models.User
// 			initializers.DB.First(&user, claims["sub"])

// 			if user.ID == 0 {
// 				c.AbortWithStatus(http.StatusUnauthorized)
// 			}

// 			c.Set("user", user)

// 			c.Next()
// 			fmt.Println(claims["foo"], claims["nbf"])
// 		} else {
// 			c.AbortWithStatus(http.StatusUnauthorized)
// 		}
// 	}
// }

package middleware

import (
	"desktop/go-projectt/initializers"
	"desktop/go-projectt/models"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5" // Changed jwt/v5 to jwt/v4
)

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check if token is valid
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil // Replace "your-secret-key" with your actual secret key
	})

	if err != nil || !token.Valid {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	expFloat, ok := claims["exp"].(float64)
	if !ok || time.Now().Unix() > int64(expFloat) {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Fetch user from database
	var user models.User
	if err := initializers.DB.First(&user, claims["sub"]).Error; err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Set("user", user)

	c.Next()
	fmt.Println(claims["foo"], claims["nbf"])
}
