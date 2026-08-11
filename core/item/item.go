package item

import "time"

// Item represents a string value that expires after a certain time.
type Item struct {
	Value     string
	ExpiresAt int64 // timestamp in seconds since epoch
}

// IsExpired returns true if the item has expired.
// If ExpiresAt is 0, the item has no expiration and returns false.
func (i *Item) IsExpired() bool {
	if i.ExpiresAt == 0 {
		return false
	}
	return time.Now().Unix() > i.ExpiresAt
}

// Expire sets the item to expire after the given TTL in seconds.
func (i *Item) Expire(ttl int64) {
	if i.ExpiresAt > 0 || i.IsExpired() {
		return
	}
	i.ExpiresAt = time.Now().Unix() + ttl
}

// LeftTime returns the number of seconds until the item expires.
func (i *Item) LeftTime() int64 {
	if i.ExpiresAt == 0 {
		return 0
	}
	return i.ExpiresAt - time.Now().Unix()
}
