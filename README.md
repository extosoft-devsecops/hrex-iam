# 🔐 HREX-IAM

> Policy-based Identity & Access Management for Gin Microservices  
> Hybrid **RBAC + PBAC** with Scope Enforcement

---

![Go](https://img.shields.io/badge/go-1.22+-00ADD8?logo=go)
![Gin](https://img.shields.io/badge/gin-framework-brightgreen)
![License](https://img.shields.io/badge/license-MIT-blue)
![Release](https://img.shields.io/github/v/tag/extosoft-devsecops/hrex-iam)

---

## 📦 Latest Release

---

## ✅ What’s Included

### 🔑 Core

- Header-based identity middleware (`AuthContextMiddleware`)
- Permission middleware with scope validation
- Permission Format:

```txt
<resource>:<action>:<scope>
```

Examples:

```txt  
users:view:self
users:update:tenant
users:delete:global
employees:view:department
```

### 🎯 Supported Scopes

| Scope        | Description                |
|--------------|----------------------------|
| `global`     | Entire organization access |
| `tenant`     | Tenant scoped access       |
| `department` | Department scoped access   |
| `self`       | Own resource only          |

Scope hierarchy:

```text
global > tenant > department > self
```


### 🔄 Scope Resolvers

Built-in helpers:
- ScopeGlobal()
- ScopeTenant()
- ScopeSelfOnly()
- ScopeSelfOrTenantFromParam(param)
- ScopeFromCustomFunc(fn)
  🎯 Target Resolver


### Extract resource targeting from:

Query string

URL Param :id

API versioned paths /vX/users/:id

JSON body

```
target := middlewares.ExtractTargets(c)

fmt.Println(target.UserID)
fmt.Println(target.TenantID)
fmt.Println(target.OrgUnitID)
```

Returns:
```golang
type TargetIdentity struct {
    UserID    string
    TenantID  string
    OrgUnitID string
}
```

### 📥 Installation

```shell
go get github.com/extosoft-devsecops/hrex-iam@latest

```

Or pin a version:

```shell
go get github.com/extosoft-devsecops/hrex-iam@v1.0.0

```
