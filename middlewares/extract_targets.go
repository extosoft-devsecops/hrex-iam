package middlewares

import (
	"bytes"
	"encoding/json"
	"io"
	"regexp"

	"github.com/gin-gonic/gin"
)

//
// ===========================
// DTO
// ===========================
//

type TargetIdentity struct {
	UserID    string
	TenantID  string
	OrgUnitID string
}

//
// ===========================
// Public API
// ===========================
//

// ExtractTargets finds target identifiers from:
// - Query string (?user_id=...)
// - Path param (:id)
// - Regex path (/v1/users/:id)
// - JSON body { "user_id": "...", "tenant_id": "...", "org_unit_id": "..." }
//
// SAFE: will restore request body for downstream consumers
func ExtractTargets(c *gin.Context) TargetIdentity {

	return TargetIdentity{
		UserID:    findUserID(c),
		TenantID:  findParamOrBody(c, "tenant_id"),
		OrgUnitID: findParamOrBody(c, "org_unit_id"),
	}
}

//
// ===========================
// Internals
// ===========================
//

// --- find user id

func findUserID(c *gin.Context) string {

	// 1) Query (?user_id=xxx)
	if v := c.Query("user_id"); v != "" {
		return v
	}

	// 2) Regex path (/v1/users/:id)
	re := regexp.MustCompile(`^/v[0-9]+/users/([^/]+)$`)
	if m := re.FindStringSubmatch(c.Request.URL.Path); len(m) == 2 {
		return m[1]
	}

	// 3) Param (:id)
	if id := c.Param("id"); id != "" {
		return id
	}

	// 4) Body
	body := readAndRestoreBody(c)
	if v, ok := body["user_id"].(string); ok {
		return v
	}

	return ""
}

// --- Generic param or json body search

func findParamOrBody(c *gin.Context, key string) string {

	// Query
	if v := c.Query(key); v != "" {
		return v
	}

	// Body
	body := readAndRestoreBody(c)
	if v, ok := body[key].(string); ok {
		return v
	}

	return ""
}

// --- Safe request body reader

func readAndRestoreBody(c *gin.Context) map[string]interface{} {

	result := map[string]interface{}{}

	if c.Request.Body == nil {
		return result
	}

	raw, err := io.ReadAll(c.Request.Body)
	if err != nil || len(raw) == 0 {
		return result
	}

	_ = json.Unmarshal(raw, &result)

	// restore body
	c.Request.Body = io.NopCloser(bytes.NewBuffer(raw))

	return result
}
