package auth

import (
    "net/http"
    "strings"
    "github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        if !strings.HasPrefix(token, "Bearer ") {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token format"})
            return
        }

        token = strings.TrimPrefix(token, "Bearer ")
        valid := false
        for userId, t := range Users {
            if token == t {
                c.Set("userId", userId)
                valid = true
                break
            }
        }

        if !valid {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
            return
        }

        c.Next()
    }
}

func GetCustodyUser(userID string) string {
    if custody, ok := CustodyMap[userID]; ok {
        return custody
    }
    return userID
}
