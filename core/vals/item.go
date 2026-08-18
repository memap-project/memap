package vals

import "time"

// Item represents a string value that expires after a certain time.
type Item struct {
	value     string
	expiresAt int64 // timestamp in seconds since epoch
}

func NewItem() *Item {
	return &Item{}
}

// GetValue returns the value of the item.
func (i *Item) GetValue() string {
	return i.value
}

// SetValue sets the value of the item.
func (i *Item) SetValue(value string) {
	i.value = value
}

// IsExpired returns true if the item has expired.
// If ExpiresAt is 0, the item has no expiration and returns false.
func (i *Item) IsExpired() bool {
	if i.expiresAt == 0 {
		return false
	}
	return time.Now().Unix() > i.expiresAt
}

// Expire sets the item to expire after the given TTL in seconds.
func (i *Item) Expire(ttl int64) bool {
	if i.IsExpired() {
		return false
	}
	if ttl <= 0 {
		i.expiresAt = 0
		return true
	}
	i.expiresAt = time.Now().Unix() + ttl
	return true
}

// LeftTime returns the number of seconds until the item expires.
func (i *Item) LeftTime() int64 {
	if i.expiresAt == 0 {
		return 0
	}
	return i.expiresAt - time.Now().Unix()
}
