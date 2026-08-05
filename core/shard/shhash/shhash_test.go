package shhash

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestShHash_SetAndGetEmptyHash(t *testing.T) {
	shh := NewShardedHash()
	shh.HSet("hash1", 0)
	result, ok := shh.HGet("hash1")

	require.True(t, ok)
	assert.Empty(t, result)
}

func TestShHash_GetEmptyExpiredHash(t *testing.T) {
	shh := NewShardedHash()
	shh.HSet("hash2", -3600)
	result, ok := shh.HGet("hash2")

	require.False(t, ok)
	assert.Empty(t, result)
}

func TestShHash_GetNonExistingHash(t *testing.T) {
	shh := NewShardedHash()
	result, ok := shh.HGet("hash3")

	require.False(t, ok)
	assert.Empty(t, result)
}

func TestShHash_Delete(t *testing.T) {
	shh := NewShardedHash()
	shh.HSet("hash4", 0)
	shh.HDelete("hash4")
	result, ok := shh.HGet("hash4")

	require.False(t, ok)
	assert.Empty(t, result)
}

func TestShHash_SetAndGetField(t *testing.T) {
	shh := NewShardedHash()
	shh.HFSet("hash5", "filed5", "value5")
	result, ok := shh.HFGet("hash5", "filed5")

	expectedRes := "value5"

	require.True(t, ok)
	assert.Equal(t, expectedRes, result)
}

func TestShHash_GetNonExistingField(t *testing.T) {
	shh := NewShardedHash()
	result, ok := shh.HFGet("hash6", "filed6")

	require.False(t, ok)
	assert.Empty(t, result)
}

func TestShHash_DeleteField(t *testing.T) {
	shh := NewShardedHash()
	shh.HFSet("hash7", "filed7", "value7")
	shh.HFDelete("hash7", "filed7")
	result, ok := shh.HFGet("hash7", "filed7")

	require.False(t, ok)
	assert.Empty(t, result)
}

func TestShHash_GetFieldFromHashWithExpiration(t *testing.T) {
	shh := NewShardedHash()
	shh.HSet("hash8", 3600)
	shh.HFSet("hash8", "filed8", "value8")
	result, ok := shh.HFGet("hash8", "filed8")

	expectedRes := "value8"

	require.True(t, ok)
	assert.Equal(t, expectedRes, result)
}

func TestShHash_Expiration(t *testing.T) {
	shh := NewShardedHash()
	shh.HFSet("hash9", "filed9", "value9")
	shh.HExpire("hash9", 3600)
	result, ok := shh.HFGet("hash9", "filed9")

	expectedRes := "value9"

	require.True(t, ok)
	assert.Equal(t, expectedRes, result)
}

func TestShHash_GetFiledFromExpiredHash(t *testing.T) {
	shh := NewShardedHash()
	shh.HFSet("hash10", "filed10", "value10")
	shh.HExpire("hash10", -3600)
	result, ok := shh.HFGet("hash10", "filed10")

	require.False(t, ok)
	assert.Empty(t, result)
}

func TestShHash_Clean(t *testing.T) {
	shh := NewShardedHash()
	shh.HFSet("hash01", "filed01", "value01")
	shh.HFSet("hash02", "filed02", "value02")
	shh.Clean()

	result, ok := shh.HFGet("hash01", "filed01")
	require.True(t, ok)
	assert.NotEmpty(t, result)
	expectedRes := "value01"
	assert.Equal(t, expectedRes, result)

	result, ok = shh.HFGet("hash02", "filed02")
	require.True(t, ok)
	assert.NotEmpty(t, result)
	expectedRes = "value02"
	assert.Equal(t, expectedRes, result)
}
