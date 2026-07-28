package item

import "time"

// Item represents a string value that expires after a certain time.
type Item struct {
	Value     string
	ExpiresAt int64 // timestamp in seconds since epoch (for easy comparison with time.Now().Unix())
}

// IsExpired returns true if the item has expired.
func (i *Item) IsExpired() bool {
	if i.ExpiresAt == 0 {
		return false
	}
	return time.Now().Unix() > i.ExpiresAt
}
