package model

import (
	"time"
)

type Tweet struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"createdAt"`
	UserID    string    `json: userId`
	User      *User     `json:"user" gorm:"OnUpdate:CASCADE, OnDelete: CASCADE"`
}
