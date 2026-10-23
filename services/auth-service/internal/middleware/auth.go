package middleware

import (
	"auth-service/internal/config"
	"auth-service/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	jwtcfg := config.JWTLoadConfig()
	publicKey, err := utils.LoadPublicKey(jwtcfg.JWT_Public_Key_Path)
	if err != nil {
		panic("Failed to load RSA public key: " + err.Error())
	}

	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")

		if err != nil || tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authentication required",
			})
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if token.Method != jwt.SigningMethodRS256 {
				return nil, jwt.ErrSignatureInvalid
			}

			return publicKey, nil
		},

			jwt.WithValidMethods([]string{"RS256"}),
			jwt.WithExpirationRequired(),
			jwt.WithIssuer("pacs-auth"),
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