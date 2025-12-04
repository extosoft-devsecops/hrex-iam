package middlewares

import (
	"net/http"

	"github.com/extosoft-devsecops/hrex-iam/types"
	"github.com/extosoft-devsecops/hrex-iam/util"
	"github.com/gin-gonic/gin"
)

// ScopeResolver is a function that decides the required scope
// for a particular request (based on path param, method, etc.).
type ScopeResolver func(c *gin.Context) string

// NewPermissionMiddleware creates a gin.HandlerFunc that checks
// permission in form of "<resource>:<action>:<scope>" against
// the permissions in context (set by AuthContextMiddleware).
func NewPermissionMiddleware(resource, action string, scopeResolver ScopeResolver) gin.HandlerFunc {
	return func(c *gin.Context) {
		userPerms := util.GetStringSlice(c, util.CtxPermissionsKey)
		if len(userPerms) == 0 {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error":   "permission_denied",
				"message": "no permissions in context",
			})
			return
		}

		var requiredScope string
		if scopeResolver != nil {
			requiredScope = scopeResolver(c)
		} else {
			// default: require tenant scope
			requiredScope = types.ScopeTenant
		}

		required := types.Permission{
			Resource: resource,
			Action:   action,
			Scope:    requiredScope,
		}

		if !types.HasPermission(userPerms, required) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error":   "permission_denied",
				"message": "insufficient permission for this resource",
			})
			return
		}

		c.Next()
	}
}

// ScopeGlobal returns a resolver that always requires global scope.
func ScopeGlobal() ScopeResolver {
	return func(c *gin.Context) string {

		return types.ScopeGlobal
	}
}

// ScopeTenant returns a resolver that always requires tenant scope.
func ScopeTenant() ScopeResolver {
	return func(c *gin.Context) string {

		return types.ScopeTenant
	}
}

// ScopeSelfOnly returns a resolver that always requires self scope.
func ScopeSelfOnly() ScopeResolver {
	return func(c *gin.Context) string {
		
		return types.ScopeSelf
	}
}

// ScopeSelfOrTenantFromParam:
// ถ้า path param (เช่น :id) == userId → ใช้ self
// ถ้าไม่ใช่ → ใช้ tenant
//
// ใช้กับ route เช่น:
//
//	GET /users/:id
//	DELETE /users/:id
func ScopeSelfOrTenantFromParam(param string) ScopeResolver {
	return func(c *gin.Context) string {
		requestID := c.Param(param)
		userID := util.GetString(c, util.CtxUserIDKey)

		if requestID != "" && userID != "" && requestID == userID {
			return types.ScopeSelf
		}
		return types.ScopeTenant
	}
}

// ScopeFromCustomFunc ใช้ custom logic เองเต็มที่
func ScopeFromCustomFunc(fn func(c *gin.Context) string) ScopeResolver {
	return func(c *gin.Context) string {
		return fn(c)
	}
}
