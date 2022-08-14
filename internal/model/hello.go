package model

import "time"

type Hello struct {
	Id        int         `json:"id,omitempty"`
	Request   interface{} `json:"request,omitempty"`
	Response  string      `json:"response,omitempty"`
	CreatedAt time.Time   `json:"created_at,omitempty"`
	UpdatedAt time.Time   `json:"updated_at,omitempty"`
}
