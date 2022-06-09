package middlewares

import (
	"fmt"
	"net/http"

	"github.com/batthanhvan/src/db"
	"github.com/batthanhvan/src/db/users"
	"github.com/batthanhvan/src/lib"
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

func Only(role lib.ROLE) gin.HandlerFunc {
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

		if err := db.DB.Model(users.User{}).Where(`"users"."id" = ? AND "users"."role" = ?"`, userID, role).Error; err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		c.Next()
	}
}

// func byPassCheckAuth(g *gin.Context) {
// 	id := g.Query("userId")
// 	role := g.Query("role")

// 	convert_role := c.ROLE(c.ROLE_value[role])
// 	g.Set("userId", id)
// 	g.Set("role", convert_role)
// 	g.Next()
// }
