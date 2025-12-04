package util

import (
	"github.com/gin-gonic/gin"
)

// Context key names – ต้องใช้ให้ตรงกันทุก service
const (
	CtxUserIDKey      = "userId"
	CtxTenantIDKey    = "tenantId"
	CtxOrgUnitIDKey   = "orgUnitId"
	CtxPermissionsKey = "permissions"
)

// SetString puts a string value into gin.Context
func SetString(c *gin.Context, key, value string) {
	if value == "" {
		return
	}
	c.Set(key, value)
}

// SetStringSlice puts []string into gin.Context
func SetStringSlice(c *gin.Context, key string, value []string) {
	if value == nil {
		return
	}
	c.Set(key, value)
}

// GetString gets a string value from gin.Context
func GetString(c *gin.Context, key string) string {
	v, exists := c.Get(key)
	if !exists {
		return ""
	}
	if s, ok := v.(string); ok {
		return s
	}
	return ""
}

// GetStringSlice gets []string value from gin.Context
func GetStringSlice(c *gin.Context, key string) []string {
	v, exists := c.Get(key)
	if !exists {
		return nil
	}
	if arr, ok := v.([]string); ok {
		return arr
	}
	return nil
}
