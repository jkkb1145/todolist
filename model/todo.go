package model

import "time"

type CreatNewList struct {
	ListID    int
	UserID    int       `json:"user_id"`
	Type      string    `json:"type"`
	Title     string    `json:"title"`
	Info      string    `json:"info"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateList struct {
	Type      string    `json:"type"`
	Title     string    `json:"title"`
	Info      string    `json:"info"`
	UpdatedAt time.Time `json:"updated_at"`
}
