package util_test

import (
	"testing"

	"github.com/extosoft-devsecops/hrex-iam/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpContext() *gin.Context {
	c, _ := gin.CreateTestContext(nil)
	return c
}

func TestStringIsStoredAndRetrievedCorrectly(t *testing.T) {
	c := SetUpContext()
	util.SetString(c, util.CtxUserIDKey, "user123")
	val := util.GetString(c, util.CtxUserIDKey)
	assert.Equal(t, "user123", val)
}

func TestStringIsNotStoredIfEmpty(t *testing.T) {
	c := SetUpContext()
	util.SetString(c, util.CtxUserIDKey, "")
	val := util.GetString(c, util.CtxUserIDKey)
	assert.Equal(t, "", val)
}

func TestStringSliceIsStoredAndRetrievedCorrectly(t *testing.T) {
	c := SetUpContext()
	util.SetStringSlice(c, util.CtxPermissionsKey, []string{"a", "b"})
	val := util.GetStringSlice(c, util.CtxPermissionsKey)
	assert.Equal(t, []string{"a", "b"}, val)
}

func TestStringSliceIsNotStoredIfNil(t *testing.T) {
	c := SetUpContext()
	util.SetStringSlice(c, util.CtxPermissionsKey, nil)
	val := util.GetStringSlice(c, util.CtxPermissionsKey)
	assert.Nil(t, val)
}

func TestGetStringReturnsEmptyIfKeyMissing(t *testing.T) {
	c := SetUpContext()
	val := util.GetString(c, "missingKey")
	assert.Equal(t, "", val)
}

func TestGetStringSliceReturnsNilIfKeyMissing(t *testing.T) {
	c := SetUpContext()
	val := util.GetStringSlice(c, "missingKey")
	assert.Nil(t, val)
}

func TestGetStringReturnsEmptyIfTypeMismatch(t *testing.T) {
	c := SetUpContext()
	c.Set(util.CtxUserIDKey, 123)
	val := util.GetString(c, util.CtxUserIDKey)
	assert.Equal(t, "", val)
}

func TestGetStringSliceReturnsNilIfTypeMismatch(t *testing.T) {
	c := SetUpContext()
	c.Set(util.CtxPermissionsKey, "notaslice")
	val := util.GetStringSlice(c, util.CtxPermissionsKey)
	assert.Nil(t, val)
}
