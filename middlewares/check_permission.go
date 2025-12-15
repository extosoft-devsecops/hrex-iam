package middlewares

import (
	"github.com/gin-gonic/gin"
)

type ScopeFunc func(c *gin.Context)

func CheckPermission(resource string, action string, scopeFunc ScopeFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}
