package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/extosoft-devsecops/hrex-iam/util"
)

// AuthContextConfig defines how to map HTTP headers to context fields.
type AuthContextConfig struct {
	IgnorePaths []string // path prefix ที่ไม่ต้องเช็ค (เช่น /health, /docs)

	HeaderUserID         string
	HeaderTenantID       string
	HeaderOrgUnitID      string
	HeaderPermissions    string
	PermissionsDelimiter string
}

// DefaultAuthContextConfig returns sane defaults for HREX platform.
func DefaultAuthContextConfig() AuthContextConfig {
	return AuthContextConfig{
		IgnorePaths: []string{
			"/health",
			"/metrics",
			"/docs",
		},
		HeaderUserID:         "X-User-Id",
		HeaderTenantID:       "X-Tenant-Id",
		HeaderOrgUnitID:      "X-Org-Unit-Id",
		HeaderPermissions:    "X-Permissions",
		PermissionsDelimiter: ",",
	}
}

// AuthContextMiddleware reads identity & permissions from headers
// and injects them into gin.Context for downstream middlewares/handlers.
func AuthContextMiddleware(cfg ...AuthContextConfig) gin.HandlerFunc {
	var c AuthContextConfig
	if len(cfg) > 0 {
		c = cfg[0]
	} else {
		c = DefaultAuthContextConfig()
	}

	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path

		// Skip ignore paths
		for _, ignore := range c.IgnorePaths {
			if strings.HasPrefix(path, ignore) {
				ctx.Next()
				return
			}
		}

		userID := strings.TrimSpace(ctx.GetHeader(c.HeaderUserID))
		tenantID := strings.TrimSpace(ctx.GetHeader(c.HeaderTenantID))
		orgUnitID := strings.TrimSpace(ctx.GetHeader(c.HeaderOrgUnitID))
		permsHeader := strings.TrimSpace(ctx.GetHeader(c.HeaderPermissions))

		if userID == "" || tenantID == "" || permsHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "unauthorized",
				"message": "missing identity or permissions in headers",
			})
			return
		}

		perms := SplitPermissions(permsHeader, c.PermissionsDelimiter)

		util.SetString(ctx, util.CtxUserIDKey, userID)
		util.SetString(ctx, util.CtxTenantIDKey, tenantID)
		util.SetString(ctx, util.CtxOrgUnitIDKey, orgUnitID)
		util.SetStringSlice(ctx, util.CtxPermissionsKey, perms)

		ctx.Next()
	}
}

func SplitPermissions(header, delimiter string) []string {
	if header == "" {
		return nil
	}
	raw := strings.Split(header, delimiter)
	out := make([]string, 0, len(raw))
	for _, v := range raw {
		v = strings.TrimSpace(v)
		if v != "" {
			out = append(out, v)
		}
	}
	return out
}
