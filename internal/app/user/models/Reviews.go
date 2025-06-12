package models

import (
	"encoding/json"
	"time"
)

type Reviews struct {
	ID             int64           `json:"id"`
	ConversationID int64           `json:"conversation_id"`
	UserID         int64           `json:"user_id"`
	SourceCode     string          `json:"source_code,omitempty"`
	Rules          string          `json:"rules,omitempty"`
	AiResult       json.RawMessage `json:"ai_result,omitempty"`
	CreatedAt      time.Time       `json:"created_at"`
}
