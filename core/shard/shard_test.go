package shard

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestShard_SetAndGet tests the Set and Get methods of the Shard.
func TestShard_SetAndGet(t *testing.T) {
	shard := NewShard[string]()
	shard.Set("key1", "value1")
	result, ok := shard.Get("key1")

	expectedRes := "value1"

	require.True(t, ok)
	assert.Equal(t, expectedRes, result)
}

// TestShard_GetNonExistingKey tests the Get method of the Shard with a non-existing key.
func TestShard_GetNonExistingKey(t *testing.T) {
	shard := NewShard[string]()
	result, ok := shard.Get("key2")

	expectedRes := ""

	require.False(t, ok)
	assert.Equal(t, expectedRes, result)
}

// TestShard_GetOrInit tests the GetOrInit method of the Shard.
func TestShard_GetOrInit(t *testing.T) {
	shard := NewShard[string]()
	result, _ := shard.GetOrInit("key3", func() string {
		return "value3"
	})

	expectedRes := "value3"

	require.NotEmpty(t, result)
	assert.Equal(t, expectedRes, result)
}

func TestShard_Delete(t *testing.T) {
	shard := NewShard[string]()
	shard.Set("key4", "value4")
	shard.Delete("key4")
	result, ok := shard.Get("key4")

	expectedRes := ""

	require.False(t, ok)
	assert.Equal(t, expectedRes, result)
}

func TestShard_Clean(t *testing.T) {
	shard := NewShard[string]()
	shard.Set("key5", "value5")
	shard.Set("key55", "value55")
	shard.Set("key555", "value555")
	shard.Clean(
		func(key string, value string) bool {
			return strings.HasPrefix(key, "key5")
		},
	)
	result1, ok1 := shard.Get("key5")
	result2, ok2 := shard.Get("key55")
	result3, ok3 := shard.Get("key555")

	expectedRes := ""

	require.False(t, ok1)
	assert.Equal(t, expectedRes, result1)

	require.False(t, ok2)
	assert.Equal(t, expectedRes, result2)

	require.False(t, ok3)
	assert.Equal(t, expectedRes, result3)
}

func TestShard_SetIdenticalKeys(t *testing.T) {
	shard := NewShard[string]()
	shard.Set("key6", "value6")
	shard.Set("key6", "value66")
	result, ok := shard.Get("key6")

	expectedRes := "value66"

	require.True(t, ok)
	assert.Equal(t, expectedRes, result)
}

func TestShard_SetAndDeleteIdenticalKeys(t *testing.T) {
	shard := NewShard[string]()
	shard.Set("key7", "value7")
	shard.Set("key7", "value77")
	shard.Delete("key7")
	result, ok := shard.Get("key7")

	require.False(t, ok)
	assert.Empty(t, result)
}

func TestShard_Update(t *testing.T) {
	shard := NewShard[string]()
	shard.Set("key8", "value8")
	shard.Update("key8", func(value *string) bool {
		*value = *value + "updated"
		return true
	})
	result, ok := shard.Get("key8")

	expectedRes := "value8updated"

	require.True(t, ok)
	assert.Equal(t, expectedRes, result)
}
