package middleware

import (
	"auth-service/internal/config"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	jwtcfg := config.JWTLoadConfig()

	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")

		if err != nil || tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authentication required",
			})
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, jwt.ErrSignatureInvalid
			}

			return []byte(jwtcfg.JWT_Secret), nil
		},

			jwt.WithValidMethods([]string{
				jwt.SigningMethodHS256.Alg(),
			}),
			jwt.WithExpirationRequired(),
		)

		if err != nil || token == nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid or expired token",
			})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token claims",
			})
			return
		}

		userID, ok := claims["id"].(string)
		if !ok || userID == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token id",
			})
			return
		}

		accountType, ok := claims["account_type"].(string)
		if !ok || accountType == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid account type",
			})
			return
		}

		c.Set("userID", userID)
		c.Set("accountType", accountType)

		c.Next()
	}
}
