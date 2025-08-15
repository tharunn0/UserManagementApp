package jwt

import (
	"fmt"
	"os"
	"strings"
	"time"

	log "github.com/tharunn0/gin-server-gorm/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
)

func Issue(username string, role string) string {

	key := []byte(os.Getenv("HS_256KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"role":     role,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(2 * time.Minute).Unix(),
	})
	fmt.Println(token)
	str, er := token.SignedString(key)
	if er != nil || str == "" {
		log.Error("jwt.Issue", zap.Error(er))
		return ""
	}
	return str
}

func ValidateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// extracting authorization header from the req
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(401, gin.H{"error": "Login to access page"})
			log.Warn("No Jwt Found")
			c.Abort()
			return
		}
		// extracted the token into authstring
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		// validating signature
		jtoken, er := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Jwt signing algorithm mismatch")
			}
			return []byte(os.Getenv("HS_256KEY")), nil
		})
		if er != nil || !jtoken.Valid {
			c.JSON(401, gin.H{"error": "Auth token is invalid"})
			c.Abort()
			return
		}
		// validating the claims
		if claims, ok := jtoken.Claims.(jwt.MapClaims); ok {
			username := claims["username"].(string)
			role := claims["role"].(string)
			if role != "user" {
				c.JSON(401, gin.H{"Unauthorized": "Login as user to continue"})
				c.Abort()
				return
			}
			c.Set("username", username)
			c.Set("role", role)
		} else {
			c.JSON(401, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		c.Next()
	}
}
func ValidateMiddlewareAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// extracting authorization header from the req
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(401, gin.H{"error": "Login to access page"})
			log.Warn("No Jwt Found")
			c.Abort()
			return
		}
		// extracted the token into authstring
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		// validating signature
		jtoken, er := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Jwt signing algorithm mismatch")
			}
			return []byte(os.Getenv("HS_256KEY")), nil
		})
		if er != nil || !jtoken.Valid {
			c.JSON(401, gin.H{"error": "Auth token is invalid"})
			c.Abort()
			return
		}
		// validating the claims
		if claims, ok := jtoken.Claims.(jwt.MapClaims); ok {
			username := claims["username"].(string)
			role := claims["role"].(string)

			if role != "admin" {
				c.JSON(401, gin.H{"Unauthorized": "Login as admin to continue"})
				c.Abort()
				return
			}

			c.Set("username", username)
			c.Set("role", role)
		} else {
			c.JSON(401, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// func ValidateToken(authHeader string) bool {

// 	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
// 		return false
// 	}
// 	// extracted the token into authstring
// 	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

// 	jtoken, er := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("Jwt signing algorithm mismatch")
// 		}
// 		return []byte(os.Getenv("HS_256KEY")), nil
// 	})
// 	if er != nil || !jtoken.Valid {
// 		return false
// 	}

// 	if claims, ok := jtoken.Claims.(jwt.MapClaims); ok {
// 		r, _ := claims["role"]
// 		role, _ := r.(string)
// 		if role == "user" {

// 			return true
// 		}
// 	}

// 	return false

// }

// func ValidateTokenAdmin(authHeader string) bool {

// 	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
// 		return false
// 	}
// 	// extracted the token into authstring
// 	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

// 	jtoken, er := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("Jwt signing algorithm mismatch")
// 		}
// 		return []byte(os.Getenv("HS_256KEY")), nil
// 	})
// 	if er != nil || !jtoken.Valid {
// 		return false
// 	}

// 	if claims, ok := jtoken.Claims.(jwt.MapClaims); ok {
// 		r, _ := claims["role"]
// 		role, _ := r.(string)
// 		if role == "admin" {
// 			return true
// 		}
// 	}

// 	return false

// }
