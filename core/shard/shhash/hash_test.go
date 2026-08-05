package shhash

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHash_ExpiredHash(t *testing.T) {
	h := NewHash()
	h.expiresAt = time.Now().Unix() - 3600
	ok := h.IsExpired()

	assert.True(t, ok)
}

func TestHash_ExpireForEmpty(t *testing.T) {
	h := NewHash()
	h.Expire(3600)
	ok := h.IsExpired()

	assert.False(t, ok)
}

func TestHash_ExpireForZero(t *testing.T) {
	h := NewHash()
	h.expiresAt = 0
	h.Expire(3600)
	ok := h.IsExpired()

	assert.False(t, ok)
}

func TestHash_ExpireForExpired(t *testing.T) {
	h := NewHash()
	h.expiresAt = time.Now().Unix() - 3600
	h.Expire(3600)
	ok := h.IsExpired()

	assert.True(t, ok)
}

func TestHash_SetAndGet(t *testing.T) {
	h := NewHash()
	h.Set("key1", "value1")
	result, ok := h.Get("key1")

	expectedRes := "value1"

	require.True(t, ok)
	assert.Equal(t, expectedRes, result)
}

func TestHash_GetNonExistentKey(t *testing.T) {
	h := NewHash()
	result, ok := h.Get("key2")

	require.False(t, ok)
	assert.Empty(t, result)
}

func TestHash_SetGetDelete(t *testing.T) {
	h := NewHash()
	h.Set("key3", "value3")
	result, ok := h.Get("key3")

	expectedRes := "value3"

	require.True(t, ok)
	assert.Equal(t, expectedRes, result)

	h.Delete("key3")
	result, ok = h.Get("key3")

	require.False(t, ok)
	assert.Empty(t, result)
}

func TestHash_GetCopy(t *testing.T) {
	h := NewHash()
	h.Set("key1", "value1")
	copy := h.GetCopy()
	result, ok := copy["key1"]

	expectedRes := "value1"

	require.NotNil(t, copy)
	require.True(t, ok)
	assert.Equal(t, expectedRes, result)
}
