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
