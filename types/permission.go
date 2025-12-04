package types

import "strings"

// Permission is the normalized permission model:
// <RESOURCE>:<ACTION>:<SCOPE>
type Permission struct {
	Resource string
	Action   string
	Scope    string
}

// Scope constants (convention)
const (
	ScopeSelf       = "self"
	ScopeDepartment = "department"
	ScopeTenant     = "tenant"
	ScopeGlobal     = "global"
)

// ParsePermission parses "resource:action:scope" string into Permission struct.
func ParsePermission(s string) Permission {
	parts := strings.Split(s, ":")
	p := Permission{}
	if len(parts) > 0 {
		p.Resource = parts[0]
	}
	if len(parts) > 1 {
		p.Action = parts[1]
	}
	if len(parts) > 2 {
		p.Scope = parts[2]
	}
	return p
}

// HasPermission checks if userPerms (from header/context)
// contain at least one permission that satisfies the required permission
// with scope dominance rules.
func HasPermission(userPerms []string, required Permission) bool {
	for _, permStr := range userPerms {
		p := ParsePermission(strings.TrimSpace(permStr))
		if p.Resource == required.Resource &&
			p.Action == required.Action &&
			ScopeMatch(p.Scope, required.Scope) {
			return true
		}
	}
	return false
}

// scopeOrder defines dominance ranking for scopes
// higher number = broader scope.
var scopeOrder = map[string]int{
	ScopeSelf:       1,
	ScopeDepartment: 2,
	ScopeTenant:     3,
	ScopeGlobal:     4,
}

// scopeRank returns the integer ranking of a scope.
func scopeRank(scope string) int {
	if v, ok := scopeOrder[scope]; ok {
		return v
	}
	// unknown scope = lowest
	return 0
}

// ScopeMatch returns true if userScope is equal or broader than requiredScope.
func ScopeMatch(userScope, requiredScope string) bool {
	return scopeRank(userScope) >= scopeRank(requiredScope)
}
