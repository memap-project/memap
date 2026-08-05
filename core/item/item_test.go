package item

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestItem_ExpiredItem(t *testing.T) {
	item := Item{ExpiresAt: time.Now().Unix() - 100}
	assert.True(t, item.IsExpired())
}

func TestItem_NotExpiredItem(t *testing.T) {
	item := Item{ExpiresAt: time.Now().Unix() + 100}
	assert.False(t, item.IsExpired())
}

func TestItem_WithoutExpiration(t *testing.T) {
	item := Item{ExpiresAt: 0}
	assert.False(t, item.IsExpired())
}

func TestItem_Expire(t *testing.T) {
	item := Item{
		Value:     "value1",
		ExpiresAt: 0,
	}
	item.Expire(3600)
	assert.False(t, item.IsExpired())
}
