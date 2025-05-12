package models

import "time"

type Notification struct {
	ID        int       `json:"id"`
	EventID   int       `json:"event_id"`
	NotifyAt  time.Time `json:"notify_at"`
	Delivered bool      `json:"delivered"`
}
