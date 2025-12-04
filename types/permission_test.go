package types_test

import (
	"testing"

	"github.com/extosoft-devsecops/hrex-iam/types"
	"github.com/stretchr/testify/assert"
)

func PermissionStringIsParsedCorrectly(t *testing.T) {
	p := types.ParsePermission("doc:read:global")
	assert.Equal(t, "doc", p.Resource)
	assert.Equal(t, "read", p.Action)
	assert.Equal(t, "global", p.Scope)
}

func PermissionStringWithMissingScopeIsParsedCorrectly(t *testing.T) {
	p := types.ParsePermission("doc:read")
	assert.Equal(t, "doc", p.Resource)
	assert.Equal(t, "read", p.Action)
	assert.Equal(t, "", p.Scope)
}

func PermissionStringWithOnlyResourceIsParsedCorrectly(t *testing.T) {
	p := types.ParsePermission("doc")
	assert.Equal(t, "doc", p.Resource)
	assert.Equal(t, "", p.Action)
	assert.Equal(t, "", p.Scope)
}

func PermissionStringEmptyReturnsEmptyPermission(t *testing.T) {
	p := types.ParsePermission("")
	assert.Equal(t, "", p.Resource)
	assert.Equal(t, "", p.Action)
	assert.Equal(t, "", p.Scope)
}

func HasPermissionReturnsTrueForExactMatch(t *testing.T) {
	perms := []string{"doc:read:global"}
	required := types.Permission{Resource: "doc", Action: "read", Scope: "global"}
	assert.True(t, types.HasPermission(perms, required))
}

func HasPermissionReturnsTrueForBroaderScope(t *testing.T) {
	perms := []string{"doc:read:global"}
	required := types.Permission{Resource: "doc", Action: "read", Scope: "tenant"}
	assert.True(t, types.HasPermission(perms, required))
}

func HasPermissionReturnsFalseForNarrowerScope(t *testing.T) {
	perms := []string{"doc:read:self"}
	required := types.Permission{Resource: "doc", Action: "read", Scope: "global"}
	assert.False(t, types.HasPermission(perms, required))
}

func HasPermissionReturnsFalseForDifferentResource(t *testing.T) {
	perms := []string{"user:read:global"}
	required := types.Permission{Resource: "doc", Action: "read", Scope: "global"}
	assert.False(t, types.HasPermission(perms, required))
}

func HasPermissionReturnsFalseForDifferentAction(t *testing.T) {
	perms := []string{"doc:write:global"}
	required := types.Permission{Resource: "doc", Action: "read", Scope: "global"}
	assert.False(t, types.HasPermission(perms, required))
}

func HasPermissionReturnsFalseIfNoPermissions(t *testing.T) {
	perms := []string{}
	required := types.Permission{Resource: "doc", Action: "read", Scope: "global"}
	assert.False(t, types.HasPermission(perms, required))
}

func ScopeMatchReturnsTrueForEqualScope(t *testing.T) {
	assert.True(t, types.ScopeMatch(types.ScopeTenant, types.ScopeTenant))
}

func ScopeMatchReturnsTrueForBroaderScope(t *testing.T) {
	assert.True(t, types.ScopeMatch(types.ScopeGlobal, types.ScopeTenant))
}

func ScopeMatchReturnsFalseForNarrowerScope(t *testing.T) {
	assert.False(t, types.ScopeMatch(types.ScopeSelf, types.ScopeDepartment))
}

func ScopeMatchReturnsTrueForUnknownRequiredScope(t *testing.T) {
	assert.True(t, types.ScopeMatch(types.ScopeGlobal, "unknown"))
}

func ScopeMatchReturnsFalseForUnknownUserScope(t *testing.T) {
	assert.False(t, types.ScopeMatch("unknown", types.ScopeGlobal))
}

func TestPermissionStringIsParsedCorrectly(t *testing.T) {
	PermissionStringIsParsedCorrectly(t)
}

func TestPermissionStringWithMissingScopeIsParsedCorrectly(t *testing.T) {
	PermissionStringWithMissingScopeIsParsedCorrectly(t)
}

func TestPermissionStringWithOnlyResourceIsParsedCorrectly(t *testing.T) {
	PermissionStringWithOnlyResourceIsParsedCorrectly(t)
}

func TestPermissionStringEmptyReturnsEmptyPermission(t *testing.T) {
	PermissionStringEmptyReturnsEmptyPermission(t)
}

func TestHasPermissionReturnsTrueForExactMatch(t *testing.T) {
	HasPermissionReturnsTrueForExactMatch(t)
}

func TestHasPermissionReturnsTrueForBroaderScope(t *testing.T) {
	HasPermissionReturnsTrueForBroaderScope(t)
}

func TestHasPermissionReturnsFalseForNarrowerScope(t *testing.T) {
	HasPermissionReturnsFalseForNarrowerScope(t)
}

func TestHasPermissionReturnsFalseForDifferentResource(t *testing.T) {
	HasPermissionReturnsFalseForDifferentResource(t)
}

func TestHasPermissionReturnsFalseForDifferentAction(t *testing.T) {
	HasPermissionReturnsFalseForDifferentAction(t)
}

func TestHasPermissionReturnsFalseIfNoPermissions(t *testing.T) {
	HasPermissionReturnsFalseIfNoPermissions(t)
}

func TestScopeMatchReturnsTrueForEqualScope(t *testing.T) {
	ScopeMatchReturnsTrueForEqualScope(t)
}

func TestScopeMatchReturnsTrueForBroaderScope(t *testing.T) {
	ScopeMatchReturnsTrueForBroaderScope(t)
}

func TestScopeMatchReturnsFalseForNarrowerScope(t *testing.T) {
	ScopeMatchReturnsFalseForNarrowerScope(t)
}

func TestScopeMatchReturnsTrueForUnknownRequiredScope(t *testing.T) {
	ScopeMatchReturnsTrueForUnknownRequiredScope(t)
}

func TestScopeMatchReturnsFalseForUnknownUserScope(t *testing.T) {
	ScopeMatchReturnsFalseForUnknownUserScope(t)
}
