package lib

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BadRequest(c *gin.Context, err error) {
	c.Set("error", err)
	err = Unwrap(err)
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
	})
}

func InternalServerError(c *gin.Context, err error) {
	c.Set("error", err)
	err = Unwrap(err)
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})
}

func Success(c *gin.Context, response interface{}) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code":    200,
		"success": true,
		"data":    response,
	})
}

func Unauthorized(c *gin.Context, err error) {
	c.Set("error", err)
	err = Unwrap(err)
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"error": err.Error(),
	})
}
