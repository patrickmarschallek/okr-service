package structures

import "time"

// KeyActions is a list of KeyAction
type KeyActions []KeyAction

// KeyAction is one goal to achieve an key result.
type KeyAction struct {
	ID          uint64     `json:"id"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `json:"-"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Grade       Grade      `json:"keyResults"`
}
