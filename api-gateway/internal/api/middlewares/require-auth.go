package middlewares

import (
	"fmt"
	"go-microservices-grpc/api-gateway/internal/pb"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type RequireAuthMiddleware struct {
	client pb.AuthServiceClient
}

func (r *RequireAuthMiddleware) RequireAuth(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatus(403)
		return 
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		c.AbortWithStatus(403)
		return 
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		log.Println(err) 
		c.AbortWithStatus(http.StatusUnauthorized)
		return 
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		id := claims["sub"]

		userID, ok := id.(string)
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		user, err := r.client.ValidateToken(c, &pb.ValidateTokenRequest{Token: userID})

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return 
		}

		if user == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return 
		}

		

		
		c.Set("user", user)
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		return 
	}

	c.Next()
}