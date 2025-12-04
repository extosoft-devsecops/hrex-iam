package middlewares_test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/extosoft-devsecops/hrex-iam/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

//
// Helpers
//

func createGinContext(method, path string, body string) *gin.Context {

	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest(method, path, bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	c.Request = req
	return c
}

//
// ==========================
// Tests
// ==========================
//

func Test_ExtractTargets_FromQueryString(t *testing.T) {

	c := createGinContext(
		http.MethodGet,
		"/users?user_id=U001&tenant_id=T001&org_unit_id=D001",
		"",
	)

	target := middlewares.ExtractTargets(c)

	assert.Equal(t, "U001", target.UserID)
	assert.Equal(t, "T001", target.TenantID)
	assert.Equal(t, "D001", target.OrgUnitID)
}

func Test_ExtractTargets_FromParam(t *testing.T) {

	c := createGinContext(
		http.MethodGet,
		"/users/999",
		"",
	)

	// Gin param mocking
	c.Params = gin.Params{
		{Key: "id", Value: "999"},
	}

	target := middlewares.ExtractTargets(c)

	assert.Equal(t, "999", target.UserID)
}

func Test_ExtractTargets_FromRegexPath(t *testing.T) {

	c := createGinContext(
		http.MethodGet,
		"/v1/users/ABC123",
		"",
	)

	target := middlewares.ExtractTargets(c)

	assert.Equal(t, "ABC123", target.UserID)
}

func Test_ExtractTargets_FromBody(t *testing.T) {

	payload := `
	{
		"user_id":"U789",
		"tenant_id":"TENANT77",
		"org_unit_id":"IT01"
	}`

	c := createGinContext(
		http.MethodPost,
		"/users",
		payload,
	)

	target := middlewares.ExtractTargets(c)

	assert.Equal(t, "U789", target.UserID)
	assert.Equal(t, "TENANT77", target.TenantID)
	assert.Equal(t, "IT01", target.OrgUnitID)
}

func Test_ExtractTargets_UserID_InBody_WithNoParam(t *testing.T) {

	payload := `{ "user_id" : "U555" }`

	c := createGinContext(
		http.MethodPost,
		"/v1/users",
		payload,
	)

	target := middlewares.ExtractTargets(c)

	assert.Equal(t, "U555", target.UserID)
}

func Test_ExtractTargets_Tenant_Only_InBody(t *testing.T) {

	payload := `{ "tenant_id": "TEN888" }`

	c := createGinContext(
		http.MethodPost,
		"/any",
		payload,
	)

	target := middlewares.ExtractTargets(c)

	assert.Equal(t, "", target.UserID)
	assert.Equal(t, "TEN888", target.TenantID)
	assert.Equal(t, "", target.OrgUnitID)
}

func Test_ExtractTargets_Body_Restore_After_Read(t *testing.T) {

	payload := `{ "user_id":"RESTORE_TEST" }`

	c := createGinContext(
		http.MethodPost,
		"/users",
		payload,
	)

	// first read
	first := middlewares.ExtractTargets(c)
	assert.Equal(t, "RESTORE_TEST", first.UserID)

	rawAfter, err := io.ReadAll(c.Request.Body)

	assert.NoError(t, err)
	assert.Contains(t, string(rawAfter), "RESTORE_TEST")
}
