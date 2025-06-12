package models

import (
	"time"
)

type Conversations struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	ReviewID  int64     `json:"review_id"`
	CreatedAt time.Time `json:"created_at"`
}
