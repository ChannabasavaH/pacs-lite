package middleware

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func RequirePatientStaff() gin.HandlerFunc {
    return func(c *gin.Context) {
        userIDValue, exists := c.Get("userID")
        if !exists {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
                "error": "Authentication required",
            })
            return
        }

        accountTypeValue, exists := c.Get("accountType")
        if !exists {
            c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
                "error": "Staff role missing",
            })
            return
        }

        userID, ok := userIDValue.(string)
        if !ok || userID == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
                "error": "Invalid account identity",
            })
            return
        }

        role, ok := accountTypeValue.(string)
        if !ok {
            c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
                "error": "Invalid account role",
            })
            return
        }

        switch role {
        case "RADIOLOGIST" :
            c.Next()
        default:
            c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
                "error": "Insufficient permissions",
            })
        }
    }
}