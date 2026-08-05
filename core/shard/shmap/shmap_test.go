package shmap

import (
	"testing"
	"time"

	"github.com/dmi3midd/memap/core/item"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestShmap_SetAndGet(t *testing.T) {
	shmap := NewShardedMap()
	i := item.Item{
		Value:     "value1",
		ExpiresAt: time.Now().Unix() + 3600,
	}
	shmap.Set("key", i)
	result, ok := shmap.Get("key")

	expectedRes := item.Item{
		Value: "value1",
	}

	require.True(t, ok)
	require.NotEmpty(t, result)
	assert.Equal(t, expectedRes.Value, result.Value)
}

func TestShmap_GetNonExistingKey(t *testing.T) {
	shmap := NewShardedMap()
	result, ok := shmap.Get("key2")

	require.False(t, ok)
	assert.Empty(t, result)
}

func TestShmap_SetAndGetExpired(t *testing.T) {
	shmap := NewShardedMap()
	i := item.Item{
		Value:     "value3",
		ExpiresAt: time.Now().Unix() - 3600,
	}
	shmap.Set("key3", i)
	result, ok := shmap.Get("key3")

	require.False(t, ok)
	assert.Empty(t, result)
}

func TestShmap_Delete(t *testing.T) {
	shmap := NewShardedMap()
	i := item.Item{
		Value:     "value4",
		ExpiresAt: time.Now().Unix() + 3600,
	}
	shmap.Set("key4", i)
	shmap.Delete("key4")
	result, ok := shmap.Get("key4")

	require.False(t, ok)
	assert.Empty(t, result)
}

func TestShmap_Clean(t *testing.T) {
	shmap := NewShardedMap()
	i := item.Item{
		Value:     "value5",
		ExpiresAt: time.Now().Unix() - 3600,
	}
	shmap.Set("key5", i)
	shmap.Set("key55", i)
	shmap.Set("key555", i)
	shmap.Clean()
	result1, ok1 := shmap.Get("key5")
	result2, ok2 := shmap.Get("key55")
	result3, ok3 := shmap.Get("key555")

	require.False(t, ok1)
	assert.Empty(t, result1)

	require.False(t, ok2)
	assert.Empty(t, result2)

	require.False(t, ok3)
	assert.Empty(t, result3)
}

func TestShmap_SetIdenticalKeys(t *testing.T) {
	shmap := NewShardedMap()
	shmap.Set("key6", item.Item{Value: "value6", ExpiresAt: time.Now().Unix() + 3600})
	shmap.Set("key6", item.Item{Value: "value66", ExpiresAt: time.Now().Unix() + 3600})
	result, ok := shmap.Get("key6")

	expectedRes := "value66"

	require.True(t, ok)
	assert.Equal(t, expectedRes, result.Value)
}

func TestShmap_SetAndDeleteIdenticalKeys(t *testing.T) {
	shmap := NewShardedMap()
	shmap.Set("key7", item.Item{Value: "value7", ExpiresAt: time.Now().Unix() + 3600})
	shmap.Set("key7", item.Item{Value: "value77", ExpiresAt: time.Now().Unix() + 3600})
	shmap.Delete("key7")
	result, ok := shmap.Get("key7")

	require.False(t, ok)
	assert.Empty(t, result)
}

func TestShmap_Expire(t *testing.T) {
	shmap := NewShardedMap()
	i := item.Item{
		Value:     "value8",
		ExpiresAt: 0,
	}
	shmap.Set("key8", i)
	shmap.Expire("key8", 3600)
	result, ok := shmap.Get("key8")

	expectedRes := "value8"

	require.True(t, ok)
	assert.Equal(t, expectedRes, result.Value)
}
