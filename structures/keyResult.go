package structures

import "time"

// KeyResults is a list of KeyResult
type KeyResults []KeyResult

// KeyResult is one goal to achieve an objective.
type KeyResult struct {
	ID          uint64     `json:"id"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `json:"-"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Grade       Grade      `json:"keyResults"`
}
