package model

import "time"

type User struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	City      string    `json:"city,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
