package item

import "time"

// Item represents a string value that expires after a certain time.
type Item struct {
	Value     string
	ExpiresAt time.Time
}

// IsExpired returns true if the item has expired.
func (i *Item) IsExpired() bool {
	if i.ExpiresAt.IsZero() {
		return false
	}
	return time.Now().After(i.ExpiresAt)
}
