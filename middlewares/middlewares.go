package middlewares

import (
	"fmt"
	"net/http"

	"github.com/batthanhvan/src/db"
	"github.com/batthanhvan/src/db/users"
	"github.com/batthanhvan/src/utils/token"
	"github.com/gin-gonic/gin"
)

// func JwtAuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		err := token.TokenValid(c)
// 		if err != nil {
// 			c.String(http.StatusUnauthorized, "Unauthorized")
// 			c.Abort()
// 			return
// 		}
// 		c.Next()
// 	}
// }

func AuthenticateRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			fmt.Println("Token Unauthorized")
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		userID, err := token.ExtractTokenID(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var temp users.User

		if err := db.DB.Model(users.User{}).Where("id = ? AND role = ?", userID, role).Take(&temp).Error; err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		c.Next()
	}
}
