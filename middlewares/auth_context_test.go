package middlewares_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/extosoft-devsecops/hrex-iam/middlewares"
	"github.com/extosoft-devsecops/hrex-iam/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthContextMiddleware_Success(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest("GET", "/api/resource", nil)
	c.Request.Header.Set("X-User-Id", "user1")
	c.Request.Header.Set("X-Tenant-Id", "tenant1")
	c.Request.Header.Set("X-Org-Unit-Id", "org1")
	c.Request.Header.Set("X-Permissions", "doc:read:global,user:update:org")

	mw := middlewares.AuthContextMiddleware()
	mw(c)

	userID := util.GetString(c, util.CtxUserIDKey)
	tenantID := util.GetString(c, util.CtxTenantIDKey)
	orgUnitID := util.GetString(c, util.CtxOrgUnitIDKey)
	perms := util.GetStringSlice(c, util.CtxPermissionsKey)

	assert.Equal(t, "user1", userID)
	assert.Equal(t, "tenant1", tenantID)
	assert.Equal(t, "org1", orgUnitID)
	assert.Equal(t, []string{"doc:read:global", "user:update:org"}, perms)
}

func TestAuthContextMiddleware_MissingHeaders(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/api/resource", nil)
	// Missing headers
	mw := middlewares.AuthContextMiddleware()
	mw(c)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestAuthContextMiddleware_IgnorePath(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/health", nil)
	mw := middlewares.AuthContextMiddleware()
	mw(c)
	// Should not abort, should continue
	assert.Equal(t, 200, w.Code)
}

func TestSplitPermissions(t *testing.T) {
	perms := middlewares.SplitPermissions("a,b,c", ",")
	assert.Equal(t, []string{"a", "b", "c"}, perms)
	perms = middlewares.SplitPermissions("", ",")
	assert.Nil(t, perms)
}
